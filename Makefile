app:
	go build -o bin/app cmd/app/main.go

app-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/app-linux-amd64 cmd/app/main.go

test:
	go test ./... -v

submission:
	bash scripts/generate-submission-file.sh