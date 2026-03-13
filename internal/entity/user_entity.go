package entity

type User struct {
	ID      int    `gorm:"column:id;primaryKey;autoIncrement"`
	Email       string `gorm:"unique;column:email"`
	PhoneNumber string `gorm:"column:phone_number"`
	Role        string `gorm:"column:role"`
	// Token       string `gorm:"column:token"`
	Password    string `gorm:"column:password"`
	CreatedAt   int64  `gorm:"column:created_at"`
}