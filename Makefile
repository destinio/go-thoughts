APP_NAME="t"
GO_MAIN_FILE="cmd/main.go"

default:
	@echo "Please specify a target to make.

build: clean
	@echo building $(APP_NAME) in $(HOME)/bin/$(APP_NAME)
	@echo Creating Database in ${HOME}/thoughts/thoughts.csv

	@go build -o $(HOME)/bin/$(APP_NAME) $(GO_MAIN_FILE)

build-dev: clean
	@echo building $(APP_NAME)

	@go build -o $(APP_NAME) $(GO_MAIN_FILE)

build-local: clean-local
	@echo building $(APP_NAME) in "./$(APP_NAME)"
	@go build -o $(APP_NAME) $(GO_MAIN_FILE)

clean:
	@echo cleaning $(APP_NAME)
	@rm -f $(HOME)/bin/$(APP_NAME)

clean-local:
	@echo cleaning $(APP_NAME)
	@rm -f $(APP_NAME)
