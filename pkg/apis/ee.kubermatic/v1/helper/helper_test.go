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
	"testing"

	kubermaticeev1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"

	corev1 "k8s.io/api/core/v1"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
)

func TestSetClusterCondition(t *testing.T) {
	conditionType := kubermaticeev1.ClusterConditionSeedResourcesUpToDate

	testCases := []struct {
		name                    string
		cluster                 *kubermaticeev1.Cluster
		conditionStatus         corev1.ConditionStatus
		conditionReason         string
		conditionMessage        string
		conditionChangeExpected bool
	}{
		{
			name: "Condition already exists, nothing to do",
			cluster: getCluster(conditionType, &kubermaticeev1.ClusterCondition{
				Status:            corev1.ConditionTrue,
				KubermaticVersion: "test",
				Reason:            "my-reason",
				Message:           "my-message",
			}),
			conditionStatus:         corev1.ConditionTrue,
			conditionReason:         "my-reason",
			conditionMessage:        "my-message",
			conditionChangeExpected: false,
		},
		{
			name:                    "Condition doesn't exist and is created",
			cluster:                 getCluster("", nil),
			conditionStatus:         corev1.ConditionTrue,
			conditionReason:         "my-reason",
			conditionMessage:        "my-message",
			conditionChangeExpected: true,
		},
		{
			name: "Update because of Kubermatic version",
			cluster: getCluster(conditionType, &kubermaticeev1.ClusterCondition{
				Status:            corev1.ConditionTrue,
				KubermaticVersion: "outdated",
				Reason:            "my-reason",
				Message:           "my-message",
			}),
			conditionStatus:         corev1.ConditionTrue,
			conditionReason:         "my-reason",
			conditionMessage:        "my-message",
			conditionChangeExpected: true,
		},
		{
			name: "Update because of status",
			cluster: getCluster(conditionType, &kubermaticeev1.ClusterCondition{
				Status:            corev1.ConditionFalse,
				KubermaticVersion: "test",
				Reason:            "my-reason",
				Message:           "my-message",
			}),
			conditionStatus:         corev1.ConditionTrue,
			conditionReason:         "my-reason",
			conditionMessage:        "my-message",
			conditionChangeExpected: true,
		},
		{
			name: "Update because of reason",
			cluster: getCluster(conditionType, &kubermaticeev1.ClusterCondition{
				Status:            corev1.ConditionTrue,
				KubermaticVersion: "test",
				Reason:            "outdated-reason",
				Message:           "my-message",
			}),
			conditionStatus:         corev1.ConditionTrue,
			conditionReason:         "my-reason",
			conditionMessage:        "my-message",
			conditionChangeExpected: true,
		},
		{
			name: "Update because of message",
			cluster: getCluster(conditionType, &kubermaticeev1.ClusterCondition{
				Status:            corev1.ConditionTrue,
				KubermaticVersion: "test",
				Reason:            "my-reason",
				Message:           "outdated-message",
			}),
			conditionStatus:         corev1.ConditionTrue,
			conditionReason:         "my-reason",
			conditionMessage:        "my-message",
			conditionChangeExpected: true,
		},
	}

	for idx := range testCases {
		tc := testCases[idx]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			initialCluster := tc.cluster.DeepCopy()
			SetClusterCondition(tc.cluster, "test", conditionType, tc.conditionStatus, tc.conditionReason, tc.conditionMessage)
			hasChanged := !apiequality.Semantic.DeepEqual(initialCluster, tc.cluster)
			if hasChanged != tc.conditionChangeExpected {
				t.Errorf("Change doesn't match expectation: hasChanged: %t: changeExpected: %t", hasChanged, tc.conditionChangeExpected)
			}
		})
	}
}

func getCluster(conditionType kubermaticeev1.ClusterConditionType, condition *kubermaticeev1.ClusterCondition) *kubermaticeev1.Cluster {
	c := &kubermaticeev1.Cluster{}
	if condition != nil {
		c.Status.Conditions = map[kubermaticeev1.ClusterConditionType]kubermaticeev1.ClusterCondition{
			conditionType: *condition,
		}
	}

	return c
}
