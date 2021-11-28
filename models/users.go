package models

type User struct {
	ID        string `bson:"_id"`
	Username  string `bson:"username"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

func (u *User) UsernameAvailable() bool {
	// res, _ := datastore.DB.Users.Find(bson.M{"username": strings.ToLower(u.Username)}).Count()
	res := 1
	return res > 0
}

func (u *User) EmailAvailable() bool {
	// res, _ := datastore.DB.Users.Find(bson.M{"email": strings.ToLower(u.Email)}).Count()
	res := 1
	return res > 0
}

func (u *User) Create() error {
	// return datastore.DB.Users.Insert(u)
	return nil
}
