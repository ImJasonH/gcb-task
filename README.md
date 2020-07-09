# Google Cloud Build for Tekton

This repo provides an experimental [Tekton Custom
Task](https://github.com/tektoncd/community/pull/128) that executes a Google
Cloud Build.

The intention is to demonstrate the kinds of things a Custom Task can do, and
to demonstrate how to write a Custom Task.

## Install

Install and configure `ko`.

```
ko apply -f config/
```

This will build and install the controller on your cluster, in the namespace
`cloudbuild-task`.

## Service Account Setup

Builds need to be created by a Service Account with appropriate permissions.

[Create a GCP Service
Account](https://cloud.google.com/kubernetes-engine/docs/tutorials/authenticating-to-cloud-platform#step_3_create_service_account_credentials)
with the **Cloud Build Editor** role, create and download a key file for that
SA, and create a K8s Secret with that Service Account Key:

```
kubectl create secret generic sa-key --from-file=key.json=PATH-TO-KEY-FILE.json -n cloudbuild-task
```

## Run a Build

Create a `Run` that refers to a Build:

```
$ kubectl create -f gcb-run.yaml 
run.tekton.dev/gcb-run-j2w5p created
$ kubectl get runs -w
NAME            SUCCEEDED   REASON         STARTTIME   COMPLETIONTIME
gcb-run-j2w5p   Unknown     BuildWorking   15s         
gcb-run-j2w5p   True        BuildSucceeded   31s         1s
```

# Outstanding Work

Thanks!😊 There's plenty more to do though.

Namely:

* Don't hard-code a trivial GCB build, either let users define their build
  using a CRD, and/or read a build config from a repo or parameter value.
* Accept a source archive in GCS as an input parameter, to integrate with
  other Tasks in a Pipeline.
* Emit the build ID as an output result.
