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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:Enum=Started;StsRebuilding;Completed;EtcdLauncherNotEnabled

// EtcdRestorePhase represents the lifecycle phase of an EtcdRestore.
type EtcdRestorePhase string

const (
	// EtcdRestorePhaseStarted indicates that the restore has started.
	EtcdRestorePhaseStarted EtcdRestorePhase = "Started"

	// EtcdRestorePhaseStsRebuilding indicates that the old Etcd statefulset has been deleted and is now rebuilding.
	EtcdRestorePhaseStsRebuilding EtcdRestorePhase = "StsRebuilding"

	// EtcdRestorePhaseCompleted indicates that the old Etcd statefulset has completed successfully.
	EtcdRestorePhaseCompleted EtcdRestorePhase = "Completed"

	// EtcdRestorePhaseEtcdLauncherNotEnabled indicates that etcd-launcher is not enabled.
	EtcdRestorePhaseEtcdLauncherNotEnabled EtcdRestorePhase = "EtcdLauncherNotEnabled"
)

// +genclient
// +kubebuilder:object:generate=true
// +kubebuilder:resource:categories=kkpee
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".spec.cluster.name",name="Cluster",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.phase",name="Phase",type="string"
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

// EtcdRestore specifies how to restore an etcd backup for a usercluster.
//
// Note that this resource is part of a KKP Enterprise feature and is not used in the Community Edition.
type EtcdRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EtcdRestoreSpec   `json:"spec,omitempty"`
	Status EtcdRestoreStatus `json:"status,omitempty"`
}

// EtcdRestoreSpec specifies details of an etcd restore.
type EtcdRestoreSpec struct {
	// Cluster is the reference to the cluster whose etcd will be backed up.
	Cluster ClusterReference `json:"cluster"`
	// BackupName is the name of the backup to restore from
	BackupName string `json:"backupName"`
	// BackupDownloadCredentialsSecret is the name of a secret in the cluster-xxx namespace containing
	// credentials needed to download the backup
	BackupDownloadCredentialsSecret string `json:"backupDownloadCredentialsSecret,omitempty"`
	// Destination indicates where the backup was stored. The destination name should correspond to a destination in
	// the cluster's Seed.Spec.EtcdBackupRestore. If empty, it will use the legacy destination configured in Seed.Spec.BackupRestore
	Destination string `json:"destination,omitempty"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// EtcdRestoreList is a list of etcd restores.
type EtcdRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []EtcdRestore `json:"items"`
}

type EtcdRestoreStatus struct {
	Phase EtcdRestorePhase `json:"phase"`
	// +optional
	RestoreTime metav1.Time `json:"restoreTime,omitempty"`
}
