.PHONY: all help build test coverage clean run

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=bin/tasktracker
BINARY_UNIX=bin/$(BINARY_NAME)_unix

# target: all - default target. Does nothing.
all:
	@echo "Hello $(LOGNAME), nothing to do by default"
	@echo "Try 'make help'"

# target: help - display callable targets.
help:
	@egrep "^# target:" [Mm]akefile

# target: build - build the binary for the program
build:
	@$(GOBUILD) -o $(BINARY_NAME) -v

# target: test - test the app
test:
	@$(GOTEST) -coverprofile cp.out

# target: coverage - print out the test coverage
coverage:
	@go tool cover -html=cp.out

# target: clean - remove the binary files
clean:
	@$(GOCLEAN)
	@rm -f $(BINARY_NAME)
	@rm -f $(BINARY_UNIX)

# target: run - run the program locally
run:
	@$(GOBUILD) -o $(BINARY_NAME) -v
	@./$(BINARY_NAME)