CURRENT_DIR = $(shell pwd)

start_db:
	mongod --dbpath $(CURRENT_DIR)/data/db

start_server:
	go run recipe.go
