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

// docker run -p 8080:8080 my-app-image  # Maps port 8080 in the container to port 8080 on your host
// docker run -e "MY_VARIABLE=my_value" my-app-image
// docker run -it my-app-image bash  # Runs bash inside the container
// docker logs <container_id_or_name (You can get the container ID using docker ps)

docker stop <container_id_or_name>
