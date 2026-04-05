package state

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

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

func NewRun(pipelineName string) (*PipelineRun){
	return &PipelineRun{
		Id: uuid.NewString(),
		Status: StatusPending,
		TriggeredAt: time.Now(),
		Stages: []*StageRun{},
	}
}

func (p *PipelineRun)StartRun() error {
	if p.Id != ""{
		p.Status = StatusRunning
		return nil
	}else{
		return fmt.Errorf("No Running PipeLine Found")
	}
}
func (p *PipelineRun) FinishRun(status Status) error{
	if p.Id != ""{
		p.Status = status
		p.FinishedAt = time.Now()
		return nil
	}else{
		return fmt.Errorf("runId Not Found") 
	}	
}


