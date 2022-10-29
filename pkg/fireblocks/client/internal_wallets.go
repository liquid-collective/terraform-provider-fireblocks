package client

type CreateInternalWalletMsg struct {
	Name          string `json:"name"`          // The wallet container's display name
	CustomerRefID string `json:"customerRefId"` // The ID for AML providers to associate the owner of funds with transactions
}

type UnmanagedWallet struct {
	ID            string        `json:"id"`            //	The ID of the Unmanaged Wallet
	Name          string        `json:"name"`          // Name of the Wallet Container
	CustomerRefID string        `json:"customerRefId"` //	[optional] The ID for AML providers to associate the owner of funds with transactions
	Assets        []WalletAsset `json:"assets"`        // Array of the assets available in the unmanaged wallet
}

type WalletAsset struct {
	ID             string                    `json:"id"`             //	The ID of the asset
	Balance        string                    `json:"balance"`        //	The balance of the wallet. Values are returned according to balance decimal precision.
	LockedAmount   string                    `json:"lockedAmount"`   // Locked amount in the wallet. Values are returned according to balance decimal precision.
	Status         ConfigChangeRequestStatus `json:"status"`         //	Status of the Internal Wallet
	ActivationTime string                    `json:"activationTime"` //	The time the wallet will be activated in case wallets activation posponed according to workspace definition
	Address        string                    `json:"address"`        // The address of the wallet
	Tag            string                    `json:"tag"`            // Destination tag (for XRP, used as memo for EOS/XLM) of the wallet.
}
