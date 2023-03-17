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

EXIT_CODE=0

try() {
  local title="$1"
  shift

  heading "$title"
  echo -e "$@\n"

  start_time=$(date +%s)

  set +e
  $@
  exitCode=$?
  set -e

  elapsed_time=$(($(date +%s) - $start_time))
  TEST_NAME="$title" write_junit $exitCode "$elapsed_time"

  if [[ $exitCode -eq 0 ]]; then
    echo -e "\n[${elapsed_time}s] SUCCESS :)"
  else
    echo -e "\n[${elapsed_time}s] FAILED."
    EXIT_CODE=1
  fi

  git reset --hard --quiet
  git clean --force

  echo
}

try "Verify go.mod" make check-dependencies
try "Verify boilerplate" ./hack/verify-boilerplate.sh
try "Verify code generation" ./hack/verify-codegen.sh
try "Verify license compatibility" ./hack/verify-licenses.sh

exit $EXIT_CODE
