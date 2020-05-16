# kubutilization
A simple dashboard for kubernetes that provides a front-end UI for node and pod utilization at 
a namespace level using the [kubernetes metrics-server](https://github.com/kubernetes-sigs/metrics-server).

## Purpose

Many Kubernetes dashboarding tools hide true utilization of 
pods behind multiple menus or in seperate UI dashboards (eg. Grafana)

This application is lightweight to run, simple to deploy and provides
a very simple view of current utilization that can be provided to users
without access to kubectl for a simple way to demonstrate current usage across nodes and pods!

Version 1.0 of this application utilizes the kubernetes metrics-server. It is mentioned in the metrics
server documentation to use something like Prometheus for accurate sourcing of usage metrics. This
application seeks to give a very high overview using the capabilities available in the metrics-server.
[Kubernetes Metrics Server Use Cases](https://github.com/kubernetes-sigs/metrics-server/blob/master/README.md#use-cases)

## Deploying

### Resource Requirements

## Contributing

### Building

### API Testing

### API Mocking for UI Testing

## License

[MIT LICENSE](./LICENSE)