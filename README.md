# eval-yaml-diff
eval-yaml-diff is a command-line tool that compares the differences between two YAML files (multi-document) and validates them against predefined policies (in YAML format).

## Features

- Compares the differences between two YAML files (multi-document) and identifies the paths where differences occurred along with the change type (add/delete/change).
- Validates the differences against predefined policies and returns an exit code indicating if there are any policy violations.
  - specifying a policy configuration file using the `--config` option.

## Installation

### Using Homebrew

1. Open a terminal and run the following commands:

    ```shell
    $ brew install yashirook/tap/eval-yaml-diff
    ```

### Using tar.gz distribution
1. Download the appropriate version of the tar.gz archive from the releases page.

2. Extract the archive by running the following command:

    ```shell
    $ tar -zxvf <archive-filename>.tar.gz
    ```
3. Navigate to the extracted directory:

    ```shell
    $ cd <extracted-directory>
    ```


## Usage
To see the usage instructions for the CLI tool, run the following command:

```shell
$ yaml-diff-checker --config config.yaml example/base.yaml example/new.yaml
PATH                                            CHANGE_TYPE     RESULT          
.spec.ports[0].port                             change          DENIED
.spec.template.metadata.labels.version          add             ALLOWED
.spec.template.spec.containers[0].image         change          ALLOWED
.spec.template.spec.containers[0].ports         change          DENIED
.spec.replicas                                  change          DENIED

# in case of exist DENIED diff, return exit status(2).
$ echo $?
2
```

### Sample config and compare files

`config.yaml`
```yaml:config.yaml
allowedPolicies:
  - path: .spec.template.metadata.labels.version
    changeType: add
    recursive: false
  - path: .spec.template.spec.containers[0].image
    changeType: change
    recursive: false
```


`base.yaml`
```yaml:base.yaml
---
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app: MyApp
  ports:
  - protocol: TCP
    port: 80
    targetPort: 9376
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app-deployment
  labels:
    app: MyApp
spec:
  replicas: 3
  selector:
    matchLabels:
      app: MyApp
  template:
    metadata:
      labels:
        app: MyApp
    spec:
      containers:
      - name: app
        image: my-app:1.0.0
        ports:
        - containerPort: 9376
```

`new.yaml`
```yaml:new.yaml
---
metadata:
  name: my-service
spec:
  ports:
  - protocol: TCP
    targetPort: 9376
    port: 8080
  selector:
    app: MyApp
apiVersion: v1
kind: Service
---
apiVersion: apps/v1
kind: Deployment
spec:
  selector:
    matchLabels:
      app: MyApp
  replicas: 10
  template:
    metadata:
      labels:
        app: MyApp
        version: 0.0.1
    spec:
      containers:
      - name: app
        ports:
        - containerPort: 9376
        - secondPort: 1010
        image: my-app:1.1.0
metadata:
  name: app-deployment
  labels:
    app: MyApp
```