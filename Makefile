ROOT=$(shell mkdir -p bin)

all: bin/app

bin/app:clean
	@go build -o bin/app cmd/main.go

clean:
	@rm -rf bin/app