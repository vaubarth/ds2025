---
sidebar_position: 4
---
# Services

A service represents a single application, like a webapp. It is always associated with a customer.

Each service has its own:
- k8s and kustomize manifests
- dashboards in our Grafana instance
- managed database if needed (MongoDB or PostgresSQL)
- CI/CD pipelines
- GitHub repository
- Space in our developer portal

## Listing services
To list all services with `ds` run

```bash
ds services

> Listing all known services
+----+-----------+------------+------------------------+
| ID | SERVICE   | CUSTOMER   | REPOSITORY             |
+----+-----------+------------+------------------------+
|  1 | Service A | Customer A | github.com/some/repo_a |
|  2 | Service B | Customer A | github.com/some/repo_b |
|  3 | Service C | Customer A | github.com/some/repo_c |
|  4 | Service Y | Customer B | github.com/some/repo_d |
|  5 | Service Z | Customer C | github.com/some/repo_e |
+----+-----------+------------+------------------------+
```

## Getting details about a specific service
You can get details such as all relevant URLs and some statistics about every service.

```bash
ds services details abc

>Details for service abc

Repository: github.com/some/repo
Dashboard:  dash.ds.dev/a17b
Pipelines:  argo.ds.dev/bjk7

Replicas:   4
Endpoints:  some.prod.com | some.dev.com
```

# Adding a new service

Creating a new service will set up everything that is needed to run and operate the service.

```bash
ds services new NewService --customer SomeCustomer

>Adding new service NewService
Creating repo ... done! [100 in 1.002s]
Setting up database ... done! [100 in 1.002s]
Setting up o11y ... done! [100 in 1.002s]
Writing manifests ... done! [100 in 1.002s]
Setting up deployments ... done! [100 in 1.002s]
Configuring Backstage ... done! [100 in 1.002s]
Creating CI/CD configurations (GH Actions) ... done! [100 in 1.002s]
```

## Flags
There are a handful of flags that can be added to customize how a service is created.


- `--customer` *required. Specifies a valid customer to which the service is added
- `--database` The type of database that should be added. Can be specified multiple times.
- `--stack` The tech and language stack for the service. Initiates different defaults for CI/CD
- `--dns` If a dns entry is required the domain name.
- `--replicas` Number of replicas if defaults are not sufficient.