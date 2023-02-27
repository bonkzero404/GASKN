include .env
export

build:
	go mod vendor
	go build -mod vendor -o dist/gaskn main.go
	cp -r templates dist/
	mkdir dist/storages || true
	mkdir dist/config || true
	cp .env.example dist/.env
	cp -r lang dist/config/lang
	cp casbin_rbac_model.conf dist/config
	rm -rf dist/vendor

vendor:
	go mod vendor

run: build
	./server

dev:
	go run main.go

watch:
	reflex -s -r '\.go$$' make dev
