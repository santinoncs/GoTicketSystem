package app

import (
	"fmt"
	"time"
)


var Jobchan1 chan Job
var Jobchan2 chan Job




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

func NewResponse(success bool,message string) *Response {
	return &Response {
		Success: success,
		Message: message,
	   }  
}

func jobProcess(priority int, question string) Job {

	responseChan1 := make(chan Response)

	j := Job{ID: priority, Question: question, ResponseChan: responseChan1}

	return j
}

// Post : escribo los jobs en jobs channel ya con los datos de prio y message
func Post(priority int, question string) Response {

	j := jobProcess(priority, question)

	Jobchan1 = make(chan Job)
	Jobchan2 = make(chan Job)


	// aqui lanzo con go func el escribir en el channel de jobs

	go func() {

		if priority == 1 {

			// escribo en jobsChan que trata con prioridad 1 , pero acepta jobs

			Jobchan1 <- j

		}
		if priority == 2 {

			// escribo en jobsChan que trata con prioridad 1 , pero acepta jobs

			Jobchan2 <- j

		}


	}()
	return j.process()
}

func (j Job) process() Response {
	// aqui leo el job, cojo el channel Response y me pongo a leer del channel  // hasta que  worker haya enviado el reponse que luego devolveré cuando     // llamen  a esta funcion

	channelListenR := j.ResponseChan

	select {
	case Response := <-channelListenR:
		return Response
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
		res := NewResponse(true, "error")
		return *res
	}
}
