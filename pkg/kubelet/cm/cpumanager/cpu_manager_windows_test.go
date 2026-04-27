//go:build windows

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
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	runtimeapi "k8s.io/cri-api/pkg/apis/runtime/v1"
	kubefeatures "k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/pkg/kubelet/winstats"
	"k8s.io/utils/cpuset"
)

func (rt mockRuntimeService) getCPUSetFromResources(resources *runtimeapi.ContainerResources) cpuset.CPUSet {
	if !utilfeature.DefaultFeatureGate.Enabled(kubefeatures.WindowsCPUAndMemoryAffinity) {
		return cpuset.New()
	}
	if resources != nil && resources.Windows != nil {
		var cpus []int
		for _, affinity := range resources.Windows.AffinityCpus {
			ga := winstats.GroupAffinity{Mask: affinity.CpuMask, Group: uint16(affinity.CpuGroup)}
			cpus = append(cpus, ga.Processors()...)
		}
		return cpuset.New(cpus...)
	}
	return cpuset.New()
}
