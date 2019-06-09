## Helm chart for mtgs

## Prerequisites

- Kubernetes 1.4+ with Beta APIs enabled
- [helm](https://helm.sh)
- PV provisioner support in the underlying infrastructure (with persistence storage enabled)
- LoadBalancer support in the underlying infrastructure

## Installing the Chart

To install the chart:

```console
$ helm dependency update
$ helm install --name mtgs --wait .
```

> **Tip**: List all releases using `helm list`

## Updating the Chart

To update the chart:

```console
$ helm dependency update
$ helm upgrade mtgs --wait .
```

## Uninstalling the Chart

To uninstall/delete:

```console
$ helm delete mtgs
```

### The following table lists the configurable parameters of the Mtgs chart and their default values.

| Parameter           | Description                                 | Default            |
|---------------------|---------------------------------------------|--------------------|
| replicaCount        | Number of MTPROTO replicas                  | 3                  |
| image.repository    | Image name                                  | "mazy/mtgs"        |
| image.tag           | Image tag                                   | "latest"           |
| image.pullPolicy    | Image pull policy                           | "IfNotPresent"     |
| service.ports       | Public ports for MTPROTO                    | {22,443,1194,3128} |
| adtag               | Tag for promoted channel                    | ""                 |
| ingress.apiToken    | Token for API secure(for Authorized header) | "secret"           |
| ingress.annotations | Ingress annotations                         | {}                 |
| ingress.hosts       | Ingress hosts                               | {"example.com"}    |
| ingress.path        | Ingress path                                | "/mtgs"            |
| ingress.tls         | Secrets for tls                             | {}                 |
| consul.Replicas     | Number or consul replicas                   | 3                  |
| consul.Resources    | Consul Resources                            | {}                 |
| resources           | Mtgs Resources                              | {}                 |