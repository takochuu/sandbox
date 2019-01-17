package archtecture

type User struct {
	ID      int
	Name    string
	Address Address
}

type Address string

type UserService interface {
	User(id int) (*User, error)
	Users() ([]*User, error)
	Create(u *User) error
	Delete(id int) error
}

type UserCache struct {
	cache   map[int]*User
	service UserService
}

func NewUserCache(service UserService) *UserCache {
	return &UserCache{
		cache:   make(map[int]*User),
		service: service,
	}
}

func (c *UserCache) User(id int) (*User, error) {
	if u := c.cache[id]; u != nil {
		return u, nil
	}

	u, err := c.service.User(id)
	if err != nil {
		return nil, err
	}

	c.cache[id] = u
	return u, nil
}
