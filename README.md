# Run scheduled tasks with cron scheduler expression

> [!IMPORTANT]  
> **Update , April 2026:** This repository is no longer maintained. Native support for cron for batch jobs will be implemented in Radix.

* Docker image is rootless
* Scheduler can run while the next command is running
* [Dockerfile.echo](./Dockerfile.echo) and [./Dockerfile.python](Dockerfile.python) are examples to get the scheduler and run it within another base image
* Run locally from docker: `docker run -it ghcr.io/equinor/task-scheduler:latest --command "echo \"test\"" --schedule "0/5 * * * * *"`

## Contribution
Want to contribute? Read our [contributing guidelines](./CONTRIBUTING.md)

## Security
This is how we handle [security issues](./SECURITY.md)
