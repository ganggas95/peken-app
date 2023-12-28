package domain

type User struct {
	Id       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"column:name;not null"`
	Email    string `gorm:"<-;unique;not null"`
	Password string `gorm:"not null" json:"-"`
	Username string `gorm:"<-;unique;not null"`
	Roles    []Role `gorm:"many2many:user_roles;" json:"roles"`
	BaseModel
}
