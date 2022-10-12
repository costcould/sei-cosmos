package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	acltypes "github.com/cosmos/cosmos-sdk/types/accesscontrol"
	"github.com/gogo/protobuf/proto"
)

var ErrNoCommitAccessOp = fmt.Errorf("MessageDependencyMapping doesn't terminate with AccessType_COMMIT")

type MessageKey string

func GenerateMessageKey(msg sdk.Msg) MessageKey {
	return MessageKey(proto.MessageName(msg))
}

func CommitAccessOp() acltypes.AccessOperation {
	return acltypes.AccessOperation{ResourceType: acltypes.ResourceType_ANY, AccessType: acltypes.AccessType_COMMIT, IdentifierTemplate: "*"}
}

// Validates access operation sequence for a message, requires the last access operation to be a COMMIT
func ValidateAccessOps(accessOps []acltypes.AccessOperation) error {
	lastAccessOp := accessOps[len(accessOps)-1]
	if lastAccessOp != CommitAccessOp() {
		return ErrNoCommitAccessOp
	}
	return nil
}

func ValidateMessageDependencyMapping(mapping acltypes.MessageDependencyMapping) error {
	return ValidateAccessOps(mapping.AccessOps)
}

func SynchronousMessageDependencyMapping(messageKey MessageKey) acltypes.MessageDependencyMapping {
	return acltypes.MessageDependencyMapping{
		MessageKey:     string(messageKey),
		DynamicEnabled: false,
		AccessOps: []acltypes.AccessOperation{
			{AccessType: acltypes.AccessType_UNKNOWN, ResourceType: acltypes.ResourceType_ANY, IdentifierTemplate: "*"},
			{AccessType: acltypes.AccessType_COMMIT, ResourceType: acltypes.ResourceType_ANY, IdentifierTemplate: "*"},
		},
	}
}

func DefaultMessageDependencyMapping() []acltypes.MessageDependencyMapping {
	return []acltypes.MessageDependencyMapping{
		SynchronousMessageDependencyMapping(""),
	}
}

func DefaultWasmDependencyMappings() []acltypes.WasmFunctionDependencyMapping {
	return []acltypes.WasmFunctionDependencyMapping{}
}

func ValidateWasmFunctionDependencyMapping(mapping acltypes.WasmFunctionDependencyMapping) error {
	lastAccessOp := mapping.AccessOps[len(mapping.AccessOps)-1]
	if lastAccessOp.AccessType != acltypes.AccessType_COMMIT {
		return ErrNoCommitAccessOp
	}
	return nil
}