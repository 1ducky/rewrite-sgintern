package assets

type AssetStatus string

const (
	AssetStatusTemp    AssetStatus = "temp"
	AssetStatusActive  AssetStatus = "active"
	AssetStatusOrphan  AssetStatus = "orphan"
	AssetStatusDeleted AssetStatus = "deleted"
)
