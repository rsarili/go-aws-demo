STACK_DIR = ./infrastructure/go-playground-stack

.PHONY: deploy
deploy:
	cd $(STACK_DIR) && cdk deploy