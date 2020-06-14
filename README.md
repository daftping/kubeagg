# Kubectl Aggregate

![](https://github.com/daftping/kubeagg/workflows/Go/badge.svg) ![Release](https://github.com/daftping/kubeagg/workflows/Release/badge.svg)

Kubectl wrapper to run against multiple contexts and namespaces.

kubeagg features which are not available in kubectl:

- get objects from multiple context (as list or as pattern)
- get objects from multiple namespaces (as list or as pattern)

## Usage

```bash
Kubectl wrapper can get any objects in any cluster in any namespace.

        You can provide list of contexts or context-pattern (regexp)
        If object you are trying to get namespaced you can provide list of 
        namespaces or namespace-pattern (regexp)

Usage:
  kubeagg [command]

Examples:
        // Get namespaces from all contexts in (kubectl config view)
        kubeagg get ns

        // Get pods from contexts matched 'dev$|test$' regexp and in 
        // namespaces matched 'default|test|dev' regexp
        kubeagg \
                --context-pattern='dev$|test$' \
                --namespace-pattern='default|test|dev' \
                get pod

        // Get all nodes in "docker-desktop" and "test" contexts
        kubeagg --contexts=docker-desktop,test get no

Available Commands:
  get         Kubectl wrapper to run against multiple contexts and namespaces
  help        Help about any command

Flags:
      --context-pattern string     Send request to contexts matched provided regexp.
                                   Ignored if --contexts is provided. (default ".*")
      --contexts strings           Send request to provided contexts. Has precedence over --context-pattern.
                                   (default: '', --context-pattern is used)
  -h, --help                       help for kubeagg
  -l, --loglevel string            Debug, Info, Warn, Error, Fatal (default "Error")
      --namespace-pattern string   Get objects from namespaces matched provided regexp.
                                   Ignored if --namespaces is provided. (default ".*")
  -n, --namespaces strings         List namespaces to get objects from.(default: '', --namespace-pattern is used)
      --no-headers                 Skip headers in output
  -o, --output string              Output format. Supported values: table, json. (default "table")

Use "kubeagg [command] --help" for more information about a command.
```

## Get namespaces in all contexts in ~/.kube/config

```bash
kubeagg get ns
CONTEXT         NAMESPACE       TYPE            NAME
docker-desktop  n\a             Namespace       default
docker-desktop  n\a             Namespace       docker
docker-desktop  n\a             Namespace       kube-node-lease
docker-desktop  n\a             Namespace       kube-public
docker-desktop  n\a             Namespace       kube-system
prod-a          n\a             Namespace       default
prod-a          n\a             Namespace       kube-public
prod-a          n\a             Namespace       kube-system
test-a          n\a             Namespace       default
test-a          n\a             Namespace       kube-public
test-a          n\a             Namespace       kube-system
test-b          n\a             Namespace       default
test-b          n\a             Namespace       kube-public
test-b          n\a             Namespace       kube-system
```

## Get pod from contexts matched pattern and namespaces matched pattern

```bash
kubeagg \
  --context-pattern='^docker-desktop$|^test' \
  --namespace-pattern='default|-public$' \
  get pod
CONTEXT         NAMESPACE       TYPE    NAME
docker-desktop  default         Pod     kubeagg
docker-desktop  default         Pod     kubeagg-68898df575-hcvs5  
test-a          default         Pod     kubeagg2
test-a          kube-public     Pod     kubeagg3
```
