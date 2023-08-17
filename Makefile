# development setup
dev-run:
	@echo "Starting development server..."
	go run main.go
	
# git config
git-add:
	@echo "Adding all files to git..."
	git add .

git-commit:
	@echo "Commiting changes..."
	read -p "Commit message: " msg; git commit -m "$$msg"

git-push:
	@echo "Pushing changes to remote..."
	read -p "Branch name: " branch; git push origin $$branch

git-pull:
	@echo "Pulling changes from remote..."
	read -p "Branch name: " branch; git pull origin $$branch


# docker config
docker-build:
	@echo "Building docker image..."
	docker build -t go-docker .

docker-run:
	@echo "Running docker image..."
	docker run -p 8080:8080 go-docker

docker-stop:
	@echo "Stopping docker container..."
	docker stop $(docker ps -a -q)

# docker compose config
docker-compose-build:
	@echo "Building docker compose image..."
	docker-compose build

docker-compose-up:
	@echo "Running docker compose image..."
	docker-compose up

docker-compose-down:
	@echo "Stopping docker compose container..."
	docker-compose down

# help config
help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  dev-run           to run the development server"
	@echo "  help              to show this help message"