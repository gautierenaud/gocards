.PHONY: dev build vendor

dev:
	wails dev -tags webkit2_41

build:
	wails build -tags webkit2_41

vendor:
	go mod vendor
	go mod tidy