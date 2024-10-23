package helpers

import (
	"encoding/json"
	"time"

	"github.com/1set/todotxt"
)

type TodoEvent struct {
	*todotxt.Task
}
type InternalTodoEvent struct {
	Label          string            `json:"label"`
	Priority       string            `json:"priority,omitempty"`
	Projects       []string          `json:"projects,omitempty"`
	Contexts       []string          `json:"contexts,omitempty"`
	AdditionalTags map[string]string `json:"additionalTags,omitempty"`
	CreatedDate    time.Time         `json:"createdDate,omitempty"`
	DueDate        time.Time         `json:"dueDate,omitempty"`
	CompletedDate  time.Time         `json:"completedDate,omitempty"`
	Completed      bool              `json:"completed,omitempty"`
	Uuid           string            `json:"uuid"`
	Running        string            `json:"running"`
}

type wrappedEvent struct {
	Timestamp string            `json:"timestamp"`
	Duration  float64           `json:"duration,omitempty"`
	Id        int               `json:"id,omitempty"`
	Data      InternalTodoEvent `json:"data"`
}

func (task TodoEvent) MarshalJSON() ([]byte, error) {
	// TODO:move datelayout to config
	dateLayout := "2006-01-02T15:04:05.999999-07:00"

	internalEvent := InternalTodoEvent{
		Label:          task.Todo,
		Uuid:           task.AdditionalTags["uuid"],
		Running:        "true",
		Priority:       task.Priority,
		Projects:       task.Projects,
		Contexts:       task.Contexts,
		AdditionalTags: task.AdditionalTags,
		CreatedDate:    task.CreatedDate,
		DueDate:        task.DueDate,
		CompletedDate:  task.CompletedDate,
		Completed:      task.Completed,
	}

	wrappedEvent := wrappedEvent{
		Timestamp: time.Now().UTC().Format(dateLayout),
		Data:      internalEvent,
	}

	return json.Marshal(wrappedEvent)
}
