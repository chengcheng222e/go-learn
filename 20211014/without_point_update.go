package main

import "fmt"

type User struct {
	userName string
}

func modify(user User) {
	user.userName = "vernonchen2222"
}

func modifyWithPoint(user *User) {
	user.userName = "vernonchen3333"
}

func main() {

	user := User{userName: "vernonchen"}

	fmt.Println("userName = ", user.userName)

	modify(user)

	fmt.Println("userName = ", user.userName)

	modifyWithPoint(&user)

	fmt.Println("userName = ", user.userName)
}
