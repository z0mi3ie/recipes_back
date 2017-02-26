db/start:
	mongod --dbpath /Users/kvickers/src/go/src/recipe/data/db

server/run:
	go run main/main.go

send/recipe:
	curl -X POST localhost:8080/recipe -d @test_recipe.json --header "Content-Type: application/json"
