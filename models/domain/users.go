package domain

type User struct {
	Id       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"column:name;not null"`
	Email    string `gorm:"<-;column:email;unique;not null"`
	Password string `gorm:"column:password;not null" json:"-"`
	Username string `gorm:"<-;column:username;unique;not null"`
	Roles    []Role `gorm:"many2many:user_roles;" json:"roles"`
	BaseModel
}
