package entity

type User struct {
	Id   uint `gorm:"primaryKey"`
	Name string
}
