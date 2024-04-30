package types

const (
	// ModuleName defines the module name
	ModuleName = "firstmodule"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_firstmodule"
)

var (
	ParamsKey = []byte("p_firstmodule")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
