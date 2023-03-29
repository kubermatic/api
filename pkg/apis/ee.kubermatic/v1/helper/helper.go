/*
Copyright 2023 The Kubermatic Kubernetes Platform contributors.

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

package helper

import (
	kubermaticeev1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"

	corev1 "k8s.io/api/core/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SetClusterCondition sets a condition on the given cluster using the provided type, status,
// reason and message. It also adds the Kubermatic version and timestamps.
func SetClusterCondition(
	c *kubermaticeev1.Cluster,
	kkpVersion string,
	conditionType kubermaticeev1.ClusterConditionType,
	status corev1.ConditionStatus,
	reason string,
	message string,
) {
	newCondition := kubermaticeev1.ClusterCondition{
		Status:            status,
		KubermaticVersion: kkpVersion,
		Reason:            reason,
		Message:           message,
	}

	oldCondition, hadCondition := c.Status.Conditions[conditionType]
	if hadCondition {
		conditionCopy := oldCondition.DeepCopy()

		// Reset the times before comparing
		conditionCopy.LastHeartbeatTime.Reset()
		conditionCopy.LastTransitionTime.Reset()

		if apiequality.Semantic.DeepEqual(*conditionCopy, newCondition) {
			return
		}
	}

	now := metav1.Now()
	newCondition.LastHeartbeatTime = now
	newCondition.LastTransitionTime = oldCondition.LastTransitionTime
	if hadCondition && oldCondition.Status != status {
		newCondition.LastTransitionTime = now
	}

	if c.Status.Conditions == nil {
		c.Status.Conditions = map[kubermaticeev1.ClusterConditionType]kubermaticeev1.ClusterCondition{}
	}
	c.Status.Conditions[conditionType] = newCondition
}

// SetSeedCondition sets a condition on the given seed using the provided type, status,
// reason and message.
func SetSeedCondition(seed *kubermaticeev1.Seed, conditionType kubermaticeev1.SeedConditionType, status corev1.ConditionStatus, reason string, message string) {
	newCondition := kubermaticeev1.SeedCondition{
		Status:  status,
		Reason:  reason,
		Message: message,
	}

	oldCondition, hadCondition := seed.Status.Conditions[conditionType]
	if hadCondition {
		conditionCopy := oldCondition.DeepCopy()

		// Reset the times before comparing
		conditionCopy.LastHeartbeatTime.Reset()
		conditionCopy.LastTransitionTime.Reset()

		if apiequality.Semantic.DeepEqual(*conditionCopy, newCondition) {
			return
		}
	}

	now := metav1.Now()
	newCondition.LastHeartbeatTime = now
	newCondition.LastTransitionTime = oldCondition.LastTransitionTime
	if hadCondition && oldCondition.Status != status {
		newCondition.LastTransitionTime = now
	}

	if seed.Status.Conditions == nil {
		seed.Status.Conditions = map[kubermaticeev1.SeedConditionType]kubermaticeev1.SeedCondition{}
	}
	seed.Status.Conditions[conditionType] = newCondition
}
