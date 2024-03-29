// Copyright © 2020 The Tekton Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package builder

import (
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
)

type TaskrunList []tkr

type tkr struct {
	TaskrunName string
	*v1beta1.PipelineRunTaskRunStatus
}

func (s TaskrunList) Len() int      { return len(s) }
func (s TaskrunList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s TaskrunList) Less(i, j int) bool {
	return s[j].Status.StartTime.Before(s[i].Status.StartTime)
}
