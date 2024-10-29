build:
	npx tailwindcss -o static/css/tw.css --minify
	templ generate
	go build -ldflags "-s -w" -o bin/server/main.bin cmd/server/main.go

dev:
	@templ generate -watch -open-browser=false -cmd="go run cmd/server/main.go"
