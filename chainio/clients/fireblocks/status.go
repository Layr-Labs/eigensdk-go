package fireblocks

type TxStatus string

// statuses for transactions
// ref: https://developers.fireblocks.com/reference/primary-transaction-statuses
const (
	Submitted            TxStatus = "SUBMITTED"
	PendingScreening     TxStatus = "PENDING_AML_SCREENING"
	PendingAuthorization TxStatus = "PENDING_AUTHORIZATION"
	Queued               TxStatus = "QUEUED"
	PendingSignature     TxStatus = "PENDING_SIGNATURE"
	PendingEmailApproval TxStatus = "PENDING_3RD_PARTY_MANUAL_APPROVAL"
	Pending3rdParty      TxStatus = "PENDING_3RD_PARTY"
	Broadcasting         TxStatus = "BROADCASTING"
	Confirming           TxStatus = "CONFIRMING"
	Completed            TxStatus = "COMPLETED"
	Cancelling           TxStatus = "CANCELLING"
	Cancelled            TxStatus = "CANCELLED"
	Blocked              TxStatus = "BLOCKED"
	Rejected             TxStatus = "REJECTED"
	Failed               TxStatus = "FAILED"
)
