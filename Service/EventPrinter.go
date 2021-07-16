package Service

import "fmt"

type EventPrinter struct {

}

func (serv *EventPrinter) Print(message string) {
	fmt.Println(message)
}