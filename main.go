package main

import (
	"fmt"
	app "github.com/santinoncs/GoTicketSystem/app"
	"sync"
)

// IncomingQuestion : here you tell us what IncomingQuestion is
type IncomingQuestion struct {
	Priority int
	Question string
}


func main() {

	st := app.NewStatus()
	var mutex = &sync.Mutex{}

	app.Start(st,mutex)

	response,st := app.Post(1, "hola",mutex, st)


	fmt.Println("message respond is:", response.Message)
	fmt.Println("Processed questions are:", st.GetProcessed())
	fmt.Println("Number of Workers:", st.GetWorkers())
	fmt.Println("average_response_time in ms:", st.GetAverage())


	


}
