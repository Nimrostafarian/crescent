package types

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// TODO: not implemented yet
	return gs.Params.Validate()
}