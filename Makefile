STACK_DIR = ./infrastructure/go-playground-stack

.PHONY: deploy
deploy:
	cd $(STACK_DIR) && cdk deploy


.PHONY:
build:
	cd functions/hello-world  && GOARCH=amd64 GOOS=linux go build -o bootstrap main.go