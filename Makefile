.PHONY: install
install:
	npm install
	go mod tidy
	go mod vendor

.PHONY: build-go
build-go:
	go build ./...

.PHONY: run-go
run-go:
	go run ./...

.PHONY: build-js
build-js:
	npx esbuild --bundle app/main.jsx app/main.css --outdir=static

.PHONY: watch-js
watch-js:
	npx esbuild --watch --bundle app/main.jsx app/main.css --outdir=static
