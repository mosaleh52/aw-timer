package timer

import (
	"fmt"
	"os"
	"time"

	"github.com/1set/todotxt"
	"github.com/mosaleh52/aw-timer/internal/api"
	"github.com/mosaleh52/aw-timer/internal/helpers"
)

func StartTimer(apiUrl, bucketId, dateLayout, taskString string) {
	todoItem, err := todotxt.ParseTask(taskString)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error parsing todo", err)
		os.Exit(-1)
	}

	// if todoItem.AdditionalTags["uuid"] == "" {
	// 	fmt.Fprintln(os.Stderr, "to start todo there should be a uuid attr")
	// 	os.Exit(-1)
	// }

	if todoItem.AdditionalTags["timeStamp"] != "" {
		fmt.Fprintln(os.Stderr, "this an already running todo")
		os.Exit(-1)
	}
	timeStamp := " timeStamp:" + time.Now().UTC().Format(dateLayout)
	modifiedTodo, err := todotxt.ParseTask(todoItem.String() + timeStamp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error in modifying to todo", err)
		os.Exit(-1)
	}
	payload, err := helpers.TodoEvent{modifiedTodo}.MarshalJSON()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error appending to file", err)
		os.Exit(-1)
	}
	err = api.CreateAwEvent(apiUrl, bucketId, payload)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("todo started")
}
