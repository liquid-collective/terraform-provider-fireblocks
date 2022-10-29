package client

import "github.com/shopspring/decimal"

type CreateTransactionMsg struct {
	AssetID            string                        `json:"assetId"`
	Source             TransferPeerPath              `json:"source"`                       // source of the transaction
	Destination        DestinationTransferPeerPath   `json:"destination"`                  // Destination of the transaction
	Destinations       []DestinationTransferPeerPath `json:"destinations,omitempty"`       // Destination of the transaction
	Amount             string                        `json:"amount"`                       // If the transfer is a withdrawal from an exchange, the actual amount that was requested to be transferred. Otherwise, the requested amount
	TreatAsGrossAmount bool                          `json:"treatAsGrossAmount,omitempty"` // For outgoing transactions, if true, the network fee is deducted from the requested amount
	Fee                string                        `json:"fee,omitempty"`                // The total fee deducted by the exchange from the actual requested amount (serviceFee = amount - netAmount)
	GasPrice           string                        `json:"gasPrice,omitempty"`           // For ETH-based assets only this will be used instead of the fee property, value is in Gwei
	GasLimit           string                        `json:"gasLimit,omitempty"`           // For ETH-based assets only
	NetworkFee         string                        `json:"networkFee,omitempty"`         // The fee paid to the network
	PriorityFee        string                        `json:"priorityFee,omitempty"`
	FeeLevel           string                        `json:"feeLevel,omitempty"`
	MaxFee             string                        `json:"maxFee,omitempty"`
	FailOnLowFee       bool                          `json:"failOnLowFee,omitempty"`
	ForceSweep         bool                          `json:"forceSweep"`                // For "DOT" transactions only, "false" by default, if set to "true" Fireblocks will allow emptying the DOT wallet.
	Note               string                        `json:"note,omitempty"`            // Customer note of the transaction
	Operation          TransactionOperation          `json:"operation"`                 // Default operation is "TRANSFER"
	CustomerRefID      string                        `json:"customerRefID,omitempty"`   // The ID for AML providers to associate the owner of funds with transactions
	ReplacedTxHash     string                        `json:"replacedTxHash,omitempty"`  //
	ExtraParameters    ExtraParameters               `json:"extraParameters,omitempty"` // Protocol / operation specific parameters.
	ExternalTxID       string                        `json:"externalTxID,omitempty"`
}

type CreateTransactionRespMsg struct {
	ID     string            `json:"id"`
	Status TransactionStatus `json:"status"`
}

type TransactionMsg struct {
	ID                            string                   `json:"id"` // ID of the transaction
	AssetID                       string                   `json:"assetId"`
	Source                        TransferPeerPathResponse `json:"source"`                        // source of the transaction
	Destination                   TransferPeerPathResponse `json:"destination"`                   // Destination of the transaction
	RequestedAmount               decimal.Decimal          `json:"requestedAmount"`               // the amount requested by the user
	AmountInfo                    AmountInfo               `json:"amountInfo"`                    // Details of the transaction's amount in string format
	FeeInfo                       FeeInfo                  `json:"feeInfo"`                       // Details of the transaction's fee in string format
	Amount                        decimal.Decimal          `json:"amount"`                        // If the transfer is a withdrawal from an exchange, the actual amount that was requested to be transferred. Otherwise, the requested amount
	NetAmount                     decimal.Decimal          `json:"netAmount"`                     // The net amount of the transaction, after fee deduction
	AmountUSD                     decimal.Decimal          `json:"amountUSD"`                     // The USD value of the requested amount
	ServiceFee                    decimal.Decimal          `json:"serviceFee"`                    // The total fee deducted by the exchange from the actual requested amount (serviceFee = amount - netAmount)
	TreatAsGrossAmount            bool                     `json:"treatAsGrossAmount"`            // For outgoing transactions, if true, the network fee is deducted from the requested amount
	NetworkFee                    decimal.Decimal          `json:"networkFee"`                    // The fee paid to the network
	CreatedAt                     int64                    `json:"createdAt"`                     // Unix timestamp
	LastUpdated                   int64                    `json:"lastUpdated"`                   // Unix timestamp
	Status                        TransactionStatus        `json:"status"`                        // The current status of the transaction
	TxHash                        string                   `json:"txHash"`                        // Blockchain hash of the transaction
	Index                         int64                    `json:"index"`                         // Blockchain hash of the transaction
	SubStatus                     TransactionSubStatus     `json:"subStatus"`                     // More detailed status of the transaction
	SourceAddress                 string                   `json:"sourceAddress"`                 // For account based assets only, the source address of the transaction
	DestinationAddress            string                   `json:"destinationAddress"`            // Address where the asset were transferred
	DestinationAddressDescription string                   `json:"destinationAddressDescription"` // Description of the address
	DestinationTag                string                   `json:"destinationTag"`                // Destination tag (for XRP, used as memo for EOS/XLM) or Bank Transfer Description for Signet/SEN
	SignedBy                      []string                 `json:"signedBy"`                      // Signers of the transaction
	CreatedBy                     string                   `json:"createdBy"`                     // Initiator of the transaction
	RejectedBy                    string                   `json:"rejectedBy"`                    // User ID of the user that rejected the transaction (in case it was rejected)
	AddressType                   string                   `json:"addressType"`                   // [ ONE_TIME, WHITELISTED ]
	Note                          string                   `json:"note"`                          // Customer note of the transaction
	ExchangeTxID                  string                   `json:"exchangeTxId"`                  // If the transaction originated from an exchange, this is the exchange tx ID
	FeeCurrency                   string                   `json:"feeCurrency"`                   // The asset which was taken to pay the fee (ETH for ERC-20 tokens, BTC for Tether Omni)
	Operation                     TransactionOperation     `json:"operation"`                     // Default operation is "TRANSFER"
	AmlScreeningResult            AmlScreeningResult       `json:"amlScreeningResult"`            // The result of the AML screening
	CustomerRefID                 string                   `json:"customerRefId"`                 // The ID for AML providers to associate the owner of funds with transactions
	NumberOfConfirmations         int                      `json:"numOfConfirmations"`            // The number of confirmations of the transaction. The number will increase until the transaction will be considered completed according to the confirmation policy.
	NetworkRecords                []NetworkRecord          `json:"networkRecords"`                // Transaction on the Fireblocks platform can aggregate several blockchain transactions, in such a case these records specify all the transactions that took place on the blockchain.
	ReplacedTxHash                string                   `json:"replacedTxHash"`                // In case of an RBF transaction, the hash of the dropped transaction
	ExternalTxID                  string                   `json:"externalTxId"`                  // Unique transaction ID provided by the user
	Destinations                  []DestinationsResponse   `json:"destinations"`                  // For UTXO based assets, all outputs specified here
	BlockInfo                     BlockInfo                `json:"blockInfo"`                     // The information of the block that this transaction was mined in, the blocks's hash and height
	AuthorizationInfo             AuthorizationInfo        `json:"authorizationInfo"`             // The information about Transaction Authorization Policy (TAP).
	SignedMessages                []SignedMessage          `json:"signedMessages"`                // A list of signed messages returned for raw signing
	ExtraParameters               map[string]interface{}   `json:"extraParameters"`               // Protocol / operation specific parameters.
}

type TransactionOperation string

const (
	TransactionOperationBurn                TransactionOperation = "BURN"
	TransactionOperationContractCall        TransactionOperation = "CONTRACT_CALL"
	TransactionOperationMint                TransactionOperation = "MINT"
	TransactionOperationRaw                 TransactionOperation = "RAW"
	TransactionOperationRedeem              TransactionOperation = "REDEEM"
	TransactionOperationRedeemFromCoumpound TransactionOperation = "REDEEM_FROM_COMPOUND"
	TransactionOperationSupplyFromCoumpound TransactionOperation = "SUPPLY_TO_COMPOUND"
	TransactionOperationTransfer            TransactionOperation = "TRANSFER"
	TransactionOperationTypedMessage        TransactionOperation = "TYPED_MESSAGE"
)

type TransactionStatus string

const (
	TransactionStatusSubmitted               TransactionStatus = "SUBMITTED"
	TransactionStatusQueued                  TransactionStatus = "QUEUED"
	TransactionPendingSignature              TransactionStatus = "PENDING_SIGNATURE"
	TransactionPendingAuthorization          TransactionStatus = "PENDING_AUTHORIZATION"
	TransactionPending3rdPartyManualApproval TransactionStatus = "PENDING_3RD_PARTY_MANUAL_APPROVAL"
	TransactionPending3rdParty               TransactionStatus = "PENDING_3RD_PARTY"
	TransactionPending                       TransactionStatus = "PENDING" // Deprecated
	TransactionBroadcasting                  TransactionStatus = "BROADCASTING"
	TransactionConfirming                    TransactionStatus = "CONFIRMING"
	TransactionConfirmed                     TransactionStatus = "CONFIRMED" // Deprecated
	TransactionCompleted                     TransactionStatus = "COMPLETED"
	TransactionPendingAmlCheckup             TransactionStatus = "PENDING_AML_CHECKUP"
	TransactionPartiallyCompleted            TransactionStatus = "PARTIALLY_COMPLETED"
	TransactionCancelling                    TransactionStatus = "CANCELING"
	TransactionCancelled                     TransactionStatus = "CANCELED"
	TransactionRejected                      TransactionStatus = "REJECTED"
	TransactionFailed                        TransactionStatus = "FAILED"
	TransactionTimeout                       TransactionStatus = "TIMEOUT"
	TransactionBlocked                                         = "BLOCKED"
)

type TransactionSubStatus string

const (
	InsufficientFunds               TransactionSubStatus = "INSUFFICIENT_FUNDS"
	AmountTooSmall                  TransactionSubStatus = "AMOUNT_TOO_SMALL"
	UnsupportedAsset                                     = "UNSUPPORTED_ASSET"
	UnauthorisedMissingPermission   TransactionSubStatus = "UNAUTHORIZED__MISSING_PERMISSION"
	InvalidSignature                TransactionSubStatus = "INVALID_SIGNATURE"
	APIInvalidSignature             TransactionSubStatus = "API_INVALID_SIGNATURE"
	UnauthorisedMissingCredentials  TransactionSubStatus = "UNAUTHORIZED__MISSING_CREDENTIALS"
	UnauthorisedUser                TransactionSubStatus = "UNAUTHORIZED__USER"
	UnauthorisedDevice              TransactionSubStatus = "UNAUTHORIZED__DEVICE"
	InvalidUnmanagedWallet          TransactionSubStatus = "INVALID_UNMANAGED_WALLET"
	InvalidExchangeAccount          TransactionSubStatus = "INVALID_EXCHANGE_ACCOUNT"
	InsufficientFundsForFee         TransactionSubStatus = "INSUFFICIENT_FUNDS_FOR_FEE"
	InvalidAddress                  TransactionSubStatus = "INVALID_ADDRESS"
	WithdrawLimit                   TransactionSubStatus = "WITHDRAW_LIMIT"
	APICallLimit                    TransactionSubStatus = "API_CALL_LIMIT"
	AddressNotWhitelisted           TransactionSubStatus = "ADDRESS_NOT_WHITELISTED"
	TIMEOUT                         TransactionSubStatus = "TIMEOUT"
	ConnectivityError               TransactionSubStatus = "CONNECTIVITY_ERROR"
	ThirdPartyInternalError         TransactionSubStatus = "THIRD_PARTY_INTERNAL_ERROR"
	CancelledExternally             TransactionSubStatus = "CANCELED_EXTERNALLY"
	InvalidThirdPartyResponse       TransactionSubStatus = "INVALID_THIRD_PARTY_RESPONSE"
	VaultWalletNotReady             TransactionSubStatus = "VAULT_WALLET_NOT_READY"
	MissingDepositAddress           TransactionSubStatus = "MISSING_DEPOSIT_ADDRESS"
	OneTimeAddressDisabled          TransactionSubStatus = "ONE_TIME_ADDRESS_DISABLED"
	InternalError                   TransactionSubStatus = "INTERNAL_ERROR"
	UnknownError                    TransactionSubStatus = "UNKNOWN_ERROR"
	AuthorizerNotFound              TransactionSubStatus = "AUTHORIZER_NOT_FOUND"
	InsufficientReservedFunding     TransactionSubStatus = "INSUFFICIENT_RESERVED_FUNDING"
	ManualDepositAddressRequired    TransactionSubStatus = "MANUAL_DEPOSIT_ADDRESS_REQUIRED"
	InvalidFee                      TransactionSubStatus = "INVALID_FEE"
	ErrorUnsupportedTransactionType TransactionSubStatus = "ERROR_UNSUPPORTED_TRANSACTION_TYPE"
	UnsupportedOperation            TransactionSubStatus = "UNSUPPORTED_OPERATION"
	T3rdPartyProcessing             TransactionSubStatus = "3RD_PARTY_PROCESSING"
	PendingBlockchainConfirmations  TransactionSubStatus = "PENDING_BLOCKCHAIN_CONFIRMATIONS"
	T3rdPartyConfirming             TransactionSubStatus = "3RD_PARTY_CONFIRMING"
	CONFIRMED                       TransactionSubStatus = "CONFIRMED"
	T3rdPartyCompleted              TransactionSubStatus = "3RD_PARTY_COMPLETED"
	RejectedByUser                  TransactionSubStatus = "REJECTED_BY_USER"
	CancelledByUser                 TransactionSubStatus = "CANCELED_BY_USER"
	T3rdPartyCancelled              TransactionSubStatus = "3RD_PARTY_CANCELED"
	T3rdPartyRejected               TransactionSubStatus = "3RD_PARTY_REJECTED"
	AmlScreeningRejected            TransactionSubStatus = "AML_SCREENING_REJECTED"
	BlockedByPolicy                 TransactionSubStatus = "BLOCKED_BY_POLICY"
	FailedAmlScreening              TransactionSubStatus = "FAILED_AML_SCREENING"
	PartiallyFailed                 TransactionSubStatus = "PARTIALLY_FAILED"
	T3rdPartyFailed                 TransactionSubStatus = "3RD_PARTY_FAILED"
	DroppedByBlockchain             TransactionSubStatus = "DROPPED_BY_BLOCKCHAIN"
	TooManyInputs                   TransactionSubStatus = "TOO_MANY_INPUTS"
	SigningError                    TransactionSubStatus = "SIGNING_ERROR"
	InvalidFeeParams                TransactionSubStatus = "INVALID_FEE_PARAMS"
	MissingTagOrMemo                TransactionSubStatus = "MISSING_TAG_OR_MEMO"
	GasLimitTooLow                  TransactionSubStatus = "GAS_LIMIT_TOO_LOW"
	MaxFeeExceeded                  TransactionSubStatus = "MAX_FEE_EXCEEDED"
	ActualFeeTooHigh                TransactionSubStatus = "ACTUAL_FEE_TOO_HIGH"
	InvalidContractCallData         TransactionSubStatus = "INVALID_CONTRACT_CALL_DATA"
	InvalidNonceTooLow              TransactionSubStatus = "INVALID_NONCE_TOO_LOW"
	InvalidNonceTooHigh             TransactionSubStatus = "INVALID_NONCE_TOO_HIGH"
	InvalidNonceForRbf              TransactionSubStatus = "INVALID_NONCE_FOR_RBF"
	FailOnLowFee                    TransactionSubStatus = "FAIL_ON_LOW_FEE"
	TooLongMempoolChain             TransactionSubStatus = "TOO_LONG_MEMPOOL_CHAIN"
	TxOutdated                      TransactionSubStatus = "TX_OUTDATED"
	IncompleteUserSetup             TransactionSubStatus = "INCOMPLETE_USER_SETUP"
	SignerNotFound                  TransactionSubStatus = "SIGNER_NOT_FOUND"
	InvalidTagOrMemo                TransactionSubStatus = "INVALID_TAG_OR_MEMO"
	ZeroBalanceInPermanentAddress   TransactionSubStatus = "ZERO_BALANCE_IN_PERMANENT_ADDRESS"
	NeedMoreToCreateDestination     TransactionSubStatus = "NEED_MORE_TO_CREATE_DESTINATION"
	NonExistingAccountName          TransactionSubStatus = "NON_EXISTING_ACCOUNT_NAME"
	EnvUnsupportedAsset             TransactionSubStatus = "ENV_UNSUPPORTED_ASSET"
)

type BlockInfo struct {
	BlockHeight string `json:"blockHeight"`
	BlockHash   string `json:"blockHash"`
}

type TransactionFee struct {
	FeePerByte string `json:"feePerByte"` // [optional] For UTXOs,
	GasPrice   string `json:"gasPrice"`   // [optional] For Ethereum assets (ETH and Tokens)
	GasLimit   string `json:"gasLimit"`   // [optional] For Ethereum assets (ETH and Tokens), the limit for how much can be used
	NetworkFee string `json:"networkFee"` // [optional] Transaction fee
}

type EstimatedTransactionFeeResponse struct {
	Low    TransactionFee `json:"low"`    // Transactions with this fee will probably take longer to be mined
	Medium TransactionFee `json:"medium"` // Average transactions fee
	High   TransactionFee `json:"high"`   // Transactions with this fee should be mined the fastest
}

type AddressStatus struct {
	IsValid     bool `json:"isValid"`
	IsActive    bool `json:"isActive"`
	RequiresTag bool `json:"requiresTag"`
}

type SignedMessage struct {
	Content        string                 `json:"content"`        // The message for signing (hex-formatted)
	Algorithm      SigningAlgorithm       `json:"algorithm"`      // The algorithm that was used for signing, one of the SigningAlgorithms
	DerivationPath string                 `json:"derivationPath"` // BIP32 derivation path of the signing key. E.g. [44,0,46,0,0]
	Signature      map[string]interface{} `json:"signature"`      // The message signature
	PublicKey      string                 `json:"publicKey"`      // Signature's public key that can be used for verification.
}

type DestinationsResponse struct {
	Amount                        decimal.Decimal          `json:"amount"`                        // The amount to be sent to this destination
	Destination                   TransferPeerPathResponse `json:"destination"`                   // Destination of the transaction
	AmountUSD                     decimal.Decimal          `json:"amountUSD"`                     // The USD value of the requested amount
	DestinationAddress            string                   `json:"destinationAddress"`            // Address where the asset were transferred
	DestinationAddressDescription string                   `json:"destinationAddressDescription"` // Description of the address
	AmlScreeningResult            AmlScreeningResult       `json:"amlScreeningResult"`            // The result of the AML screening
	CustomerRefID                 string                   `json:"customerRefId"`                 // The ID for AML providers to associate the owner of funds with transactions

}

type NetworkRecord struct {
	Source             TransferPeerPathResponse `json:"source"`             // Source of the transaction
	Destination        TransferPeerPathResponse `json:"destination"`        // Destination of the transaction
	TxHash             string                   `json:"txHash"`             // Blockchain hash of the transaction
	NetworkFee         decimal.Decimal          `json:"networkFee"`         // The fee paid to the network
	AssetID            string                   `json:"assetId"`            // transaction asset
	NetAmount          decimal.Decimal          `json:"netAmount"`          // The net amount of the transaction, after fee deduction
	Status             NetworkStatus            `json:"status"`             // Status of the blockchain transaction
	OpType             string                   `json:"type"`               // Type of the operation
	DestinationAddress string                   `json:"destinationAddress"` // Destination address
	SourceAddress      string                   `json:"sourceAddress"`      // For account based assets only, the source address of the transaction

}

type NetworkStatus string

const (
	NetworkStatusDropped      NetworkStatus = "DROPPED"
	NetworkStatusBroadcasting NetworkStatus = "BROADCASTING"
	NetworkStatusConfirming   NetworkStatus = "CONFIRMING"
	NetworkStatusFailed       NetworkStatus = "FAILED"
	NetworkStatusConfirmed    NetworkStatus = "CONFIRMED"
)

type AmlScreeningResult struct {
	Provider string `json:"provider"` // The AML service provider
	Payload  string `json:"payload"`  // The response of the AML service provider
}

type AmountInfo struct {
	Amount          string `json:"amount"`          // If the transfer is a withdrawal from an exchange, the actual amount that was requested to be transferred. Otherwise, the requested amount
	RequestedAmount string `json:"requestedAmount"` // The amount requested by the user
	NetAmount       string `json:"netAmount"`       // The net amount of the transaction, after fee deduction
	AmountUSD       string `json:"amountUSD"`       // The USD value of the requested amount
}

type FeeInfo struct {
	NetworkFee string `json:"networkFee"` // The fee paid to the network
	ServiceFee string `json:"serviceFee"` // The total fee deducted by the exchange from the actual requested amount (serviceFee = amount - netAmount)
}

type SigningAlgorithm string

const (
	MpcEcdsaSecp256k1 SigningAlgorithm = "MPC_ECDSA_SECP256K1"
	MpcEddsaEd25519   SigningAlgorithm = "MPC_EDDSA_ED25519"
)

type OneTimeAddress struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
}

type TransferPeerPathResponse struct {
	TransferType string `json:"type"` // [ PeerTypeVaultAccount, EXCHANGE_ACCOUNT, INTERNAL_WALLET, EXTERNAL_WALLET, ONE_TIME_ADDRESS, NETWORK_CONNECTION, FIAT_ACCOUNT, COMPOUND ]
	ID           string `json:"id"`   // The ID of the exchange account to return
	Name         string `json:"name"` // The name of the exchange account
	Subtype      string `json:"subType"`
}

type TransferPeerPath struct {
	ID   string   `json:"id"` // The peer ID
	Type PeerType `json:"type"`
}

type DestinationTransferPeerPath struct {
	ID             string         `json:"id"` // The peer ID (not needed for ONE_TIME_ADDRESS)
	Type           PeerType       `json:"type"`
	OneTimeAddress OneTimeAddress `json:"oneTimeAddress"`
}

type ExtraParameters struct {
	ContractCallData string `json:"contractCallData"`
}

type PeerType string

const (
	PeerTypeVaultAccount      PeerType = "VAULT_ACCOUNT"
	PeerTypeExchangeAccount   PeerType = "EXCHANGE_ACCOUNT"
	PeerTypeInternalWallet    PeerType = "INTERNAL_WALLET"
	PeerTypeExternalWallet    PeerType = "EXTERNAL_WALLET"
	PeerTypeOneTimeAddress    PeerType = "ONE_TIME_ADDRESS"
	PeerTypeFiatAccount       PeerType = "FIAT_ACCOUNT"
	PeerTypeNetworkConnection PeerType = "NETWORK_CONNECTION"
	PeerTypeCompound          PeerType = "COMPOUND"
)

type AuthorizationInfo struct {
	AllowOperatorAsAuthorizer bool                 `json:"allowOperatorAsAuthorizer"` // Set to "true" if the intiator of the transaction can be one of the approvers
	Logic                     Logic                `json:"logic"`                     // "AND" or "OR", this is the logic that is applied between the different authorization groups listed below
	Groups                    []AuthorizationGroup `json:"groups"`                    // The list of authorization groups and users that are required to approve this transaction. The logic applied between the different groups is the “logic” field above. Each element in the response is the user ID (the can found see via the users endpoint) and theirApprovalStatus
}

type Logic string

const (
	LogicOr  Logic = "OR"
	LogicAnd Logic = "AND"
)

type AuthorizationGroup struct {
	Threshold int64                     `json:"th"`    // The threshold of required approvers in this authorization group
	Users     map[string]ApprovalStatus `json:"users"` // The list of users that the threshold number is applied to for transaction approval. Each user in the response is a "key:value" where the key is the user ID (the can found see via the users endpoint) and the value is the theirApprovalStatus.
}

type ApprovalStatus string

const (
	ApprovalStatuPendingAuthorization ApprovalStatus = "PENDING_AUTHORIZATION"
	ApprovalStatuApproved             ApprovalStatus = "APPROVED"
	ApprovalStatuRejected             ApprovalStatus = "REJECTED"
	ApprovalStatuNA                   ApprovalStatus = "NA"
)
