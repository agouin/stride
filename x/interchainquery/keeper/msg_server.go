package keeper

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/v3/modules/core/23-commitment/types"
	tmclienttypes "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint/types"
	"github.com/spf13/cast"

	"github.com/Stride-Labs/stride/x/interchainquery/types"
)

type msgServer struct {
	*Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: &keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) SubmitQueryResponse(goCtx context.Context, msg *types.MsgSubmitQueryResponse) (*types.MsgSubmitQueryResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp]... ICQ FROM GO RELAYER HAS MSG: %#v", msg))

	q, found := k.GetQuery(ctx, msg.QueryId)
	if !found {
		k.Logger(ctx).Info("Ignoring duplicate query")
		return &types.MsgSubmitQueryResponseResponse{}, nil // technically this is an error, but will cause the entire tx to fail if we have one 'bad' message, so we can just no-op here.
	}

	defer func() {
		k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] query_response event emission: %s", q.Id))
		ctx.EventManager().EmitEvent(sdk.NewEvent(
			"query_response",
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(types.AttributeKeyChainId, q.ChainId),
			sdk.NewAttribute(types.AttributeKeyQueryId, q.Id),
		))
	}()

	// ABORT PROCESSING QUERY RESPONSE IF WE EXCEEDED THE TTL
	k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] q: %#v.", q))
	k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] query %s with ttl: %d, resp time: %d.", msg.QueryId, q.Ttl, ctx.BlockHeader().Time.UnixNano()))
	curT, err := cast.ToUint64E(ctx.BlockTime().UnixNano())
	if err != nil {
		k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] error gathering block time!"))
		return nil, err
	}

	if q.Ttl < curT {
		errMsg := fmt.Sprintf("[ICQ Resp] aborting query callback due to ttl expiry! ttl is %d, time now %d for query of type %s with id %s, on chain %s", q.Ttl, ctx.BlockHeader().Time.UnixNano(), q.QueryType, q.ChainId, msg.QueryId)
		k.DeleteQuery(ctx, msg.QueryId)
		k.Logger(ctx).Error(errMsg)
		return &types.MsgSubmitQueryResponseResponse{}, nil
	}

	// PROCESS QUERY RESPONSE
	pathParts := strings.Split(q.QueryType, "/")

	// verify the query results are valid by checking the assocaited proof
	if pathParts[len(pathParts)-1] == "key" {
		k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] key query...!"))
		if msg.ProofOps == nil {
			errMsg := "unable to validate proof. No proof submitted"
			return nil, fmt.Errorf(errMsg)
		}
		connection, _ := k.IBCKeeper.ConnectionKeeper.GetConnection(ctx, q.ConnectionId)
		k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] 68!"))

		msgHeight, err := cast.ToUint64E(msg.Height)
		if err != nil {
			return nil, err
		}
		k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] 74!"))

		height := clienttypes.NewHeight(clienttypes.ParseChainID(q.ChainId), msgHeight+1)
		consensusState, found := k.IBCKeeper.ClientKeeper.GetClientConsensusState(ctx, connection.ClientId, height)
		k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] 68!"))

		if !found {
			return nil, fmt.Errorf("unable to fetch consensus state")
		}
		k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] 83!"))

		clientState, found := k.IBCKeeper.ClientKeeper.GetClientState(ctx, connection.ClientId)
		if !found {
			return nil, fmt.Errorf("unable to fetch client state")
		}
		path := commitmenttypes.NewMerklePath([]string{pathParts[1], url.PathEscape(string(q.Request))}...)
		k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] 90!"))

		merkleProof, err := commitmenttypes.ConvertProofs(msg.ProofOps)
		if err != nil {
			return nil, fmt.Errorf("error converting proofs")
		}
		k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] 96!"))

		tmclientstate, ok := clientState.(*tmclienttypes.ClientState)
		if !ok {
			return nil, fmt.Errorf("error unmarshaling client state %v", clientState)
		}

		k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp] key checking len...!"))

		if len(msg.Result) != 0 {
			// if we got a non-nil response, verify inclusion proof.
			if err := merkleProof.VerifyMembership(tmclientstate.ProofSpecs, consensusState.GetRoot(), path, msg.Result); err != nil {
				return nil, fmt.Errorf("unable to verify proof: %s", err)
			}
			k.Logger(ctx).Info(fmt.Sprintf("Proof validated! module: %s, queryId %s", types.ModuleName, q.Id))

		} else {
			// if we got a nil response, verify non inclusion proof.
			if err := merkleProof.VerifyNonMembership(tmclientstate.ProofSpecs, consensusState.GetRoot(), path); err != nil {
				return nil, fmt.Errorf("unable to verify proof: %s", err)
			}
			k.Logger(ctx).Info(fmt.Sprintf("Non-inclusion Proof validated! module: %s, queryId %s", types.ModuleName, q.Id))
		}

		moduleNames := []string{}
		for moduleName := range k.callbacks {
			moduleNames = append(moduleNames, moduleName)
		}
		sort.Strings(moduleNames)

		for _, module := range moduleNames {
			k.Logger(ctx).Info(fmt.Sprintf("Executing callback for queryId (%s), module (%s)", q.Id, module))
			moduleCallbackHandler := k.callbacks[module]

			if moduleCallbackHandler.Has(q.CallbackId) {
				k.Logger(ctx).Info(fmt.Sprintf("ICQ Callback (%s) found for module (%s)", q.CallbackId, module))
				err := moduleCallbackHandler.Call(ctx, q.CallbackId, msg.Result, q)
				if err != nil {
					k.Logger(ctx).Error(fmt.Sprintf("error in ICQ callback, error: %s, msg: %s, result: %v, type: %s, params: %v", err.Error(), msg.QueryId, msg.Result, q.QueryType, q.Request))
					return nil, err
				}
			} else {
				k.Logger(ctx).Info(fmt.Sprintf("ICQ Callback not found for module (%s)", module))
			}
		}

		k.DeleteQuery(ctx, msg.QueryId)
	}

	k.Logger(ctx).Info(fmt.Sprintf("[ICQ Resp2] q: %#v.", q))
	k.DeleteQuery(ctx, msg.QueryId)

	return &types.MsgSubmitQueryResponseResponse{}, nil
}
