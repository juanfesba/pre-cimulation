# pre-cimulation
Create Branch.
Publish.

Make local changes. Me :D!

Commit!
(Sync)
On Git, Create Pull Request

More Local changes.
Commit Amend!
Sync... Save... Blah blah!

Merge and delete branch.
Don't forget to go to main branch again.

VS Code Extensions: Docker, Go, Kubernetes, vscode-proto4, YAML, Copilot

===

Local changes.
Commit.

# Creating Go Module
*  go mod init github.com/juanfesba/pre-cimulation
*  go run hello_world/main.go

===

docker build -t hello-image .
// docker build --no-cache --progress=plain -t hello-image .
docker run hello-image

docker run -it hello-image

// docker run -p 8080:8080 my-app-image  # Maps port 8080 in the container to port 8080 on your host
// docker run -e "MY_VARIABLE=my_value" my-app-image
// docker run -it my-app-image bash  # Runs bash inside the container
// docker logs <container_id_or_name (You can get the container ID using docker ps)

docker stop <container_id_or_name>

===

Installing Go Gin (https://github.com/gin-gonic/gin) (https://medium.com/geekculture/full-stack-application-with-go-gin-react-and-mongodb-37b63ef71133)

go mod tidy

# docker system prune
docker build --no-cache -t hello-image .
go mod tidy

# docker system prune
docker build --no-cache -t hello-image .
docker run -p 8080:8080 -d hello-image
docker run -it -p 8080:8080 hello-image

docker ps
docker stop <container-id>
docker rm <container-id>

====

[optional] sometimes, just use incognito
[optional] docker system prune -a
[optional] docker rmi -f $(docker images -a -q)
[optional] docker compose build --no-cache
docker compose up
  docker compose up --build
docker compose up -d (detached)
  docker compose up --build -d (detached)
docker compose down
docker compose build
docker compose logs webapp

====

Use new GitBash terminal (not power shell)
curl localhost:8080/throw --include --header "Content-Type: application/json" -d @curl_body_test.json  --request "POST"
