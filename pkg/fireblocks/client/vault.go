package client

type PageInfo struct {
	Paging      Paging `json:"paging"`
	PreviousURL string `json:"previousUrl"`
	NextURL     string `json:"nextUrl"`
}

type Paging struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

type UpdateVaultAccountMsg struct {
	Name string `json:"name"`
}

type CreateVaultAccountMsg struct {
	Name          string `json:"name"`
	HiddenOnUI    bool   `json:"hiddenOnUI"`
	CustomerRefID string `json:"customerRefId"`
	AutoFuel      bool   `json:"autoFuel"`
}

type VaultAccount struct {
	ID            string       `json:"id"`
	Name          string       `json:"name"`
	HiddenOnUI    bool         `json:"hiddenOnUI"`
	CustomerRefID string       `json:"customerRefId"`
	AutoFuel      bool         `json:"autoFuel"`
	Assets        []VaultAsset `json:"assets"`
}

type VaultAsset struct {
	ID                   string `json:"id"`
	Total                string `json:"total"`
	Available            string `json:"available"`
	Pending              string `json:"pending"`
	LockedAmount         string `json:"lockedAmount"`
	TotalStackedCPU      string `json:"totalStackedCPU"`
	TotalStackedNetwork  string `json:"totalStackedNetwork"`
	SelfStackedCPU       string `json:"selfStackedCPU"`
	SelfStakedNetwork    string `json:"selfStakedNetwork"`
	PendingRefundCPU     string `json:"pendingRefundCPU"`
	PendingRefundNetwork string `json:"pendingRefundNetwork"`
}

type VaultAccountAssetAddress struct {
	AssetID           string `json:"assetId"`           // The ID of the asset
	Address           string `json:"address"`           // Address of the asset in a Vault Account, for BTC/LTC the address is in Segwit (Bech32) format, for BCH cash format
	LegacyAddress     string `json:"legacyAddress"`     // For BTC/LTC/BCH the legacy format address
	Description       string `json:"description"`       // Description of the address
	Tag               string `json:"tag"`               // Destination tag for XRP, used as memo for EOS/XLM, for the fiat providers (Signet by Signature, SEN by Silvergate, BLINC by BCB Group), it is the Bank Transfer Description
	Type              string `json:"type"`              // Address type
	Change            string `json:"change"`            // The change address for BTC transactions
	CustomerRefID     string `json:"customerRefId"`     // The ID for AML providers to associate the owner of funds with transactions
	Bip44AddressIndex int64  `json:"bip44AddressIndex"` // The address_index, addressFormat, and enterpriseAddress in the derivation path of this address based on BIP44

}

type CreateVaultAssetResponse struct {
	ID             string `json:"id"`             // The ID of the Vault Account
	Address        string `json:"address"`        // Address of the asset in a Vault Account, for BTC/LTC the address is in Segwit (Bech32) format, cash address format for BCH
	LegacyAddress  string `json:"legacyAddress"`  // Legacy address format for BTC/LTC/BCH
	Tag            string `json:"tag"`            // Destination tag for XRP, used as memo for EOS/XLM
	EosAccountName string `json:"eosAccountName"` // Returned for EOS, the account name
}

type VaultAccountsWithPageInfoMsg struct {
	Accounts []*VaultAccount `json:"accounts"`
	PageInfo
}
