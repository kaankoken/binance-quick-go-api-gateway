proto:
	bash ./scripts/generate_protos.sh

server-dev:
	cp ./config/envs/dev.env ./config/envs/.env
	go run cmd/main.go 

server-prod:
	cp ./config/envs/prod.env ./config/envs/.env
	go run cmd/main.go

run-vendor:
	bash ./scripts/run_vendor.sh
