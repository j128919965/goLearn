package po

type User struct {
	Id   int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
	Age  *int
}
