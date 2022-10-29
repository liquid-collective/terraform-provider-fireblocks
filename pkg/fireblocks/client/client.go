package client

import (
	"context"
)

type UserMsg struct {
	ID        string `json:"id"`        // User ID on the Fireblocks platform
	FirstName string `json:"firstName"` // First name
	LastName  string `json:"lastName"`  // Last name
	Role      string `json:"role"`      // The role of the user in the workspace
	Email     string `json:"email"`     // The email of the user
	Enabled   bool   `json:"enabled"`   // The status of the user in the workspace
}

type Client interface {
	ListUsers(ctx context.Context) ([]*UserMsg, error)

	CreateVaultAccount(ctx context.Context, msg *CreateVaultAccountMsg) (*VaultAccount, error)
	GetVaultAccount(ctx context.Context, vaultAccountID string) (*VaultAccount, error)
	CreateVaultAccountAsset(ctx context.Context, vaultID, assetID string) (*CreateVaultAssetResponse, error)
	UpdateVaultAccount(ctx context.Context, vaultAccountID string, msg *UpdateVaultAccountMsg) error
	HideVaultAccount(ctx context.Context, vaultAccountID string) error
	UnhideVaultAccount(ctx context.Context, vaultAccountID string) error
	GetVaultAccountAssetBalance(ctx context.Context, vaultID, assetID string) (*VaultAsset, error)
	ListVaultAccountAssetAddresses(ctx context.Context, vaultID, assetID string) ([]*VaultAccountAssetAddress, error)

	CreateExternalWallet(ctx context.Context, msg *CreateExternalWalletMsg) (*ExternalWallet, error)
	GetExternalWallet(ctx context.Context, walletID string) (*ExternalWallet, error)
	DeleteExternalWallet(ctx context.Context, walletID string) error
	AddExternalWalletAsset(ctx context.Context, walletID, assetID string, msg *AddExternalWalletAssetMsg) (*ExternalWalletAsset, error)
	GetExternalWalletAsset(ctx context.Context, walletID, assetID string) (*ExternalWalletAsset, error)
	DeleteExternalWalletAsset(ctx context.Context, walletID, assetID string) error

	CreateTransaction(ctx context.Context, msg *CreateTransactionMsg) (*CreateTransactionRespMsg, error)
	GetTransaction(ctx context.Context, txID string) (*TransactionMsg, error)
	CancelTransaction(ctx context.Context, txID string) error
}
