# Getting started

- Create a new folder for your project
- Initialize the module by running
	- `go mod init github.com/my/package`
	- replace the package with your own repository!
- Get Go modules:
	- `go get -u github.com/labstack/echo/v4`
	- `go get -u github.com/pocketbase/pocketbase`
	- `go get -u github.com/joho/godotenv`
	- `go get -u github.com/golang-jwt/jwt`
    - `go get -u github.com/xuri/excelize/v2`
- Install Templ CLI:
    - `go install github.com/a-h/templ/cmd/templ@latest`
- Install TailwindCSS and DaisyUI:
	- `npm i -D tailwindcss @tailwindcss/typography daisyui`
- Initialize TailwindCSS:
	- `npx tailwindcss init`
- Configure `tailwind.config.js`, which the previous command generated, to look like this:
```javascript
module.exports = {
	content: ["views/**/*.templ"],
	theme: {
		extend: {},
	},
	plugins: [require("@tailwindcss/typography"), require("daisyui")],
	daisyui: {
		themes: ["light"]
	}
}
```

- Create `input.css` at the base of your project with the following contents:
```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```
- Create `Makefile` at the base of your project with the following contents:
```make
tw:
	@npx tailwindcss -i input.css -o static/css/tw.css --watch

dev:
	@templ generate -watch -proxy="http://localhost:8080" -open-browser=false -cmd="go run main.go"
```
