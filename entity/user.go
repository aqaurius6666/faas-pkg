package entity

type User struct {
	ID        int64  `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	FirstName string `json:"first_name" gorm:"column:first_name;type:varchar(255)"`
	LastName  string `json:"last_name" gorm:"column:last_name;type:varchar(255)"`
	Username  string `json:"username" gorm:"column:username;type:varchar(64);idx_user_username;unique"`
	Password  string `json:"password" gorm:"column:password;type:varchar(255)"`
}

type UserQuery struct {
	User
}
