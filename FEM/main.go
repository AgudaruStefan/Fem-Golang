package main

import (
	"fmt"
	"learningo.com/go/server/data"
)

func main() {
	stefan := data.Instructor{Id: 1, LastName: "Agudaru"}
	stefan.FirstName = "Stefan"

	goCourse := data.Course{Id: 21, Name: "Go fundamentals", Instructor: stefan}

	swiftWS := data.NewWorkshop("Swift with iOS", stefan)

	var courses [2]data.Signable
	courses[0] = goCourse
	courses[1] = swiftWS

	for _, course := range courses {
		fmt.Println(course)
	}
}
