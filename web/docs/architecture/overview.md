---
sidebar_position: 0
---
# Overview

`dsenv` is the overall system designed to self-host and run your applications.
It consists of several connected services as well as a cli tool called `ds`.

`dsenv` is built on [kubernetes](/docs/architecture/k8s) and uses namespaces to provide separation and isolation between customers.
There are three additional namespaces configured that hold services for [observability](/docs/architecture/o11y), [CI/CD via ArgoCD](/docs/architecture/cicd) and internal services.

Users interact primarily via our dedicated CLI `ds` with their services as well as our internal ones.
Outside of k8s we are running [managed database services](/docs/architecture/databases), currently for PostgreSQL and MongoDB, as well as databases for observability.

The internal services are the [developer portal](/docs/architecture/portal) and a service called [Blueprinter](/docs/architecture/blueprinter). The developer portal integrates with [GitHub](/docs/architecture/cicd) which is also our source of truth for all configurations.

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
 subgraph k8s["K8s"]
        argo
        dssys
        o11y
        cB
        cA
  end
    db["Managed DB"]
    db@{ shape: cyl}

    o11ydb["O11y DB"]
    o11ydb@{ shape: cyl}

    cli["CLI"]
    cli@{ shape: rounded}

    repo["GitHub repo"]
    repo@{ shape: h-cyl}

    cli --> dssys

    cBw2 --> db & ingest
    cAw1 --> db & ingest
    cAw2 --> db & ingest
    cBw1 --> cBw2
    
    o11y --> o11ydb

    repo <--> argo

    portal --> repo

    argo --> cA & cB

    blueprinter --> argo & repo & o11y & db

```