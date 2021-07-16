package main

import (
	"context"
	"goMicroserviceDemo/Service"
	"goMicroserviceDemo/kafka"
	"log"
	"os"
	"os/signal"
)
const topic = "demo-topic"
const broker ="localhost:9092"

func main() {
	receivedEvent := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())

	eventPrinterService := Service.EventPrinter{}
	eventListener := kafka.NewEventListener(broker,topic, ctx)

	var terminate chan os.Signal = make(chan os.Signal, 1)
	signal.Notify(terminate, os.Interrupt)

	go eventListener.StartAndListenAndPushToChannel("deafult", receivedEvent)

	for{
		select {
		case received := <-receivedEvent:
			eventPrinterService.Print(received)
		case  <- terminate:
			cancel()
			<-receivedEvent
			log.Println("exiting..")
			os.Exit(1)

		}
	}
}





/*
type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title:"Test Title", Desc: "Test Description", Content: "Test Content"},
	}
	fmt.Println("EndPoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}


func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
*/
