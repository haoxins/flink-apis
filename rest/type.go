package rest

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

const (
	SavepointStateInProgress = "IN_PROGRESS"
	SavepointStateCompleted  = "COMPLETED"
)

// Client - The Flink API client.
type Client struct {
	h       resty.Client
	baseURL string
}

type ResponseError struct {
	StatusCode int
	Status     string
}

func (e *ResponseError) Error() string {
	return e.Status
}

type JobException struct {
	Exception string `json:"exception"`
	Location  string `json:"location"`
}

type JobExceptions struct {
	Exceptions []JobException `json:"all-exceptions"`
}

// Job defines the Flink job status.
type Job struct {
	Id        string `json:"jid"`
	State     string `json:"state"`
	Name      string `json:"name"`
	StartTime int64  `json:"start-time"`
	EndTime   int64  `json:"end-time"`
	Duration  int64  `json:"duration"`
}

// JobList defines the Flink job list.
type JobList struct {
	Jobs []Job
}

type JobByStartTime []Job

func (jst JobByStartTime) Len() int           { return len(jst) }
func (jst JobByStartTime) Swap(i, j int)      { jst[i], jst[j] = jst[j], jst[i] }
func (jst JobByStartTime) Less(i, j int) bool { return jst[i].StartTime > jst[j].StartTime }

// SavepointTriggerID defines trigger ID of an async savepoint operation.
type SavepointTriggerID struct {
	RequestID string `json:"request-id"`
}

// SavepointFailureCause defines the cause of savepoint failure.
type SavepointFailureCause struct {
	ExceptionClass string `json:"class"`
	StackTrace     string `json:"stack-trace"`
}

// SavepointStateID - enum("IN_PROGRESS", "COMPLETED").
type SavepointStateID struct {
	ID string `json:"id"`
}

// SavepointStatus defines savepoint status of a job.
type SavepointStatus struct {
	JobID        string
	TriggerID    string
	Completed    bool
	Location     string
	FailureCause SavepointFailureCause
}

func (s *SavepointStatus) IsSuccessful() bool {
	return s.Completed && s.FailureCause.StackTrace == ""
}

func (s *SavepointStatus) IsFailed() bool {
	return s.Completed && s.FailureCause.StackTrace != ""
}

type RawJSON map[string]*json.RawMessage
