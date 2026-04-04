package state

import "time"

type Status string


// For storing status Constent
const (
	StatusPending	Status = "pending"
	StatusRunning	Status = "running"
	StatusSuccess	Status = "success"
	StatusFailed	Status = "Failed"
	StatusCancelled	Status = "cancelled"
)

// Struct to define Data used During Stage Management Function
type PipelineRun struct{
	Id 			string
	Status		Status
	TriggeredAt	time.Time
	FinishedAt 	time.Time
	Stages		[]*StageRun
}

type StageRun struct{
	Name		string
	Status		Status
	StartedAt	time.Time
	FinishedAt	time.Time
	Jobs		[]*JobRun
}

type JobRun struct{
	Name		string
	Status		Status
	ExitCode	int
	StartedAt	time.Time
	FinishedAt	time.Time
}

