include .env
export

build:
	go build -o dist/gaskn main.go
	cp -r templates dist/
	mkdir dist/storages
	cp .env.example dist/.env
	cp -r lang dist/lang
	cp casbin_rbac_model.conf dist/

run: build
	./server

dev:
	go run main.go

watch:
	reflex -s -r '\.go$$' make dev
