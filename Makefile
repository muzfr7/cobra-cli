APP_NAME=CLI
BINARY_NAME=cli

build:
	@echo "Building ${APP_NAME}..."
	@go build -o bin/${BINARY_NAME} ./cmd/cli/
	@echo "${APP_NAME} built"

run: build
	@echo "Starting ${APP_NAME}..."
	@./bin/${BINARY_NAME} &
	@echo "${APP_NAME} started"

clean:
	@echo "Cleaning..."
	@go clean
	@rm bin/${BINARY_NAME}
	@echo "Cleaned"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done"

start: run

stop:
	@echo "Stopping ${APP_NAME}..."
	@-pkill -SIGTERM -f "./bin/${BINARY_NAME}"
	@echo "Stopped ${APP_NAME}"

restart: stop start
