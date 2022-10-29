package client

type UnmanagedContract struct {
	ID     string          `json:"id"`     // The ID of the unmanaged contract wallet
	Name   string          `json:"name"`   // Name of the contract wallet container
	Assets []ContractAsset `json:"assets"` // Array of ContractAsset	Array of the assets available in the unmanaged contract wallet
}

type ContractAsset struct {
	ID             string                    `json:"id"`             // The ID of the contract wallet
	Balance        string                    `json:"balance"`        // The balance of the contract wallet
	LockedAmount   string                    `json:"lockedAmount"`   // 	Locked amount in the contract wallet
	Status         ConfigChangeRequestStatus `json:"status"`         // 	Status of the contract wallet
	ActivationTime string                    `json:"activationTime"` // 	The time the contract wallet will be activated if case wallets activation is posponed according to the workspace definition
	Address        string                    `json:"address"`        // 	The address of the contract wallet
	Tag            string                    `json:"tag"`            // 	Destination tag (for XRP, used as memo for EOS/XLM) of the contract wallet.
}

type ConfigChangeRequestStatus string

const (
	ConfigChangeRequestStatusWaitingForApproval ConfigChangeRequestStatus = "WAITING_FOR_APPROVAL"
	ConfigChangeRequestStatusApproved           ConfigChangeRequestStatus = "APPROVED"
	ConfigChangeRequestStatusCancelled          ConfigChangeRequestStatus = "CANCELED"
	ConfigChangeRequestStatusRejected           ConfigChangeRequestStatus = "REJECTED"
	ConfigChangeRequestStatusFailed             ConfigChangeRequestStatus = "FAILED"
)
