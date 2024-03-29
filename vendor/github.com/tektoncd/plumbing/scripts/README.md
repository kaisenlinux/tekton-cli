# Helper scripts

This directory contains helper scripts used by Prow test jobs, as well as
local development scripts.

## Using the `presubmit-tests.sh` helper script

This is a helper script to run the presubmit tests. To use it:

1. Source this script.

1. [optional] Define the function `build_tests()`. If you don't define this
   function, the default action for running the build tests is to:

   - check markdown files
   - run `go build` on the entire repo
   - run `/hack/verify-codegen.sh` (if it exists)
   - check licenses in all go packages

   The markdown link checker tool doesn't check `localhost` links by default.
   Its configuration file, `markdown-link-check-config.json`, lives in the
   `test-infra/scripts` directory. To override it, create a file with the same
   name, containing the custom config in the `/test` directory.

   The markdown lint tool ignores long lines by default. Its configuration file,
   `markdown-lint-config.rc`, lives in the `test-infra/scripts` directory. To
   override it, create a file with the same name, containing the custom config
   in the `/test` directory.

1. [optional] Customize the default build test runner, if you're using it. Set
   the following environment variables if the default values don't fit your needs:

   - `DISABLE_MD_LINTING`: Disable linting markdown files, defaults to 0 (false).
   - `DISABLE_MD_LINK_CHECK`: Disable checking links in markdown files, defaults
     to 0 (false).
   - `PRESUBMIT_TEST_FAIL_FAST`: Fail the presubmit test immediately if a test fails,
     defaults to 0 (false).

1. [optional] Define the functions `pre_build_tests()` and/or
   `post_build_tests()`. These functions will be called before or after the
   build tests (either your custom one or the default action) and will cause
   the test to fail if they don't return success.

1. [optional] Define the function `unit_tests()`. If you don't define this
   function, the default action for running the unit tests is to run all go tests
   in the repo.

1. [optional] Define the functions `pre_unit_tests()` and/or
   `post_unit_tests()`. These functions will be called before or after the
   unit tests (either your custom one or the default action) and will cause
   the test to fail if they don't return success.

1. [optional] Define the function `integration_tests()`. If you don't define
   this function, the default action for running the integration tests is to run
   all `./test/e2e-*tests.sh` scripts, in sequence.

1. [optional] Define the functions `pre_integration_tests()` and/or
   `post_integration_tests()`. These functions will be called before or after the
   integration tests (either your custom one or the default action) and will cause
   the test to fail if they don't return success.

1. [optional] Define the function `conformance_tests()`. If you don't define
   this function, the default action for running the conformance tests is to run
   `go test -race -v ./test -tags=conformance` in the repo.

1. [optional] Define the functions `pre_conformance_tests()` and/or
   `post_conformance_tests()`. These functions will be called before or after the
   conformance tests (either your custom one or the default action) and will cause
   the test to fail if they don't return success.

1. Call the `main()` function passing `$@` (without quotes).

Running the script without parameters, or with the `--all-tests` flag causes
all tests to be executed, in the right order (i.e., build, then unit, then
integration, then conformance tests).

Use the flags `--build-tests`, `--unit-tests`, `--integration-tests` and `--conformance-tests`
to run a specific set of tests. The flag `--emit-metrics` is used to emit metrics when
running the tests, and is automatically handled by the default action for
integration tests (see above).

The script will automatically skip all presubmit tests for PRs where all changed
files are exempt of tests (e.g., a PR changing only the `OWNERS` file).

Also, for PRs touching only markdown files, the unit, integration and conformance tests are
skipped.

### Sample presubmit test script

```bash
source vendor/github.com/knative/test-infra/scripts/presubmit-tests.sh

function post_build_tests() {
  echo "Cleaning up after build tests"
  rm -fr ./build-cache
}

function unit_tests() {
  make -C tests test
}

function pre_integration_tests() {
  echo "Cleaning up before integration tests"
  rm -fr ./staging-area
}

# We use the default integration test runner.

main $@
```

## Using the `e2e-tests.sh` helper script

This is a helper script for Knative E2E test scripts. To use it:

1. [optional] Customize the test cluster. Set the following environment variables
   if the default values don't fit your needs:

   - `E2E_CLUSTER_REGION`: Cluster region, defaults to `us-central1`.
   - `E2E_CLUSTER_BACKUP_REGIONS`: Space-separated list of regions to retry test cluster creation in case of stockout. Defaults to `us-west1 us-east1`.
   - `E2E_CLUSTER_ZONE`: Cluster zone (e.g., `a`), defaults to none (i.e. use a regional
     cluster).
   - `E2E_CLUSTER_BACKUP_ZONES`: Space-separated list of zones to retry test cluster creation in case of stockout. If defined, `E2E_CLUSTER_BACKUP_REGIONS` will be ignored thus it defaults to none.
   - `E2E_CLUSTER_MACHINE`: Cluster node machine type, defaults to `n1-standard-4}`.
   - `E2E_MIN_CLUSTER_NODES`: Minimum number of nodes in the cluster when autoscaling,
     defaults to 1.
   - `E2E_MAX_CLUSTER_NODES`: Maximum number of nodes in the cluster when autoscaling,
     defaults to 3.

1. Source the script.

1. [optional] Write the `knative_setup()` function, which will set up your
   system under test (e.g., Knative Serving). This function won't be called if you
   use the `--skip-knative-setup` flag.

1. [optional] Write the `knative_teardown()` function, which will tear down your
   system under test (e.g., Knative Serving). This function won't be called if you
   use the `--skip-knative-setup` flag.

1. [optional] Write the `test_setup()` function, which will set up the test
   resources.

1. [optional] Write the `test_teardown()` function, which will tear down the test
   resources.

1. [optional] Write the `cluster_setup()` function, which will set up any resources
   before the test cluster is created.

1. [optional] Write the `cluster_teardown()` function, which will tear down any
   resources after the test cluster is destroyed.

1. [optional] Write the `dump_extra_cluster_state()` function. It will be
   called when a test fails, and can dump extra information about the current state
   of the cluster (typically using `kubectl`).

1. [optional] Write the `parse_flags()` function. It will be called whenever an
   unrecognized flag is passed to the script, allowing you to define your own flags.
   The function must return 0 if the flag is unrecognized, or the number of items
   to skip in the command line if the flag was parsed successfully. For example,
   return 1 for a simple flag, and 2 for a flag with a parameter.

1. Call the `initialize()` function passing `$@` (without quotes).

1. Write logic for the end-to-end tests. Run all go tests using `go_test_e2e()`
   (or `report_go_test()` if you need a more fine-grained control) and call
   `fail_test()` or `success()` if any of them failed. The environment variable
   `KO_DOCKER_REPO` and `E2E_PROJECT_ID` will be set according to the test cluster.
   You can also use the following boolean (0 is false, 1 is true) environment
   variables for the logic:

   - `EMIT_METRICS`: true if `--emit-metrics` was passed.

   All environment variables above are marked read-only.

**Notes:**

1. Calling your script without arguments will create a new cluster in the GCP
   project `$PROJECT_ID` and run the tests against it.

1. Calling your script with `--run-tests` and the variable `KO_DOCKER_REPO` set
   will immediately start the tests against the cluster currently configured for
   `kubectl`.

1. You can force running the tests against a specific GKE cluster version by using
   the `--cluster-version` flag and passing a full version as the flag value.

### Sample end-to-end test script

This script will test that the latest Knative Serving nightly release works. It
defines a special flag (`--no-knative-wait`) that causes the script not to
wait for Knative Serving to be up before running the tests. It also requires that
the test cluster is created in a specific region, `us-west2`.

```bash

# This test requires a cluster in LA
E2E_CLUSTER_REGION=us-west2

source vendor/github.com/knative/test-infra/scripts/e2e-tests.sh

function knative_setup() {
  start_latest_knative_serving
  if (( WAIT_FOR_KNATIVE )); then
    wait_until_pods_running knative-serving || fail_test "Knative Serving not up"
  fi
}

function parse_flags() {
  if [[ "$1" == "--no-knative-wait" ]]; then
    WAIT_FOR_KNATIVE=0
    return 1
  fi
  return 0
}

WAIT_FOR_KNATIVE=1

initialize $@

# TODO: use go_test_e2e to run the tests.
kubectl get pods || fail_test

success
```

## Using the `deploy-release.sh` helper script

This is a helper script to deploy a Tekton release to the dogfooding cluster.

Prerequisites:

1. kubectl installed

1. cluster gke_tekton-nightly_europe-north1-a_robocat defined in the local kubeconfig

Usage:

`deploy-release.sh -p project -v version [-b bucket] [-e extra-path] [-f file] [-g post-file]`

Where:

- `project` is `pipeline`, `triggers` or `dashboard`

- `version` is the release number i.e. `v0.13.2`

- `bucket` is the URL of the bucket where the release file is stored, `gs://tekton-releases` by default

- `extra-path` is the root path within the bucket where releases are stored, empty by default

- `file` is the name of the release file, `release.yaml` by default

- `post-file` is the name of the 2nd release file, none by default, `interceptors.yaml` by default for triggers

To summarize, the deployment job will look for the release file in `<bucket>/<extra-path>/<project>/previous/<version>/<file>`
