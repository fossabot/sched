// Copyright 2018 Changkun Ou. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package tests

import (
	"fmt"
	"time"
)

// NonExportTask implements task.Interface However it has no
// export field.
// NOTE: such tasks are unable to be persist by sched
type NonExportTask struct {
	id        string
	execution time.Time
}

// NewNonExportTask creates a task
func NewNonExportTask(id string, e time.Time) *NonExportTask {
	return &NonExportTask{
		id:        id,
		execution: e,
	}
}

// GetID get task id
func (t *NonExportTask) GetID() (id string) {
	id = t.id
	return
}

// GetExecution get execution time
func (t *NonExportTask) GetExecution() (execute time.Time) {
	execute = t.execution
	return
}

// GetTimeout get timeout of execution
func (t *NonExportTask) GetTimeout() (executeTimeout time.Duration) {
	return time.Second
}

// GetRetryDuration get retry execution duration
func (t *NonExportTask) GetRetryDuration() (duration time.Duration) {
	return time.Second
}

// SetID sets the id of a task
func (t *NonExportTask) SetID(id string) {
	t.id = id
}

// SetExecution sets the execution time of a task
func (t *NonExportTask) SetExecution(current time.Time) (old time.Time) {
	old = t.execution
	t.execution = current
	return
}

// Execute is the actual execution block
func (t *NonExportTask) Execute() (retry bool, fail error) {
	O.Push(t.id)
	fmt.Printf("Execute non export task %s.\n", t.id)
	return false, nil
}
