build:
	go build -o bin/viewer .

run: build
	./bin/viewer

test:
	go test -v ./... -count=1

tidy:
	go mod tidy

watch:
	clear && tailwindcss -i ./public/css/input.css -o ./public/css/style.css --watch

minify:
	tailwindcss -i ./public/css/input.css -o ./public/css/style.min.css --minify