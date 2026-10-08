package types

type Password struct {
	Id          uint `gorm:"primaryKey;autoIncrement"`
	ServiceName string 
	Password 	string 
	SecLevel 	string 
}
