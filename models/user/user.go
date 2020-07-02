package user

type User struct {
	id   int
	name string
}

func (u *User) GetTableName() string {
	return "user_tbl"
}
