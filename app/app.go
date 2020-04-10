package app

import (
	"fmt"
	"time"
	"sync"
)

var jobchan1 chan Job
var jobchan2 chan Job

// type App struct {

// }

// Response : here you tell us what Response is
type Response struct {
	Success bool
	Message string
}

// Status : here you tell us what Status is
type Status struct {
	workers   int
	processed int
	timeProcessed   time.Duration
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
func Start(st *Status,mutex *sync.Mutex) {

	numWorkers := 2

	for j := 1; j <= numWorkers; j++ {
		mutex.Lock()
		st.workers ++
		mutex.Unlock()
		go newWorker(j,st,mutex)
	}

	go func () {
		fmt.Println("Number of Workers:" , st.workers)
	}()
}

func (j Job) process() Response {

// call to newReponse to initialize a response struct
// return response 

	res := newResponse(true, "bye")

	return *res

}

func newWorker(j int,st *Status,mutex *sync.Mutex) {

	for {
		select {
		case msg1 := <-jobchan1:
			time.Sleep(100 * time.Millisecond)
			msg1.ResponseChan <- msg1.process()
			close(msg1.ResponseChan)
		case msg2 := <-jobchan2:
			time.Sleep(4 * time.Second)
			msg2.ResponseChan <- msg2.process()
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
func Post(priority int, question string,mutex *sync.Mutex,st *Status) (Response,*Status) {

	start := time.Now()

	j := newJob(priority, question)

	jobchan1 = make(chan Job)
	jobchan2 = make(chan Job)

	// aqui lanzo con go func el escribir en el channel de jobs

	go func() {

		if priority == 1 {
			jobchan1 <- j
		}
		if priority == 2 {
			jobchan2 <- j
		}

	}()

	channelListenR := j.ResponseChan

	select {
	case Response := <-channelListenR:
		t := time.Now()
		elapsed := t.Sub(start)
		mutex.Lock()
		st.timeProcessed = elapsed
		st.processed ++
		mutex.Unlock()
		fmt.Println(st.timeProcessed)
		return Response,st
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
		res := newResponse(true, "error")
		return *res,st
	}

}

func (st *Status ) GetProcessed() int{
	 return st.processed
}

func (st *Status ) GetWorkers() int{
	 return st.workers
}

