.PHONY: install
install:
	npm install
	go mod tidy
	go mod vendor

.PHONE: build-go
build-go:
	go build ./...

.PHONE: run-go
run-go:
	go run ./...

.PHONY: build-js
build-js:
	npx esbuild --bundle app/main.jsx app/main.css --outdir=static

.PHONE: watch-js
watch-js:
	npx esbuild --watch --bundle app/main.jsx app/main.css --outdir=static
