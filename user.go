package koifer

// User is a couple user/password
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserRepository is anything that can return a User from a username.
type UserRepository interface {
	// GetUserByName returns a pointer to the user if it exist and a nil error.
	// If the user does not exists, it returns nil and a nil error
	// An errror is returns in case of problem accessing the data
	GetUserByName(name string) (*User, error)
}
