#!/usr/bin/env bash

bin_dir=$1/bin/

mkdir -p ${bin_dir}
GOBIN=${bin_dir} go install sigs.k8s.io/controller-runtime/tools/setup-envtest@v0.0.0-20230216140739-c98506dc3b8e
