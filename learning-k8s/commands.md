## Switch between clusters

```
kubectl config current-context
kubectl config get-contexts
kubectl config use-context <CONTEXT_NAME>
```

## Accessing pods and debugging from inside

### Enter the DockerDesktop VM
```
docker run --net=host --ipc=host --uts=host --pid=host -it --security-opt=seccomp=unconfined --privileged --rm -v /:/host alpine chroot /host
```
- The container is created from the alpine image.
- The --net, --ipc, --uts and --pid flags make the container use the host’s namespaces instead of being sandboxed, and the --privileged and --security-opt flags give the container unrestricted access to all sys-calls.
- The -it flag runs the container interactive mode and the --rm flags ensures the container is deleted when it terminates.
- The -v flag mounts the host’s root directory to the /host directory in the container. The chroot /host command then makes this directory the root directory in the container

### SSH into a container
```
kubectl exec -it kiada -- bash
kubectl exec -it <POD_NAME> -c <CONTAINER_NAME> -- bash
```

### Run debugger side-pod (https://medium.com/@the_good_guy/get-shell-access-to-pods-nodes-in-kubernetes-using-kubectl-1d8fc10e89eb)
```
kubectl debug node/<node-name> -it --image=<image name>
kubectl run --image=curlimages/curl -it --restart=Never --rm client-pod curl 10.1.0.94:8080
```

### Port-forwarding
```
kubectl port-forward <POD_NAME> <PORT>
kubectl port-forward <POD_NAME> <PORT1> <PORT2> <...>
kubectl port-forward <POD_NAME> <LOCAL_PORT>:<POD_PORT>
```

### Attach to pod (std in/out/err)
```
kubectl attach <POD_NAME>
```