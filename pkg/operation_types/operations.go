package operations

const (
	DEPOSIT  = 1
	WITHDRAW = 2
	TRANSFER = 3
)

var NameByType = map[int]string{
	DEPOSIT:  "DEPOSIT",
	WITHDRAW: "WITHDRAW",
	TRANSFER: "TRANSFER",
}
