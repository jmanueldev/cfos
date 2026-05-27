PROTO_DIR=api/proto

proto:
	buf generate

scheduler:
	cd control-plane && go run ./scheduler

runtime:
	cd runtime && cargo run

ml:
	cd ml && python rl/train.py

dev-up:
	docker compose up -d

lint:
	golangci-lint run

test:
	go test ./...

build:
	docker compose build