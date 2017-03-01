package models

import (
	"errors"
	"time"
)

var (
	// UserList is a list of users for use in testing
	UserList map[int]*User
)

func init() {
	UserList = make(map[int]*User)
	u := User{
		1,
		"user_11111",
		"astaxie",
		"11111",
		time.Now(),
		time.Now(),
		Profile{
			"male",
			20,
			"astaxie@gmail.com",
			"a person",
			"fake.com/icon.png",
		},
	}
	UserList[11111] = &u
}

// User represents an end user on the site. Includes all info necessary to identify a user and includes that user's playlists.
// swagger:model
type User struct {
	// the ID of the user
	//
	// required: true
	// read only: true
	Id int `json:"id"`

	// the username of the user
	//
	// required: true
	Username string `json:"username"`

	// the way a user's name will be displayed, a nickname, if you will
	DisplayName string `json:"display_name"`

	// the user's password. we may be able to avoid doing this.
	Password string `json:"password,omitempty"`

	// the time the user was created
	Created time.Time `json:"created"`

	// the last time the user interacted with the site
	LastActivity time.Time `json:"last_activity"`

	Profile Profile
}

type Profile struct {
	// the user's gender
	//
	// required: true
	Gender string

	// the user's age
	//
	// required: true
	Age int

	// the user's email
	//
	// required: true
	// pattern: [^@]+@[\w\d\.]+
	Email string `json:"email,omitempty"`

	// a brief self-description of the user
	Description string `json:"description"`

	// a URL pointing to the user's icon
	IconURL string `json:"icon_url"`
}

func AddUser(u User) int {
	u.Id = int(time.Now().Unix())
	UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid int) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[int]*User {
	return UserList
}

func UpdateUser(uid int, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		if uu.Password != "" {
			u.Password = uu.Password
		}
		if uu.Profile.Age != 0 {
			u.Profile.Age = uu.Profile.Age
		}
		if uu.Profile.Gender != "" {
			u.Profile.Gender = uu.Profile.Gender
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid int) {
	delete(UserList, uid)
}
