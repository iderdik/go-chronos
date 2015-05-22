package chronos

import (
	"net/url"
	"path"
)

// Jobs is a slice of jobs
type Jobs []Job

// A Job defines a chronos job
// https://github.com/mesos/chronos/blob/master/docs/docs/api.md#job-configuration
type Job struct {
	Name                   string              `json:"name,omitempty"`
	Command                string              `json:"command,omitempty"`
	Shell                  bool                `json:"shell,omitempty"`
	Epsilon                string              `json:"epsilon,omitempty"`
	Executor               string              `json:"executor,omitempty"`
	ExecutorFlags          string              `json:"executorFlags,omitempty"`
	Retries                int                 `json:"retries,omitempty"`
	Owner                  string              `json:"owner,omitempty"`
	OwnerName              string              `json:"ownerName,omitempty"`
	Description            string              `json:"description,omitempty"`
	Async                  bool                `json:"async,omitempty"`
	SuccessCount           int                 `json:"successCount,omitempty"`
	ErrorCount             int                 `json:"errorCount,omitempty"`
	LastSuccess            string              `json:"lastSuccess,omitempty"`
	LastError              string              `json:"lastError,omitempty"`
	CPUs                   float32             `json:"cpus,omitempty"`
	Disk                   float32             `json:"disk,omitempty"`
	Mem                    float32             `json:"mem,omitempty"`
	Disabled               bool                `json:"disabled,omitempty"`
	SoftError              bool                `json:"softError,omitempty"`
	DataProcessingJobType  bool                `json:"dataProcessingJobType,omitempty"`
	ErrorsSinceLastSuccess int                 `json:"errorsSinceLastSuccess,omitempty"`
	URIs                   []string            `json:"uris,omitempty"`
	EnvironmentVariables   []map[string]string `json:"environmentVariables,omitempty"`
	Arguments              []string            `json:"arguments,omitempty"`
	HighPriority           bool                `json:"highPriority,omitempty"`
	RunAsUser              string              `json:"runAsUser,omitempty"`
	Container              *Container          `json:"container,omitempty"`
	Schedule               string              `json:"schedule,omitempty"`
	ScheduleTimeZone       string              `json:"scheduleTimeZone,omitempty"`
	Constraints            []map[string]string `json:"constraints,omitempty"`
	Parents                []string            `json:"parents,omitempty"`
}

// Jobs gets all jobs that chronos knows about
func (client *Client) Jobs() (*Jobs, error) {
	jobs := new(Jobs)

	err := client.apiGet(ChronosAPIJobs, jobs)

	if err != nil {
		return nil, err
	}

	return jobs, nil
}

// DeleteJob will delete a chronos job
// name: The name of job you wish to delete
func (client *Client) DeleteJob(name string) error {
	return client.apiDelete(path.Join(ChronosAPIJob, name), nil)
}

// DeleteJobTasks will delete all tasks associated with a job.
// name: The name of the job whose tasks you wish to delete
func (client *Client) DeleteJobTasks(name string) error {
	return client.apiDelete(path.Join(ChronosAPIKillJobTask, name), nil)
}

// StartJob can manually start a job
// name: The name of the job to start
// args: A map of arguments to append to the job's command
func (client *Client) StartJob(name string, args map[string]string) error {
	queryValues := url.Values{}
	for key, value := range args {
		queryValues.Add(key, value)
	}

	uri := path.Join(ChronosAPIJob, name) + "?" + queryValues.Encode()
	return client.apiPut(uri, nil)
}
