# echo

## Usage

## Prepare

Add minikube.local to your /etc/hosts
```
$ echo "$(minikube ip) minikube.local" | sudo tee -a /etc/hosts
```

Enable ingress
```
$ minikube addons enable ingress
```

Enable heapster
```
$ minikube addons enable heapster
```