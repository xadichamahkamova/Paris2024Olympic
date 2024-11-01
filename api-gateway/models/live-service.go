package models 

type LiveStream struct {
	EventID   string            `json:"event_id"`
	LeftSide  string            `json:"left_side"`
	RightSide string            `json:"right_side"`
	Action    map[string]string `json:"action"`
	Timestamp string            `json:"timestamp"`
}
