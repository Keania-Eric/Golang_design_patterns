package proxy

import (
	"math/rand"
	"testing"
)

func TestUserListProxy(t *testing.T) {
	database := UserList{}

	rand.Seed(2342342)
	for i := 0; i < 100000; i++ {
		n := rand.Int31()
		database = append(database, User{ID: n})
	}

	proxy := UserListProxy{AppDatabase: database, StackCapacity: 2, StackCache: UserList{}}

	knownIds := [3]int32{database[3].ID, database[4].ID, database[5].ID}

	// embeded test -- i see it enables us to resuse declared structs
	t.Run("Find user -- empty cache", func(t *testing.T) {
		user, err := proxy.Find(knownIds[0])
		if err != nil {
			t.Fatal(err.Error())
		}

		if user.ID != knownIds[0] {
			t.Error("Returned user id doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("After one successful search size of cache must be 1")
		}

		if proxy.LastSearchUsedCache {
			t.Error("No user can be returned from an empty cache")
		}
	})

	t.Run("find - one user, request for same user", func(t *testing.T) {
		user, err := proxy.Find(knownIds[0])
		if err != nil {
			t.Fatal(err.Error())
		}

		if user.ID != knownIds[0] {
			t.Error("Returned user does not match expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("size of cache must not grow above 1")
		}

		if !proxy.LastSearchUsedCache {
			t.Error("The user should have been returned from the cache")
		}
	})

	t.Run("find - overflow the stack", func(t *testing.T) {
		user1, err := proxy.Find(knownIds[0])
		if err != nil {
			t.Fatal(err.Error())
		}

		user2, _ := proxy.Find(knownIds[1])
		if proxy.LastSearchUsedCache {
			t.Error("This user should'nt be seen in the cache")
		}

		user3, _ := proxy.Find(knownIds[2])
		if proxy.LastSearchUsedCache {
			t.Error("This user shouldn't be seen on the cache yet")
		}

		// check if we can find user1 in the cache
		for i := 0; i < len(proxy.StackCache); i++ {
			if proxy.StackCache[i].ID == user1.ID {
				t.Error("User1 should not be in the cahe")
			}
		}

		if len(proxy.StackCache) != 2 {
			t.Error("After inserting 3 users the cache length should be 2")
		}

		// check if an item exists that is not user2 or user3
		for _, v := range proxy.StackCache {
			if v != user2 && v != user3 {
				t.Error("A non expected user was found on the cache")
			}
		}
	})
}
