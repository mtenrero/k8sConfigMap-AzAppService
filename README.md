# k8sconfigmap-azappservice

This CLI executable program, transform a given Kubernetes ConfigMap into a JSON object ready to consume by Azure
AppService container as environment variables.

It's quite useful for avoiding to maintain multiple config files if you have hybrid deployment with kubernetes and 
Azure App Service containers.

## Usage

Download the latest binary for your platform, (its written in go and doesn't need any dependency, it's all compiled!)

Execute it in your CI/CD pipeline:

```bash
./k8sconfigmap-azappservice -path <k8s-configmap-file> -out <json-output-file>
```

## Contributing 

PR! are accepted. Its planned to improve the program for supporting multiple ConfigMaps in the same file and the ability
to specify environment-scoped variables for multiple apps deployment inside the same Azure App Service.
