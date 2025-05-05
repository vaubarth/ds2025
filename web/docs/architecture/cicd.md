---
sidebar_position: 7
---
# Code & CI/CD

## Code
We use GitHub to host our repositories.
The repository functions as the single source of truth for all configuration.

This includes k8s manifests, backstage manifests but also observability definitions like dashboards.

:::info
Only changes to the code can trigger actions like releases or configuration changes.
:::

## CI
GitHub actions are used to build, test and integrate any changes.
Every part of a pipeline is encapsulated in its own action.
This way we ensure that the actual workflows are simple and don’t contain unnecessary logic. We also deduplicate individual steps across services and teams.

## CD
For deployment we use [ArgoCD](https://argo-cd.readthedocs.io). 
Argo makes sure that we can easily automate the rollout to different clusters while maintaining versions in code.

All services have have automatic sync enabled, therefore updates in the manifests are automatically applied and the CD pipelines stay small.

In a typical deployment workflow we create a new git tag and build a corresponding docker image.
We then update a kustomization file declaring the image new tag to be used in the deployment. Argo will sync this change across the relevant clusters and therefore update the service.