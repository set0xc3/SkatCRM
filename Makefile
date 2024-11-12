build:
	npx tailwindcss -o static/css/tw.css --minify
	templ generate
	go build -ldflags "-s -w" -o bin/server/main.bin cmd/server/main.go

tw:
	@npx tailwindcss -i input.css -o static/css/tw.css --watch

dev:
	@templ generate -watch -proxy="http://localhost:8080" -open-browser=false -cmd="go run cmd/server/main.go"
