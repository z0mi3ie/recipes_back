# recipes_back

# TODO
# Data Validation/Sanitization
- [x] make sure incoming name has no spaces

# MongoDB 
- [ ] set db index info and config https://gist.github.com/border/3489566
- [ ] pass ENV struct with db connection around (move away from creating brand new session, copy session per user)
- [ ] setup command line arg scripts to manage db (separate program)

# Authentication (expand this)
- authN
- authZ
- OAuth? 

# Basic Functionality
- [x] Fix conflict for single recipe and all recipe search 
- [x] Delete recipe
- [ ] Update recipe
- [ ] Pagination for all recipes (low priority)
- [ ] Search for recipes by category
	- Categories
- [ ] Make sure NotFound and BadRequest are handled (404 and 400)

# Logging
- [ ] Add true logging

# Devops
- Vagrant / Ansible full deploy
	- db server
	- recipe front server
	- recipe back server
- Docker
	- Unit test
	- Functional test (python func test)

# Testing
- unit
- functional

# Front End
- [ ] view single recipe page
- [ ] browes recipes page
- [ ] ajax call on add recipe to hit this service
- [ ] connect them basic for now

# Resources
https://codeplanet.io/principles-good-restful-api-design/
