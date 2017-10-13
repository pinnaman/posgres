package main

import (
	"ajay/models"
	"fmt"
)

func main() {

	evts, err := models.AllEvents()
	if err != nil {
		fmt.Println("Error")
		return
	}

	for _, evt := range evts {
		fmt.Println(evt)
		//fmt.Fprintf(w, "%s, %s, %s, %s, %s, %s\n", evt.Id, evt.Created, evt.Start, evt.End, evt.Title, evt.Completed)
	}

}
