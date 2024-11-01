package auth

type UserRepository interface {
	FindByEmail() (User, error)
}

type UserManager interface {
	Create(user User) (User, error)
	Delete(user User) error
}
