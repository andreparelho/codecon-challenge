package repository

type Users interface {
	SaveUsers(users map[string]User)
}

type user struct {
	Users map[string]User
}

func NewUsers(users map[string]User) *user {
	return &user{
		Users: users,
	}
}

func (u *user) SaveUsers(users []User) {
	for _, user := range users {
		u.Users[user.Id] = user
	}
}
