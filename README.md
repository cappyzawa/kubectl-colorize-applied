# kubectl-colorize-applied

Enjoy coloring the applied results.

![](https://github.com/cappyzawa/demo/raw/master/kubectl-colorize-applied/image.png)

## How to install

```shell
kubectl krew install colorize-applied
```

## How to use

This plugin colorizes the results of apply/dry-run.  
Each line containing the following strings is colored: `pruned`, `configured`, `created`, and `unchanged`.

```shell
kubectl apply --prune -f manifest.yaml --dry-run=server | kubectl colorize-applied
```
