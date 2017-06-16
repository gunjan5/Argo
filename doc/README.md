# Running status of the meetings:

## 4/6 Thursday:

### Discussion:
- What is the scope of the project
- What are we trying to prove
- What are the possible benchmarks

### Tasks:
- Define A vs B (what are we comparing)
- Experimental design
- Find 2 related papers

### Work:
- http://media.taricorp.net/performance-evaluation-unikernels.pdf
- (not really a paper) https://www.slideshare.net/BodenRussell/performance-characteristics-of-traditional-v-ms-vs-docker-containers-dockercon14


## 4/13 Thursday:

### Discussion:
- Explain the hypothesis
- Defend the hypothesis

### Tasks:
- add "motivation" (specific numbers, boot up time)
- intent (new doc for the final report)
- Add `syscall` to the project

### Work:
- added a diagram, motivation, and intent section draft
- method:
  - run a basic compiled program directly on the OS
  - now run it inside a docker container and measure the startup + code run time
  - now run it inside the container but measure startup + code run + tear-down time
  - now just "start" an already created container instance (real startup time + code run time)
  // need to figure out the details for unikernel


## Week 5:
  - separate dev and test environment
  - define the environment with unikernels (on a VM? hypervisor? baremetal?)
  - ubuntu on metal, not 
  - how is this different from exokernel/user runtimes model
  - read https://pdos.csail.mit.edu/6.828/2008/readings/engler95exokernel.pdf


## Week 6: 
- low variance between cold and hot run
- plot them and see the difference


## Week 7:
watch https://www.youtube.com/watch?v=57g9hNGGOuk&feature=youtu.be

- numebrs 
- experimental setup
- short abstract
- short summery
