build:
	cp ./config/envs/prod.env ./config/envs/config.env
	go build cmd/main.go

proto:
	bash ./scripts/generate_protos.sh

server-dev-debug:
	bash ./scripts/change_flavor.sh --flavor dev --debug
	go run cmd/main.go 

server-dev-release:
	bash ./scripts/change_flavor.sh --flavor dev --release
	go run cmd/main.go 

server-prod-debug:
	bash ./scripts/change_flavor.sh --flavor prod --debug
	go run cmd/main.go

server-prod-release:
	bash ./scripts/change_flavor.sh --flavor prod --release
	go run cmd/main.go

run-vendor:
	bash ./scripts/run_vendor.sh
