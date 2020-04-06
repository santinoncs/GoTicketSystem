package app

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

func jobProcess(priority int, question string) Job {

	responseChan1 := make(chan Response)

	j := Job{ID: priority, Question: question, ResponseChan: responseChan1}

	return j
}

// Post : escribo los jobs en jobs channel ya con los datos de prio y message
func Post(priority int, question string, jobchan chan Job) Response {

	j := jobProcess(priority, question)

	// aqui lanzo con go func el escribir en el channel de jobs

	go func() {

		if priority == 1 {

			// escribo en jobsChan que trata con prioridad 1 , pero acepta jobs

			jobchan <- j

		}
		close(jobchan)
	}()
	return j.process()
}

func (j Job) process() Response {
	// aqui leo el job, cojo el channel Response y me pongo a leer del channel  // hasta que  worker haya enviado el reponse que luego devolveré cuando     // llamen  a esta funcion

	channelListenR := j.ResponseChan

	select {
	case Response := <-channelListenR:
		return Response
	}

}
