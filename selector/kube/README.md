# Kube selector

The Kube selector returns the kubernetes service name as a node for every request. Any "." in the service name will be replaced with a "-".  Example, service name is micro.service.greeter, the kube selector will replace that service name with micro-service-greeter.  The kubernetes service name needs to be the same.

This DOES however require a static port assignment (because we no longer have the ability to look up metadata). This defaults to port 8080, but can be overriddden at runtime using env-vars.

An optional domain-name can be appended too.


## Environment variables

* "KUBE_SELECTOR_DOMAIN_NAME": An optional domain-name to append to the speicified service name.
* "KUBE_SELECTOR_PORT_NUMBER": Override the default port (8080) for "discovered" services.


## Usage

```go
selector := kube.NewSelector()

service := micro.NewService(
	client.NewClient(client.Selector(selector))
)
```
