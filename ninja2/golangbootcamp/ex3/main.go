package main

import "fmt"

// User is the basic user structure
type User struct {
	Id             int
	Name, Location string
}

// Greetings prints a hello message
func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

// Player is a composite struct of User and GameId
type Player struct {
	*User
	GameId int
}

func main() {
	p := Player{
		&User{
			Id:       131,
			Name:     "Oscar Almgren",
			Location: "Sverige",
		}, 1,
	}

	fmt.Printf("Id: %d, Name: %s, Location: %s, Game id: %d\n", p.Id, p.Name, p.Location, p.GameId)
}
