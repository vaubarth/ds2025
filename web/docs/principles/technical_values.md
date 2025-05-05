---
sidebar_position: 2
---
# Technical Values


## Standards Enhanced
We use standard tooling and process as our foundation, we don't reinvent the wheel.
Standards are combined to create seamless and tailored workflows.
We add glue-services and tools to add value. 
Workflows are integrated into the tools developers use everyday, like GitHub or Slack.

## Engineering Practices

Everything we do follows proper software engineering practices.
Every tool is well tested, has CI and has its own documentation.

We ship single binaries or docker containers - except for docker, no dependencies should be needed on a system.

## Composable workflows
We create individual, small tools that do a great job. We favor CLIs over complex GUIs
Every tool is designed to be integrated into larger flows. 
This does not mean that we force tools to be tiny, but it means that they are always composable and targeted.

## Transparent Abstraction
Tools hide complexity behind abstractions but are transparent when needed.
Users should not need to use the underlying systems like Kubernetes directly. 
We provide abstractions that work in 90% of cases.
We are aware however that abstractions are never perfect and therefore provide escape hatches to directly use the abstracted systems.