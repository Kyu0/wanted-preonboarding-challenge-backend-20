package products

// 예약 상태에 대한 타입
type TransactionStatus string

const (
	SELLING  TransactionStatus = "SELLING"  // 판매 중
	RESERVED TransactionStatus = "RESERVED" // 예약 중
	DONE     TransactionStatus = "DONE"     // 판매 완료
)

func (t TransactionStatus) IsVaild() bool {
	switch t {
	case SELLING, RESERVED, DONE:
		return true
	default:
		return false
	}
}

type Product struct {
	ID     uint              `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	Name   string            `gorm:"type:varchar(64);not null" json:"name"`
	Price  uint32            `json:"price"`
	Status TransactionStatus `gorm:"type:varchar(10);embedded" json:"status" binding:"enum"`
	Amount uint16            `json:"amount"`
	UserId uint              `json:"userId"`
}
