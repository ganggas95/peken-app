package domain

type Member struct {
	Name    string `gorm:"column:name;not null"`
	Id      uint   `gorm:"primaryKey;autoIncrement"`
	Address string `gorm:"column:address;not null"`
	User    User   `gorm:"foreignKey:Id"`
	BaseModel
}
