package main

import (
	"fmt"
	"time"
)

// Response : here you tell us what Response is
type Response struct{
	success bool
	message string
}

// IncomingQuestion : here you tell us what IncomingQuestion is
type IncomingQuestion struct {
	Priority   int         
	Question   string      
}

// Job : here you tell us what Job is
type Job struct{
	id int
	question string
	responseChan chan Response
}

func newWorker(j int,jobs <-chan Job) {

	d := Response{
			success: true,
			message: "bye",
	}

	for j := range jobs {
		
		fmt.Println("worker", j, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", j, "finished job", j)
		j.responseChan <- d
	}
	
}

func jobProcess(priority int, question string) Job{

	responseChan1 := make (chan Response)

	j := Job{id: priority, question: question, responseChan: responseChan1}

	return j
}

// Post : escribo los jobs en jobs channel ya con los datos de prio y message
func Post(priority int, question string,jobchan1 chan Job ) Response{


	j := jobProcess(priority , question)

	// aqui lanzo con go func el escribir en el channel de jobs

	go func() {
		
	
		if priority == 1 {

			// escribo en jobsChan que trata con prioridad 1 , pero acepta jobs 
			
			jobchan1 <- j

		}
		close(jobchan1)
	}()
    return j.process()
}



func (j Job) process() Response {
	// aqui leo el job, cojo el channel Response y me pongo a leer del channel  // hasta que  worker haya enviado el reponse que luego devolveré cuando     // llamen  a esta funcion desde el main 

	channelListenR := j.responseChan
	

	select {
		case Response := <-channelListenR:
			return Response
	}


}

func main() {


	jobchan1 := make(chan Job)

	for j := 1; j <= 2; j++ {
        go newWorker(j,jobchan1)
    }

	response := Post(1, "hola", jobchan1)

	fmt.Println(response.message)

}