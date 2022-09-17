# math-controller

CRDをKubernetesに適用する
```
kustomize build config/crd | kubectl apply -f -
```

Cluster Role Bind
```
kustomize build config/default | kubectl apply -f -
```

Numberを
```
kubectl apply -f config/samples
```

Controllerをデプロイする
```
go build -o bin/manager main.go
```


