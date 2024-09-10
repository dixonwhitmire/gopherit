build:
	rm -rf target/
	mkdir target
	go build -o target/gogetem cmd/main.go

test:
	go test ./... -coverprofile=coverage.out

view-coverage:
	go tool cover -html=coverage.out
