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
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
)

func init() {
	if err := AddToScheme(scheme.Scheme); err != nil {
		panic(fmt.Sprintf("failed to add ee.kubermatic scheme: %v", err))
	}
}

// GroupName is the group name use in this package.
const GroupName = "ee.kubermatic.k8c.io"
const GroupVersion = "v1"

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme

	// SchemeGroupVersion is group version used to register these objects.
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: GroupVersion}
)

// Resource takes an unqualified resource and returns a Group qualified GroupResource.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&AddonConfig{},
		&AddonConfigList{},
		&Addon{},
		&AddonList{},
		&AdmissionPlugin{},
		&AdmissionPluginList{},
		&Alertmanager{},
		&AlertmanagerList{},
		&AllowedRegistry{},
		&AllowedRegistryList{},
		&Cluster{},
		&ClusterList{},
		&ClusterTemplate{},
		&ClusterTemplateList{},
		&ClusterTemplateInstance{},
		&ClusterTemplateInstanceList{},
		&Constraint{},
		&ConstraintList{},
		&ConstraintTemplate{},
		&ConstraintTemplateList{},
		&EtcdBackupConfig{},
		&EtcdBackupConfigList{},
		&EtcdRestore{},
		&EtcdRestoreList{},
		&ExternalCluster{},
		&ExternalClusterList{},
		&GroupProjectBinding{},
		&GroupProjectBindingList{},
		&IPAMPool{},
		&IPAMPoolList{},
		&IPAMAllocation{},
		&IPAMAllocationList{},
		&KubermaticConfiguration{},
		&KubermaticConfigurationList{},
		&DashboardConfiguration{},
		&DashboardConfigurationList{},
		&MLAClusterConfiguration{},
		&MLAClusterConfigurationList{},
		&MLARuleGroup{},
		&MLARuleGroupList{},
		&Preset{},
		&PresetList{},
		&ResourceQuota{},
		&ResourceQuotaList{},
		&Seed{},
		&SeedList{},
		&SSHKeyBinding{},
		&SSHKeyBindingList{},
		&User{},
		&UserList{},
		&UserProjectBinding{},
		&UserProjectBindingList{},
		&UserSSHKey{},
		&UserSSHKeyList{},
	)

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
