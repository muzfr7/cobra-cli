APP_NAME=Cobra CLI
BINARY_NAME=cobra-cli

build:
	@echo "Building ${APP_NAME}..."
	@go build -o bin/${BINARY_NAME} .
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
