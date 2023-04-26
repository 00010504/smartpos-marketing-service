// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/555082f38110f65b60d470107d211fc354a5c55a


package types

// ParentTaskInfo type.
//
// https://github.com/elastic/elasticsearch-specification/blob/555082f38110f65b60d470107d211fc354a5c55a/specification/tasks/_types/TaskListResponseBase.ts#L45-L47
type ParentTaskInfo struct {
	Action             string            `json:"action"`
	Cancellable        bool              `json:"cancellable"`
	Cancelled          *bool             `json:"cancelled,omitempty"`
	Children           []TaskInfo        `json:"children,omitempty"`
	Description        *string           `json:"description,omitempty"`
	Headers            map[string]string `json:"headers"`
	Id                 int64             `json:"id"`
	Node               string            `json:"node"`
	ParentTaskId       *TaskId           `json:"parent_task_id,omitempty"`
	RunningTime        *Duration         `json:"running_time,omitempty"`
	RunningTimeInNanos int64             `json:"running_time_in_nanos"`
	StartTimeInMillis  int64             `json:"start_time_in_millis"`
	Status             *TaskStatus       `json:"status,omitempty"`
	Type               string            `json:"type"`
}

// NewParentTaskInfo returns a ParentTaskInfo.
func NewParentTaskInfo() *ParentTaskInfo {
	r := &ParentTaskInfo{
		Headers: make(map[string]string, 0),
	}

	return r
}
