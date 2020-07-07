default:
		@echo "Choose one task"
		@echo "make test: run the tests"
		@echo "make run-app: Run the application and database"
		@exit 1

test:
		@echo "Running the tests"
		@cd src/tests && ENV=test go test

run-app:
		@echo "Running the application"
		@docker-compose up

