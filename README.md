## ECR Token Helper

Helper for using private AWS ECR registries with Kubernetes - Generates a [Kubernetes Secrets](https://kubernetes.io/docs/concepts/configuration/secret/) resource.

Intended to be used locally or in CI environments which have access to the Kubernetes API via kubectl

### Dependencies
1. `AWS_REGION`, `AWS_ACCESS_KEY_ID`, and `AWS_SECRET_ACCESS_KEY` environment variables are set
    - The IAM user must have sufficient priviledges
    - For CI environments, it's recommended to have a user with only the **AmazonEC2ContainerRegistryReadOnly** managed policy
2. Optional - `kubectl` for direct apply
    - Needs v1.10 or greater due to this [bug](https://github.com/kubernetes/kubernetes/issues/61780)

### Build and install

> `make build && make install`

### Use
To print the Secrets YAML to stdout :
> `kube-ecr-helper get-apply --secret docker-credentials -e douglas.vaz`

Alternatively, the generated YAML can be applied to the cluster by running:
> `$(kube-ecr-helper get-apply --secret docker-credentials -e douglas.vaz) | kubectl apply -f -`

### Configuring Kubernetes image resources

Resources which have a container spec pointing to the private ECR need to refer to the created named secret via `imagePullSecrets`

Example YAML:
```yaml
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: web-app
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: web-app
    spec:
      containers:
        - name: web-app
          image: 3047566.dkr.ecr.ap-southeast-1.amazonaws.com/web-app:stable
          ports:
            - containerPort: 8000
      imagePullSecrets:
        - name: docker-credentials
```