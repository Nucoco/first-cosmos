package firstcosmos

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"first-cosmos/testutil/sample"
	firstcosmossimulation "first-cosmos/x/firstcosmos/simulation"
	"first-cosmos/x/firstcosmos/types"
)

// avoid unused import issue
var (
	_ = firstcosmossimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgHello = "op_weight_msg_hello"
	// TODO: Determine the simulation weight value
	defaultWeightMsgHello int = 100

	opWeightMsgCreatePeople = "op_weight_msg_people"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePeople int = 100

	opWeightMsgUpdatePeople = "op_weight_msg_people"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePeople int = 100

	opWeightMsgDeletePeople = "op_weight_msg_people"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePeople int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	firstcosmosGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		PeopleList: []types.People{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		PeopleCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&firstcosmosGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgHello int
	simState.AppParams.GetOrGenerate(opWeightMsgHello, &weightMsgHello, nil,
		func(_ *rand.Rand) {
			weightMsgHello = defaultWeightMsgHello
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgHello,
		firstcosmossimulation.SimulateMsgHello(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreatePeople int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePeople, &weightMsgCreatePeople, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePeople = defaultWeightMsgCreatePeople
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePeople,
		firstcosmossimulation.SimulateMsgCreatePeople(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePeople int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePeople, &weightMsgUpdatePeople, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePeople = defaultWeightMsgUpdatePeople
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePeople,
		firstcosmossimulation.SimulateMsgUpdatePeople(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePeople int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePeople, &weightMsgDeletePeople, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePeople = defaultWeightMsgDeletePeople
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePeople,
		firstcosmossimulation.SimulateMsgDeletePeople(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgHello,
			defaultWeightMsgHello,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				firstcosmossimulation.SimulateMsgHello(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePeople,
			defaultWeightMsgCreatePeople,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				firstcosmossimulation.SimulateMsgCreatePeople(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdatePeople,
			defaultWeightMsgUpdatePeople,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				firstcosmossimulation.SimulateMsgUpdatePeople(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeletePeople,
			defaultWeightMsgDeletePeople,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				firstcosmossimulation.SimulateMsgDeletePeople(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
