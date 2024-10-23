package api

type QueryData struct {
	Query       []string `json:"query"`
	Timeperiods []string `json:"timeperiods"`
}

var queryData QueryData = QueryData{
	Query: []string{
		// fmt.Sprintf("stop = query_bucket(find_bucket(\"%s\"));", bucketId),
		"run = filter_keyvals(stop, \"running\", [true , \"true\"]);",
		"RETURN = sort_by_duration(run);",
		";",
	},
	Timeperiods: []string{
		// fmt.Sprintf("%s/%s", (time.Now().Add(-24 * time.Hour)).Format(dateLayout), (time.Now().Add(24 * time.Hour)).Format(dateLayout)),
	},
}

type AwResponseEvent struct {
	ID        int                    `json:"id,omitempty"`
	Duration  float64                `json:"duration,omitempty"`
	Timestamp string                 `json:"timestamp"`
	Data      map[string]interface{} `json:"data"`
}
