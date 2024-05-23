package products

// 예약 상태에 대한 타입
type TransactionStatus string

const (
	SELLING  TransactionStatus = "SELLING"  // 판매 중
	RESERVED TransactionStatus = "RESERVED" // 예약 중
	DONE     TransactionStatus = "DONE"     // 판매 완료
)

type Product struct {
	ID     uint `gorm:"primaryKey"`
	Name   string
	Price  uint
	Status TransactionStatus
	Amount uint
	UserId uint
}
