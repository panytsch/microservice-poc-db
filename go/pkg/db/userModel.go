package db

type User struct {
	Model
	Name     string
	Password string
	Balance  int
	CCNumber string `gorm:"column:cc_number"`
}

func (u *User) Create() {
	DB.Omit("created_at", "updated_at", "deleted_at").Create(u)
}

func (*User) TableName() string {
	return "users"
}

func (u *User) FindByNameAndPass(name string, pass string) *User {
	u.Name = name
	u.Password = pass
	DB.Unscoped().Where(u).Find(u)
	return u
}
