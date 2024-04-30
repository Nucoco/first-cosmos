package types

const (
	// ModuleName defines the module name
	ModuleName = "firstcosmos"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_firstcosmos"
)

var (
	ParamsKey = []byte("p_firstcosmos")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
