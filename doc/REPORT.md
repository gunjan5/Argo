
## Motivation:
Motivation for doing this research is to compare boot up time performance of containers vs. unikernels.
Boot up times are very crucial for Serverless/FaaS (Function as a Service) architecture. Moving from VMs
to containers enabled us to expand rapidly with Microservices architecture. Now, to take it one step farther
Serverless/FaaS architecture breaks down services even grannular and runs instances of them, those instances
only boot up when they are needed, so boot up time is one of the most critical aspect for better performance. 

## Problem:
![Bootup time problem with Serverless](/doc/img/fig-1-res-vs-time.png "Resources vs Time")

## Intent:
Intent of this experiment is to see if unikernels can boot up faster compared to containers.
Instances boot up per request in Serverless/FaaS architecture.

The goals for this experiment are:

- Get a unikernel runtime working with a serverless/FaaS orchestrator (doesn't exist right now)
- Measure boot up time of unikernels vs containers (for the same code serving a simple HTTP service)
- Measure total size of unikernel instance vs container instance for a simple server (same code)
