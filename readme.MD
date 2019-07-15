Hi Gophers!

# Chat app - backeng golang, front-end reackt js 

## for golang created with  go modules 

1. Create go modules

```bash
mkdir backed & cd backend
go mog init
```

2. Create a scaffold project in go (with mux) and build it


3. Create a simple Dockerfile

```
FROM golang:1.11.2-stretch
ARG GO111MODULE=on
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
ENTRYPOINT ["/app/<image_name>"]
```

4. Run docker build - `dot` at the end is important

```bash
$ docker build -t <image_name> .
```

5. Check if the image was created

``` bash
$ docker images
to remove image
$ docker rmi <image_name>  (if force necessary use '-f' flag)

```

6. Run the app in a docker (first in interactive mode) first port is port to be expose in docker second default in app

```bash
docker run -p 8080:8080 -it <image_name>
```

7. Run in detatched mode (background)

``` bush
docker run -p 8080:8080 -d <image_name>
```

8. Check docker - process should be in the background

```bash
docker ps
```

9. Other useful commands
a. Get logs of the container
docker logs <containerId>

b.Stop running container
docker stop <containerId>

c.Restart exited container
docker restart <containerId>

d.Remove stopped container
docker rm <containerId>



## Version

1.0.0

## Contributing

1. Fork it
2. Creates your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request