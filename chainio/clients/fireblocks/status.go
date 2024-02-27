package fireblocks

type Status string

// statuses for transactions
// ref: https://developers.fireblocks.com/reference/primary-transaction-statuses
const (
	Submitted            Status = "SUBMITTED"
	PendingScreening     Status = "PENDING_AML_SCREENING"
	PendingAuthorization Status = "PENDING_AUTHORIZATION"
	Queued               Status = "QUEUED"
	PendingSignature     Status = "PENDING_SIGNATURE"
	PendingEmailApproval Status = "PENDING_3RD_PARTY_MANUAL_APPROVAL"
	Pending3rdParty      Status = "PENDING_3RD_PARTY"
	Broadcasting         Status = "BROADCASTING"
	Confirming           Status = "CONFIRMING"
	Completed            Status = "COMPLETED"
	Cancelling           Status = "CANCELLING"
	Cancelled            Status = "CANCELLED"
	Blocked              Status = "BLOCKED"
	Rejected             Status = "REJECTED"
	Failed               Status = "FAILED"
)
