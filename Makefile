.PHONY: deploy-pipeline
deploy:
	docker build . -t ghcr.io/equinor/task-scheduler:latest
	docker push ghcr.io/equinor/task-scheduler:latest
