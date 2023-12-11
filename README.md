# Run scheduled tasks with cron scheduler expression

* Docker image is rootless
* Scheduler can run while the next command is running
* [Dockerfile.echo](./Dockerfile.echo) and [./Dockerfile.python](Dockerfile.python) are examples to get the scheduler and run it within another base image
* Run locally from docker: `docker run -it ghcr.io/equinor/task-scheduler:latest --command "echo \"test\"" --schedule "0/5 * * * * *"`