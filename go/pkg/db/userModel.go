package db

type User struct {
	ID          uint
	Transaction []Transaction `gorm:"foreignkey:UserID"`
	Name        string
	Password    string
}

func (u *User) Create() {
	DB.Create(u)
}

func (*User) TableName() string {
	return "users"
}

func (u *User) FindByNameAndPass(name string, pass string) *User {
	u.Name = name
	u.Password = pass
	DB.Find(u)
	return u
}
