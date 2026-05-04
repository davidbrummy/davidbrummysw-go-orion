# davidbrummysw-go-orion

## Local Postgres

Create the shared Docker network once:

```sh
./createNetwork.sh
```

Start Postgres:

```sh
make db-up
```

The default local connection string is:

```sh
postgres://orion:orion@localhost:5432/orion?sslmode=disable
```

Override it with `DATABASE_URL` when needed. The Docker dev environment uses the `DATABASE_URL` in `dev.env`.

Stop Postgres:

```sh
make db-down
```

To run as K8s

minikube start --driver=docker
minikube addons enable ingress


# build the image, load it into minikube. 
# deploy image as a pod
docker build . -t davidbrummysw-go-orion-service1.0
minikube image load davidbrummysw-go-orion-service1.0
minikube image ls --format table
kubectl apply -f k8-orion-deployment.yaml

# Option 1 - Expose pod via port forwarding

kubectl get pods
kubectl port-forward pod/orion-service-deployment-55bf4f7bb6-p8pn4 8080:8080

# Option 2 - create a service and create a tunnel

kubectl expose deployment/orion-service-deployment --type="NodePort" --port 8080
minikube service  orion-service-deployment

# Option 3 - Deploy as ingress

sudo minikube tunnel
kubectl expose deployment/orion-service-deployment --type="NodePort" --port 8080
kubectl apply -f ingress.yaml

minikube dashboard



minikube addons enable ingress
