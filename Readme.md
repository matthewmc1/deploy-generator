# Deployment Configuration Generation

Creating Deployment manifests for K8s requires knowledge of current best pracitces and we want to simplify this by using a reference template a smaller set of values you need to define.

```yaml
name: app
highAvailability: true 
image: <link>
labels: |
 "service-type:frontend"
 "application:client-facing"
```

You can then pass this on build through `deploy` the deploy command will have some options:

- version - lists the current version
- log - outputs the log from what has been created
- create - creates a deployment specification based on the last template and accepts `args...` for as many deployments that you want to generate

