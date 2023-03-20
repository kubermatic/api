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

export CGO_ENABLED ?= 0
export GOFLAGS ?= -mod=readonly -trimpath
export GO111MODULE = on
export KUBERMATIC_EDITION ?= ce

.PHONY: test
test:
	./hack/run-tests.sh

.PHONY: lint
lint:
	@# we want tagliatelle to check only CRDs
	golangci-lint run \
		--verbose \
		--config .golangci.apis.yml \
		./pkg/apis/kubermatic/... ./pkg/apis/apps.kubermatic/...

	golangci-lint run \
		--verbose \
		--print-resources-usage \
		./pkg/...

.PHONY: codegen
codegen:
	./hack/update-codegen.sh

.PHONY: verify
verify:
	./hack/verify.sh

.PHONY: check-dependencies
check-dependencies:
	go mod tidy
	go mod verify
	git diff --exit-code
