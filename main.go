package main

import (
	"fmt"
	app "github.com/santinoncs/GoTicketSystem/app"
	"time"
)




// IncomingQuestion : here you tell us what IncomingQuestion is
type IncomingQuestion struct {
	Priority int
	Question string
}


func newWorker(j int) {


	res := app.NewResponse(true, "bye")


	fmt.Println( "arrancando workers")

	for {
		select {
		case msg1 := <-app.Jobchan1:
			fmt.Println("escuchando en jobchan1")
			msg1.ResponseChan <- *res
			close(msg1.ResponseChan)
		case msg2 := <-app.Jobchan2:
			time.Sleep(4 * time.Second)
			fmt.Println("escuchando en jobchan2")
			msg2.ResponseChan <- *res
			close(msg2.ResponseChan)
		default:
			break
		}
	}
}

func main() {


	numWorkers := 2


	for j := 1; j <= numWorkers ; j++ {
		go newWorker(j)
	}

	response := app.Post(2, "hola")

	fmt.Println(response.Message)

	

}
