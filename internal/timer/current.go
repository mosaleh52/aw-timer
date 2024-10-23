package timer

import (
	"fmt"
	"time"

	"github.com/mosaleh52/aw-timer/internal/api"
)

func GetCurrentTodos(apiUrl, bucketId, dateLayout string) []api.AwResponseEvent {
	url := fmt.Sprintf("%s/query/", apiUrl)

	queryData := api.QueryData{
		Query: []string{
			fmt.Sprintf("stop = query_bucket(find_bucket(\"%s\"));", bucketId),
			"run = filter_keyvals(stop, \"running\", [true , \"true\"]);",
			"RETURN = sort_by_duration(run);",
			";",
		},
		Timeperiods: []string{
			fmt.Sprintf("%s/%s", (time.Now().Add(-24 * time.Hour)).Format(dateLayout), (time.Now().Add(24 * time.Hour)).Format(dateLayout)),
		},
	}
	response := api.SendAwQuery(url, queryData)
	return response
}
