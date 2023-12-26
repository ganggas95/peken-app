package domain

type User struct {
	Id       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"column:name;not null"`
	Email    string `gorm:"<-;unique;not null"`
	Password string `gorm:"not null" json:"-"`
	Username string `gorm:"<-;unique;not null"`
	BaseModel
}
