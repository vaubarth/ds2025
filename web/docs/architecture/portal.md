---
sidebar_position: 3
---
# Developer Portal

We use [Backstage](https://backstage.io) as our developer portal.
Here you can view all services and customers. 
Services integrate with different parts of our system, like k8s and alerts from Grafana.

This way the portal functions as a central UI for all teams. 

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

 subgraph k8s["K8s"]

        dssys
  end


    cli["CLI"]
    cli@{ shape: rounded}

    repo["GitHub repo"]
    repo@{ shape: h-cyl}

    cli --> dssys


    portal --> repo


```

## Backstage manifests

Each service needs to have its own Backstage manifest. These are persisted in the GitHub repo of the service.
Backstage automatically syncs any changes back from GitHub.

An initial manifest is created when you bootstrap your service through `ds`.