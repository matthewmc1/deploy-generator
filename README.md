## Service Generator 

Deploying and building k8s applications takes time, and we aim to simplfy this by stripping it down to the things you care about, image, name and high-availability.

We take care of the rest, ensuring it conforms to the latest security standards, topology awareness, and handling rolling updates plus resource management.

sample:

```yaml
name: notes
highAvailability: true 
image: "http:///link.com"
labels: 
 application: "client-facing"
 service-type: "frontend"
```

output:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Deployment.name }}
spec:
  replicas: {{ if .Deployment.highAvailability == true : 3 else 1}}
  selector:
    matchLabels:
        {{ foreach line in .Deployment.labels }}
  template:
    metadata:
      labels:
        {{ foreach line in .Deployment.labels }}
    spec:
      topologySpreadConstraints:
      - maxSkew: 1 
        topologyKey: topology.kubernetes.io/zone  
        whenUnsatisfiable: DoNotSchedule  
      securityContext:
        readOnlyRootFilesystem: true
      containers:
      - name: {{ .Deployment.name }}
        image: {{ .Deployment.image }}
        securityContext:
          runAsUser: 1000  
          capabilities:
            drop: ["ALL"]  
        resources:
          requests:
            memory: "50Mi"
            cpu: "100m"
          limits:
            memory: "200Mi"
```

