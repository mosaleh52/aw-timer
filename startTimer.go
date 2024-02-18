package main

import (
	"fmt"
	"os"
	"time"

	"github.com/1set/todotxt"
)

func startTimer(apiUrl, bucketId, dateLayout, taskString string) {
	todoItem, err := todotxt.ParseTask(taskString)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error parsing todo", err)
		os.Exit(-1)
	}

	if todoItem.AdditionalTags["uuid"] == "" {
		fmt.Fprintln(os.Stderr, "to start todo there should be a uuid attr")
		os.Exit(-1)
	}

	if todoItem.AdditionalTags["timeStamp"] != "" {
		fmt.Fprintln(os.Stderr, "this an already running todo")
		os.Exit(-1)
	}
	timeStamp := " timeStamp:" + time.Now().UTC().Format(dateLayout)
	running := " running:true"
	modifiedTodo, err := todotxt.ParseTask(todoItem.String() + timeStamp + running)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error in modifying to todo", err)
		os.Exit(-1)
	}
	payload, err := TodoEvent{modifiedTodo}.MarshalJSON()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error appending to file", err)
		os.Exit(-1)
	}
	err = createAwEvent(apiUrl, bucketId, payload)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("todo started")

	// appendToFileFunc(timeFilePath, modifiedTodo.String())
	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, "error appending to file", err)
	// 	os.Exit(-1)
	// } else {
	// 	fmt.Fprintln(os.Stderr, "there is already running todo :", getCurrentTimmer(timeFilePath))
	// 	os.Exit(-1)
	// }
	// }
}
