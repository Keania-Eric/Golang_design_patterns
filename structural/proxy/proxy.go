package proxy

import (
	"fmt"
)

type Repository interface {
	Find(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

func (t *UserList) Find(id int32) (User, error) {
	for i := 0; i < len(*t); i++ { // notice pointer deferencing in go
		if (*t)[i].ID == id {
			return (*t)[i], nil
		}
	}
	return User{}, fmt.Errorf("user %d could not be found", id)
}

func (t *UserList) addUser(user User) {
	*t = append(*t, user)
}

type UserListProxy struct {
	AppDatabase         UserList
	StackCache          UserList
	StackCapacity       int
	LastSearchUsedCache bool
}

func (p *UserListProxy) addUserToStack(user User) {
	if len(p.StackCache) >= p.StackCapacity {
		p.StackCache = append(p.StackCache[1:], user) // more reading of slice
	} else {
		p.StackCache.addUser(user)
	}
}

func (p *UserListProxy) Find(id int32) (User, error) {
	user, err := p.StackCache.Find(id)

	if err == nil {
		fmt.Println("Returning user from cache")
		p.LastSearchUsedCache = true
		return user, nil
	}

	user, err = p.AppDatabase.Find(id)
	if err != nil {
		return User{}, err
	}

	p.addUserToStack(user)

	fmt.Println("Returning user from database")
	p.LastSearchUsedCache = false
	return user, nil
}
