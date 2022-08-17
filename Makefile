include .env
export

build:
	go build -o dist/server main.go
	cp -r templates dist/
	mkdir dist/storages
	cp .env dist/
	cp -r lang dist/lang
	cp casbin_rbac_model dist/

run: build
	./server

dev:
	go run main.go

watch:
	reflex -s -r '\.go$$' make dev
