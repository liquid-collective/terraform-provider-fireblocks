package client

type AssetTypeResponse struct {
	ID              string `json:"id"`              // The ID of the asset
	Name            string `json:"name"`            // The name of the asset
	Type            string `json:"type"`            //	[ ALGO_ASSET, BASE_ASSET, BEP20, COMPOUND, ERC20, FIAT, SOL_ASSET, TRON_TRC20, XLM_ASSET ]
	ContractAddress string `json:"contractAddress"` //	Contract address for ERC-20 smart contracts
	NativeAsset     string `json:"nativeAsset"`     // The name of the native asset
	Decimals        int64  `json:"decimals"`        // The number of digits after the decimal point
}
