---
sidebar_position: 3
---
# Developer Portal

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