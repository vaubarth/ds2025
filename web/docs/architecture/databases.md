---
sidebar_position: 6
---
# Databases

We provide managed databases that can be used from any service.
- **PostgresSQL** for structured SQL data
- **MongoDB** for unstructured document data

The databases are configured to be redundant and failsafe. They are also kept on new versions automatically.

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
 subgraph cB["NS: Customer B"]
        cBw1["Webapp 1"]
        cBw2["Webapp 2"]
  end
 subgraph cA["NS: Customer A"]
        cAw1["Webapp 1"]
        cAw2["Webapp 2"]
  end
 subgraph k8s["K8s"]
        dssys
        cB
        cA
  end
    db["Managed DB"]
    db@{ shape: cyl}

    cBw2 --> db
    cAw1 --> db
    cAw2 --> db
    cBw1 --> cBw2

    blueprinter --> db

```

## Usage

In order to use a Postgres/Mongo database with your service we need to create a new database within it.
This is usually done via `ds`. You will receive all required secrets and connection data.