#!/usr/bin/env bash

# Copyright 2023 The Kubermatic Kubernetes Platform contributors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail

cd $(dirname $0)/..
source hack/lib.sh

CRD_DIR=crd/k8c.io
CODEGEN_DIR=pkg/generated

echodate "Removing old generated clients"
rm -rf "$CODEGEN_DIR"

echodate "Creating vendor directory"
go mod vendor

echodate "Generating Kubernetes clientset"

echo "" > /tmp/headerfile

# no deepcopy here, as controller-gen takes care of that
bash vendor/k8s.io/code-generator/generate-groups.sh client,lister,informer \
  k8c.io/api/v3/$CODEGEN_DIR \
  k8c.io/api/v3/pkg/apis \
  "kubermatic:v1 apps.kubermatic:v1" \
  --go-header-file /tmp/headerfile

# move generated code to the correct location; this should work regardless where
# this repository has been cloned to
mv $GOPATH/src/k8c.io/api/v3/pkg/generated pkg/

# in case the repository was cloned to the module path in $GOPATH, make sure to
# remove the leftover v3 directory
rm -rf v3

# cleanup
rm -rf vendor

# generate CRDs from the Go types
echodate "Generating CRDs"
go run sigs.k8s.io/controller-tools/cmd/controller-gen \
  crd \
  object:headerFile=./hack/boilerplate/ce/boilerplate.go.txt \
  paths=./pkg/... \
  output:crd:dir=./$CRD_DIR

# beautify CRDs just because we can
for f in $CRD_DIR/*.yaml; do
  yq --inplace --no-doc 'del(.metadata.creationTimestamp)' "$f"
  mv "$f" "$f.bak"
  echo -e "# This file has been generated by hack/update-codegen.sh, DO NOT EDIT.\n" > "$f"
  cat "$f.bak" >> "$f"
  rm "$f.bak"
done

annotation="kubermatic.k8c.io/location"
locationMap='{
  "applicationdefinitions.apps.kubermatic.k8c.io": "master,seed",
  "applicationinstallations.apps.kubermatic.k8c.io": "usercluster",
  "addonconfigs.kubermatic.k8c.io": "master",
  "addons.kubermatic.k8c.io": "seed",
  "admissionplugins.kubermatic.k8c.io": "master",
  "alertmanagers.kubermatic.k8c.io": "seed",
  "allowedregistries.kubermatic.k8c.io": "master",
  "clusters.kubermatic.k8c.io": "seed",
  "clustertemplateinstances.kubermatic.k8c.io": "seed",
  "clustertemplates.kubermatic.k8c.io": "master,seed",
  "constraints.kubermatic.k8c.io": "master,seed",
  "constrainttemplates.kubermatic.k8c.io": "master,seed",
  "customoperatingsystemprofiles.operatingsystemmanager.k8c.io": "seed",
  "etcdbackupconfigs.kubermatic.k8c.io": "seed",
  "etcdrestores.kubermatic.k8c.io": "seed",
  "externalclusters.kubermatic.k8c.io": "master",
  "groupprojectbindings.kubermatic.k8c.io": "master",
  "ipamallocations.kubermatic.k8c.io": "master",
  "ipampools.kubermatic.k8c.io": "master",
  "kubermaticconfigurations.kubermatic.k8c.io": "master,seed",
  "kubermaticsettings.kubermatic.k8c.io": "master",
  "mlaadminsettings.kubermatic.k8c.io": "seed",
  "presets.kubermatic.k8c.io": "master,seed",
  "projects.kubermatic.k8c.io": "master,seed",
  "resourcequotas.kubermatic.k8c.io": "master",
  "rulegroups.kubermatic.k8c.io": "master",
  "seeds.kubermatic.k8c.io": "master,seed",
  "userprojectbindings.kubermatic.k8c.io": "master,seed",
  "usersshkeys.kubermatic.k8c.io": "master",
  "users.kubermatic.k8c.io": "master,seed"
}'

failure=false
echodate "Annotating CRDs"

for filename in $CRD_DIR/*.yaml; do
  crdName="$(yq '.metadata.name' "$filename")"
  location="$(echo "$locationMap" | jq -rc --arg key "$crdName" '.[$key] + ""')"

  if [ -z "$location" ]; then
    echo "Error: No location defined for CRD $crdName"
    failure=true
    continue
  fi

  yq --inplace ".metadata.annotations.\"$annotation\" = \"$location\"" "$filename"
done

if $failure; then
  exit 1
fi
