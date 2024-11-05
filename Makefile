run-site:
	tailwindcss -i tailwind.css -o static/app.css --minify
	go run ./cmd/site/main.go
