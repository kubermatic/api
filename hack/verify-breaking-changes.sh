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

if [ -z "${PULL_BASE_SHA:-}" ]; then
  echodate "This script is meant to run as a Prowjob."
  exit 1
fi

oasdiff=oasdiff
if ! [ -x "$(command -v $oasdiff)" ]; then
  version=2.0.1
  url="https://github.com/Tufin/oasdiff/releases/download/v$version/oasdiff_${version}_linux_amd64.tar.gz"
  oasdiff=/tmp/oasdiff

  echodate "Downloading oasdiff v$version..."
  wget -O- "$url" | tar xzOf - oasdiff > $oasdiff
  chmod +x $oasdiff
  echodate "Done!"
fi

apiVersion=v1
mergedHead="$(git rev-parse HEAD)"
directories=(
   crd/community/
   crd/enterprise/kcp/
   crd/enterprise/seed/
)

echodate "Base commit: $PULL_BASE_SHA ($PULL_BASE_REF)"
echodate "Revision...: $(git rev-parse HEAD)"

exitCode=0

for dir in "${directories[@]}"; do
  echodate "Checking for breaking changes in $dir…"

  echodate "Combining base CRDs into spec…"
  git checkout --quiet "$PULL_BASE_SHA" -- "$dir"
  create_openapi_from_crds "$dir" base.json "$apiVersion"

  echodate "Combining revision CRDs into spec…"
  git reset --hard --quiet
  create_openapi_from_crds "$dir" revision.json "$apiVersion"

  echodate "Detecting breaking changes…"
  $oasdiff breaking base.json revision.json --format=json > breaking.json

  if [[ $(jq 'length' breaking.json) -eq 0 ]]; then
    echo "No breaking changes."
    continue
  fi

  # group output by API group
  apiGroups="$(oasdiff_breaking_apigroups breaking.json)"

  for apiGroup in $apiGroups; do
    heading "$apiGroup"
    echo

    # filter all breaking changes down to this API group
    oasdiff_breaking_changes breaking.json "$apiGroup" > "breaking.$apiGroup.json"

    # get affected CRDs
    for apiPath in $(oasdiff_breaking_paths "breaking.$apiGroup.json"); do
      filename="$(basename $apiPath)"
      crdFile="$dir/$filename"

      if [ -f "$crdFile" ]; then
        kind="$(yq e '.spec.names.kind' $crdFile)"

        subheading "$kind has breaking changes"
        echo
        oasdiff_breaking_markdown_list "breaking.$apiGroup.json" "$apiPath"
      else
        subheading "$kind was removed"
      fi
      echo
    done

    echo
  done

  exitCode=1
done

exit $exitCode
