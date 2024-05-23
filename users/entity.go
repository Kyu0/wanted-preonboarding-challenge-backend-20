package users

type User struct {
	ID       uint   `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	Username string `gorm:"type:varchar(32);uniqueIndex;not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
}
