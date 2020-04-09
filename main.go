package main

import (
	"fmt"
	app "github.com/santinoncs/GoTicketSystem/app"
)

// IncomingQuestion : here you tell us what IncomingQuestion is
type IncomingQuestion struct {
	Priority int
	Question string
}

func main() {

	app.Start()

	response := app.Post(1, "hola")

	fmt.Println(response.Message)

}
