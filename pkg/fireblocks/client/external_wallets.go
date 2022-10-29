package client

type CreateExternalWalletMsg struct {
	Name          string `json:"name"`          // Name of the external wallet container
	CustomerRefID string `json:"customerRefId"` // The ID for AML providers to associate the owner of funds with transactions
}

type ExternalWallet struct {
	ID            string                `json:"id"`            // The ID of the Unmanaged Wallet
	Name          string                `json:"name"`          // Name of the external wallet container
	CustomerRefID string                `json:"customerRefId"` // The ID for AML providers to associate the owner of funds with transactions
	Assets        []ExternalWalletAsset `json:"assets"`        // Array of the assets available in the external wallet
}

type AddExternalWalletAssetMsg struct {
	Address string `json:"address"` //	The wallet's address or, for EOS wallets, the account name
	Tag     string `json:"tag"`     //	For XRP wallets, the destination tag; for EOS/XLM, the memo; for the fiat providers (Signet by Signature, SEN by Silvergate, BLINC by BCB Group), the Bank Transfer Description
}

type ExternalWalletAsset struct {
	ID             string                    `json:"id"`             // The ID of the asset
	Status         ConfigChangeRequestStatus `json:"status"`         // Status of the external wallet
	ActivationTime string                    `json:"activationTime"` // The time the wallet will be activated in case wallets activation posponed according to workspace definition
	Address        string                    `json:"address"`        // The address of the wallet
	Tag            string                    `json:"tag"`            // Destination tag (for XRP, used as memo for EOS/XLM) of the contract wallet.
}
