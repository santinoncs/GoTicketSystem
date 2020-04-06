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

func newWorker(j int, jobs <-chan app.Job) {

	d := app.Response{
		Success: true,
		Message: "bye",
	}

	for j := range jobs {

		fmt.Println("worker", j, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", j, "finished job", j)
		j.ResponseChan <- d
	}

}

func main() {

	jobchan1 := make(chan app.Job)

	for j := 1; j <= 2; j++ {
		go newWorker(j, jobchan1)
	}

	response := app.Post(1, "hola", jobchan1)

	fmt.Println(response.Message)

}
