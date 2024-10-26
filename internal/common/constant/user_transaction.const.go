package constant

import "time"

const (
	UserTransactionTypeDebit  = "DEBIT"
	UserTransactionTypeCredit = "CREDIT"

	UserTransactionActionTypeTopUp    = "TOP_UP"
	UserTransactionActionTypePayment  = "PAYMENT"
	UserTransactionActionTypeTransfer = "TRANSFER"

	UserTransactionStatusPending = "PENDING"
	UserTransactionStatusSuccess = "SUCCESS"
	UserTransactionStatusFailed  = "FAILED"
)

var (
	LockUserTransferTimeout = time.Minute * 10
)
