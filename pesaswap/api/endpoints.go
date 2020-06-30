package api

var (
	MpesaPesaSwap             = "api/pesaswap/mobile-payment"
	GetTransactionRecord      = "api/get-transaction-record"
	FilterByTransactionId     = "api/get-transaction-record/filter-by-transaction-id"
	ViewAllTransactionsRecord = "api/view-all-transactions-record"
	IveriTransaction          = "api/view/iveri-transaction"
	FilterByApplicationId     = "api/view/iveri-transaction/filter-by-application-id"
	CreateUser                = "api/pesaswap/create/customer"
	UpdateUser                = "api/pesaswap/update/customer"
	QueryCustomer             = "api/pesaswap/query/customer"
	PaymentRequest            = "api/payment/request"
	CouponPayment             = "api/coupon/payment"
	CouponQueryEndpoint       = "api/coupon/query"
	CardPayment               = "api/pesaswap/credit-card-payment"
	AccountBalance            = "api/account-balance-query"
	FundTransfer              = "api/funds-transfer"
	SendToMpesa               = "api/send-to-mpesa"
	StatusQuery               = "api/status-query"
)
