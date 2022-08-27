build:
	cp ./config/envs/prod.env ./config/envs/config.env
	go build main.go

lint:
	golint ./...

proto:
	bash ./scripts/generate_protos.sh

run-vendor:
	bash ./scripts/run_vendor.sh

run-vet:
	go vet ./...

server-dev-debug:
	bash ./scripts/change_flavor.sh --flavor dev --debug
	go run main.go 

server-dev-release:
	bash ./scripts/change_flavor.sh --flavor dev --release
	go run main.go 

server-prod-debug:
	bash ./scripts/change_flavor.sh --flavor prod --debug
	go run main.go

server-prod-release:
	bash ./scripts/change_flavor.sh --flavor prod --release
	go run main.go

static-check:
	staticcheck ./...

test:
	bash ./scripts/run_tests.sh
