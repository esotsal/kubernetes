//go:build !windows

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cpumanager

import (
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"
	"k8s.io/utils/cpuset"
)

func (rt mockRuntimeService) getCPUSetFromResources(resources *runtimeapi.ContainerResources) cpuset.CPUSet {
	if resources != nil && resources.Linux != nil {
		set, err := cpuset.Parse(resources.Linux.CpusetCpus)
		if err != nil {
			rt.t.Errorf("(%v) Cannot parse Linux CPUSet resources %v", rt.testCaseDescription, resources.Linux.CpusetCpus)
			return cpuset.New()
		}
		return set
	}
	return cpuset.New()
}
