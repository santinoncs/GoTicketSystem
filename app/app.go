package app

import (
	"fmt"
	"time"
)

var Jobchan1 chan Job
var Jobchan2 chan Job

// type App struct {

// }

// Response : here you tell us what Response is
type Response struct {
	Success bool
	Message string
}

// Job : here you tell us what Job is
type Job struct {
	ID           int
	Question     string
	ResponseChan chan Response
}

func newResponse(success bool, message string) *Response {
	return &Response{
		Success: success,
		Message: message,
	}
}

// Start : starting workers
func Start() {

	numWorkers := 2

	for j := 1; j <= numWorkers; j++ {
		go newWorker(j)
	}
}

func newWorker(j int) {

	res := newResponse(true, "bye")

	for {
		select {
		case msg1 := <-Jobchan1:
			fmt.Println("escuchando en jobchan1")
			msg1.ResponseChan <- *res
			close(msg1.ResponseChan)
		case msg2 := <-Jobchan2:
			time.Sleep(4 * time.Second)
			fmt.Println("escuchando en jobchan2")
			msg2.ResponseChan <- *res
			close(msg2.ResponseChan)
		}
	}
}

func newJob(priority int, question string) Job {

	responseChan1 := make(chan Response)

	j := Job{ID: priority, Question: question, ResponseChan: responseChan1}

	return j
}

// Post : escribo los jobs en jobs channel ya con los datos de prio y message
func Post(priority int, question string) Response {

	j := newJob(priority, question)

	Jobchan1 = make(chan Job)
	Jobchan2 = make(chan Job)

	// aqui lanzo con go func el escribir en el channel de jobs

	go func() {

		if priority == 1 {
			Jobchan1 <- j
		}
		if priority == 2 {
			Jobchan2 <- j
		}

	}()

	channelListenR := j.ResponseChan

	select {
	case Response := <-channelListenR:
		return Response
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
		res := newResponse(true, "error")
		return *res
	}

}

// func (j Job) process() Response {
// 	// aqui leo el job, cojo el channel Response y me pongo a leer del channel  // hasta que  worker haya enviado el reponse que luego devolveré cuando     // llamen  a esta funcion

// }
