package model

import "fmt"

// User define a user
type User struct {
	ScreenName string `json:screen_name`
}

// String stringify the User struct
func (u *User) String() string {
	return fmt.Sprintf("User : %v", u.ScreenName)
}
