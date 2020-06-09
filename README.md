# Kubectl Aggregate
![](https://github.com/daftping/kubeagg/workflows/Go/badge.svg) ![Release](https://github.com/daftping/kubeagg/workflows/Release/badge.svg)

Kubectl wrapper to run against multiple contexts.

kubeagg features which are not available in kubectl:
- provide multiple context (literally or as pattern)


# Usage
```bash
Kubectl wrapper to run against multiple contexts

Usage:
  kubeagg get [flags]

Flags:
  -p, --context-pattern string   Send request to contexts matched provided regexp. Ignored if --contexts is provided. (default ".*")
  -c, --contexts strings         Send request to provided contexts. Has precedence over --context-pattern.(default: '', --context-pattern is used)
  -h, --help                     help for get
  -n, --namespace string         Namespace to operate in. (default "default")
  -o, --output string            Output format. Supported values: table, wide, json. (default "table")

Global Flags:
  -l, --loglevel string   Debug, Info, Warn, Error, Fatal (default "Error")
```
## All contexts in ~/.kube/config
```bash
kubeagg get ns
CONTEXT      NAME
prod-a       default
prod-a       kube-public
prod-a       kube-system
test-a       default
test-a       kube-public
test-a       kube-system
test-b       default
test-b       kube-public
test-b       kube-system
test-c       default
test-c       kube-public
test-c       kube-system
```
## List of selected namespaces
```bash
kubeagg get ns --contexts=test-a,test-b -o wide
CONTEXT      TYPE    NAME
test-a       ns      default
test-a       ns      kube-public
test-a       ns      kube-system
test-b       ns      default
test-b       ns      kube-public
test-b       ns      kube-system
```
## Regexp matched deployments
```bash
kubeagg get deploy --context-pattern='test-[b|c]$' --namespace=kube-system -o wide 
CONTEXT      TYPE    NAME
test-b       deploy  cluster-autoscaler
test-b       deploy  kube-dns
test-b       deploy  metrics-server
test-c       deploy  cluster-autoscaler
test-c       deploy  kube-dns
test-c       deploy  metrics-server

```
