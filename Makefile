TFFLAGS = -var "do_api_token=$(DO_API_TOKEN)"

run: install
	terraform apply $(TFFLAGS)

install: .terraform
	terraform init

stop: install
	terraform destroy $(TFFLAGS)

.PHONY: install run stop
