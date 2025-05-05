---
sidebar_position: 1
---
# Kubernetes

`dsenv` is built around Kubernetes as our compute platform.
We run two K8s clusters, one for development and one for production.
Each cluster is further separated into namespaces based on customers internal usage.
With that we make sure that we have clean multi-tenant capabilities and never leak development data or workloads into production.


```mermaid
---
config:
  theme: redux
---
flowchart TD
 subgraph dssys["NS: DSSys"]
        blueprinter["Blueprinter"]
        portal["Portal"]
  end
  subgraph argo["NS: Argo"]
        argocd["Argo CD"]
  end
  subgraph o11y["NS: O11y"]
        dashboard["Dashboards"]
        ingest["Ingest"]
  end
 subgraph cB["NS: Customer B"]
        cBw1["Webapp 1"]
        cBw2["Webapp 2"]
  end
 subgraph cA["NS: Customer A"]
        cAw1["Webapp 1"]
        cAw2["Webapp 2"]
  end
   subgraph k8sdev["K8s: Dev"]
        argodev["NS: Argo"]
        dssysdev["NS: DSSys"]
        o11ydev["NS: O11y"]
        cBdev["Customer B"]
        cAdev["Customer A"]
  end
 subgraph k8sprod["K8s: Prod"]
        argo
        dssys
        o11y
        cB
        cA
 end
    cBw1 --> cBw2
    argo --> cA & cB

    blueprinter --> argo & o11y

```

## Services and definitions

We use [kustomize](https://kustomize.io/) to define flexible application configuration.
When you bootstrap a new service via the `ds` CLI a basic structure is created for you.
This includes a base manifest as well as overlays for development and production.
`ds` provides sensible defaults for most use-cases but you can further adapt the manifests for your specific needs.

Where needed, we use [Helm](https://helm.sh/) for packaging and dependency management.
However, we try to not rely on Helm for customizations and templating. Wherever possible we use kustomize as the outcomes are easier to predict.

## Workloads

We aim to only run compute workloads in k8s. All persistent data is usually stored outside of the clusters.
We provide [managed databases](/docs/architecture/databases) as well as S3 compatible storage if your application needs it.