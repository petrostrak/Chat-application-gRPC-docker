# Chat-application-gRPC-docker
In this tutorial, we take a look at how we can build a small chat app with gRPC and Docker. 

### To build from dockerfile
`
docker build --tag=docker_example .
`

### To run docker container with port binding on port 8080
`
docker run -it -p 8080:8080 docker_example
`

### Now under client directory, run the main.go file with your name as a parameter
`
go run main.go -N Petros
`
