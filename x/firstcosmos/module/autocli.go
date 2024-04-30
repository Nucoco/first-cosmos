package firstcosmos

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "first-cosmos/api/firstcosmos/firstcosmos"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "PeopleAll",
					Use:       "list-people",
					Short:     "List all people",
				},
				{
					RpcMethod:      "People",
					Use:            "show-people [id]",
					Short:          "Shows a people by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "Hello",
					Use:            "hello [from-address] [to-address]",
					Short:          "Send a hello tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "fromAddress"}, {ProtoField: "toAddress"}},
				},
				{
					RpcMethod:      "CreatePeople",
					Use:            "create-people [address] [name]",
					Short:          "Create people",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}, {ProtoField: "name"}},
				},
				{
					RpcMethod:      "UpdatePeople",
					Use:            "update-people [id] [address] [name]",
					Short:          "Update people",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "address"}, {ProtoField: "name"}},
				},
				{
					RpcMethod:      "DeletePeople",
					Use:            "delete-people [id]",
					Short:          "Delete people",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
