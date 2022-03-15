package client

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/go-resty/resty/v2"
)

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL: baseURL,
		h:       *resty.New(),
	}
}

func (c *Client) GetJobList() (*JobList, error) {
	resp, err := c.h.R().SetResult(JobList{}).
		Get(c.baseURL + "/jobs/overview")

	if err != nil {
		return nil, err
	}

	jobList := resp.Result().(JobList)
	sort.Sort(JobByStartTime(jobList.Jobs))

	return &jobList, err
}

func (c *Client) StopJob(jobID string) error {
	_, err := c.h.R().Patch(fmt.Sprintf("%s/jobs/%s?mode=cancel", c.baseURL, jobID))

	if err != nil {
		return err
	}

	return nil
}

// TriggerSavepoint triggers an async savepoint operation.
func (c *Client) TriggerSavepoint(jobID string, dir string, cancel bool) (*SavepointTriggerID, error) {
	url := fmt.Sprintf("%s/jobs/%s/savepoints", c.baseURL, jobID)
	jsonStr := fmt.Sprintf(`{
		"target-directory" : "%s",
		"cancel-job" : %v
	}`, dir, cancel)

	resp, err := c.h.R().
		SetResult(SavepointTriggerID{}).
		SetBody(jsonStr).
		Post(url)

	if err != nil {
		return nil, err
	}

	triggerID := resp.Result().(SavepointTriggerID)

	return &triggerID, err
}

func (c *Client) GetSavepointStatus(
	jobID string, triggerID string) (*SavepointStatus, error) {
	url := fmt.Sprintf("%s/jobs/%s/savepoints/%s", c.baseURL, jobID, triggerID)
	status := &SavepointStatus{JobID: jobID, TriggerID: triggerID}

	var stateID SavepointStateID

	resp, err := c.h.R().SetResult(RawJSON{}).Get(url)
	if err != nil {
		return nil, err
	}

	rootJSON := resp.Result().(RawJSON)

	if state, ok := rootJSON["status"]; ok && state != nil {
		err = json.Unmarshal(*state, &stateID)
		if err != nil {
			return nil, err
		}
		if stateID.ID == SavepointStateCompleted {
			status.Completed = true
		} else {
			status.Completed = false
		}
	}

	if op, ok := rootJSON["operation"]; ok && op != nil {
		var opJSON RawJSON
		err = json.Unmarshal(*op, &opJSON)
		if err != nil {
			return nil, err
		}

		// Success
		if location, ok := opJSON["location"]; ok && location != nil {
			err = json.Unmarshal(*location, &status.Location)
			if err != nil {
				return nil, err
			}
		}

		// Failure
		if failureCause, ok := opJSON["failure-cause"]; ok && failureCause != nil {
			err = json.Unmarshal(*failureCause, &status.FailureCause)
			if err != nil {
				return nil, err
			}
		}
	}

	return status, err
}

// TakeSavepoint takes savepoint, blocks until it succeeds or fails.
func (c *Client) TakeSavepoint(jobID string, dir string) (*SavepointStatus, error) {
	status := &SavepointStatus{JobID: jobID}

	triggerID, err := c.TriggerSavepoint(jobID, dir, false)
	if err != nil {
		return nil, err
	}

	for i := 0; i < 12; i++ {
		status, err = c.GetSavepointStatus(jobID, triggerID.RequestID)
		if err == nil && status.Completed {
			return status, nil
		}
		time.Sleep(5 * time.Second)
	}

	return status, err
}

func (c *Client) TakeSavepointAsync(jobID string, dir string) (string, error) {
	triggerID, err := c.TriggerSavepoint(jobID, dir, false)
	if err != nil {
		return "", err
	}

	return triggerID.RequestID, err
}

func (c *Client) GetJobExceptions(jobId string) (*JobExceptions, error) {
	url := fmt.Sprintf("%s/jobs/%s/exceptions", c.baseURL, jobId)

	resp, err := c.h.R().SetResult(JobExceptions{}).Get(url)
	if err != nil {
		return nil, err
	}

	exceptions := resp.Result().(JobExceptions)

	return &exceptions, nil
}
