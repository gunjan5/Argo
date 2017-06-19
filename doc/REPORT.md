
# This is a draft report. See [UnikernelsforServerlessArchitecture.pdf](UnikernelsforServerlessArchitecture.pdf) for the full report.

## Motivation:
Motivation for doing this research is to compare boot up time performance of containers vs. unikernels.
Boot up times are very crucial for Serverless/FaaS (Function as a Service) architecture. Moving from VMs
to containers enabled us to expand rapidly with Microservices architecture. Now, to take it one step farther
Serverless/FaaS architecture breaks down services even grannular and runs instances of them, those instances
only boot up when they are needed, so boot up time is one of the most critical aspect for better performance. 

## Problem:
![Bootup time problem with Serverless](/doc/img/fig-1-res-vs-time.png "Resources vs Time")
[ Figure 1: Resource utilization vs time comparison for always running system compared with Serveless system]

- The problem with having an always running service, an application server for example, is that it racks up a
huge compute, power, memory, network and storage bill from the cloud providers since they usually tend to charge
for these resources per unit of time. Serverless architecture is designed to overcome this issue by bringing up
the service only when it's needed and then powering it down to save the resources. This solves part of the problem
where it's not utilizing the resources when the service is not in use, but introduces another problem which
is the service instance will need to be booted up whenever it's the service is needed, and this introduces
extra latency in serving the requests.
- Boot up times have increased radically since the introduction of containers, compared to Virtual Machines.
However, containers still take some time to boot up, and request response time is still not as good as an 
always running service. 
- This request response time can be improved further if we use unikernels to host such service in Serverless
architecture.

## Intent:
Intent of this experiment is to see if unikernels can boot up faster compared to containers.
Instances boot up per request in Serverless/FaaS architecture.

The goals for this experiment are:

- Get a unikernel runtime working with a serverless/FaaS orchestrator (doesn't exist right now)
- Measure boot up time of unikernels vs containers (for the same code serving a simple HTTP service)
- Measure total size of unikernel instance vs container instance for a simple server (same code)



