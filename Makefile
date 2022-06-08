.PHONY: all

PROJECT_DIR = $(PWD)
FRONT_END_DIR = $(PROJECT_DIR)/minnie-test-frontend
BACK_END_DIR = $(PROJECT_DIR)/minnie-test-backend
MIGRATIONS_FOLDER = $(PWD)/minnie-test-import-data-script

APP_NAME = packform-rest-api
BACKEND_BUILD_DIR = $(BACK_END_DIR)/build/

all: copy-dotenv-to-all-packages create-db-and-import-data install-backend-packages install-frontend-packages run-backend-tests run-frontent-tests indicate-run-commands
run-fresh-build: copy-dotenv-to-all-packages clean-previous-builds generate-builds test run-both-builds
generate-builds: build-frontend copy-frontend-build build-backend
test: run-backend-tests run-frontent-tests

run-both-builds:
	cd $(BACK_END_DIR) && go run main.go
	# $(BACKEND_BUILD_DIR)/$(APP_NAME)


copy-dotenv-to-all-packages:
	echo "\n\nCAREFULL!!!\n\nAssuming .env is having all environment variables set, using it for migrations and running servers\n\n"  &&  rm -rf $(MIGRATIONS_FOLDER)/.env  $(BACK_END_DIR)/.env && cp -r .env $(MIGRATIONS_FOLDER)/ && cp -r .env $(BACK_END_DIR)/

create-db-and-import-data:
	cd $(MIGRATIONS_FOLDER) && go get -d ./... && go run .

install-backend-packages:
	cd $(BACK_END_DIR) && go mod tidy && go get -d ./...
	

install-frontend-packages: 
	cd $(FRONT_END_DIR) && rm -rf node_modules && npm install

run-backend:
	cd $(BACK_END_DIR) && rm -rf server && go build -o server main.go && ./server

run-frontend:
	cd $(FRONT_END_DIR) && npm run serve

build-frontend:
	cd $(FRONT_END_DIR) && npm run build

copy-frontend-build:
	cd $(BACK_END_DIR) && cp -r $(FRONT_END_DIR)/dist/* ./frontend/

build-backend:
	# cd $(BACK_END_DIR) && CGO_ENABLED=0 go build -ldflags="-w -s" -o ./build/$(APP_NAME) main.go

indicate-run-commands:
	echo "\n\nNow generate and run fresh builds for both frontend and backend by \nmake run-fresh-build\n"

clean-previous-builds:
	rm -rf $(BACK_END_DIR)/frontend/* && rm -rf $(BACKEND_BUILD_DIR)/* && rm -rf $(FRONT_END_DIR)/dist/*

run-backend-tests:
	echo "\n\nBegninning to run backend unit tests\n\n" && cd $(BACK_END_DIR) && go test ./internals/routes/orders/orders_test.go

run-frontent-tests:
	echo "\n\nBegninning to run frontend tests\n\n" && cd $(FRONT_END_DIR) && npm run test


# clean:
# 	rm -rf ./build

# build: clean test
# 	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BACKEND_BUILD_DIR)/$(APP_NAME) main.go

# security:
# 	gosec -quiet ./...

# test: security
# 	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
# 	go tool cover -func=cover.out


# swag:
# 	swag init

# runSwag: swag build
# 	$(BACKEND_BUILD_DIR)/$(APP_NAME)

# run-backend:
# 	cd $(BACKEND_BUILD_DIR)/$(APP_NAME) && go run main.go