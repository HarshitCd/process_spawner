build:
	@go mod tidy
	@go build -o ./target/spawner main.go

run: build
	@./target/spawner

clean:
	rm -rf target
