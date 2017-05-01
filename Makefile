make:


docker:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	docker build -t gosfs .
install:
	go install
runDocker:
	docker run -p 8000:8000 -v "$$PWD":/data --rm -it gosfs
clean:
	go clean
