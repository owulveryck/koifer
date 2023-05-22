package memory

import (
	"github.com/owulveryck/koifer"
)

type DB struct {
	users map[string]*koifer.User
}

func NewDB() *DB {
	return &DB{
		users: make(map[string]*koifer.User),
	}
}

func (d *DB) UpsertUser(user, pass string) {
	d.users[user] = &koifer.User{
		Name:     user,
		Password: pass,
	}
}

func (d *DB) GetUserByName(name string) (*koifer.User, error) {
	user, ok := d.users[name]
	if !ok {
		return nil, nil
	}
	return user, nil
}
