/*
Copyright 2025 The Kubernetes Authors.

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

package lifecycle

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/kubernetes/pkg/security/apparmor"
)

const (
	// AppArmorNotAdmittedReason is used to denote that the pod was
	// rejected admission to the node because of AppArmor
	AppArmorNotAdmittedReason = "AppArmor"
)

type appArmorAdmitHandler struct {
	apparmor.Validator
}

var _ PodAdmitHandler = &appArmorAdmitHandler{}

// NewAppArmorAdmitHandler returns a PodAdmitHandler which is used to evaluate
// if a pod can be admitted from the perspective of AppArmor.
func NewAppArmorAdmitHandler(validator apparmor.Validator) PodAdmitHandler {
	return &appArmorAdmitHandler{
		Validator: validator,
	}
}

func (a *appArmorAdmitHandler) Admit(attrs *PodAdmitAttributes) PodAdmitResult {
	// If the pod is already running or terminated, no need to recheck AppArmor.
	if attrs.Pod.Status.Phase != v1.PodPending {
		return PodAdmitResult{Admit: true}
	}

	err := a.Validate(attrs.Pod)
	if err != nil {
		return PodAdmitResult{
			Admit:   false,
			Reason:  AppArmorNotAdmittedReason,
			Message: fmt.Sprintf("Cannot enforce AppArmor: %v", err),
		}
	}
	return PodAdmitResult{
		Admit: true,
	}
}
