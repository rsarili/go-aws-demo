STACK_DIR = ./infrastructure/go-playground-stack

.PHONY: deploy
deploy:
	cd $(STACK_DIR) && cdk deploy


.PHONY: build
build:
	cd functions/hello-world  && GOARCH=amd64 GOOS=linux go build -o dist/bootstrap main.go

.PHONY: run-function
run-function:
	aws lambda invoke --function-name hello-world --cli-binary-format raw-in-base64-out --payload '{"message": "hello world"}' --log-type Tail response.json