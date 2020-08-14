FROM golang:1.14-alpine AS build
WORKDIR /src/
COPY main.go /src/
COPY go.mod /src/
COPY go.sum /src/
RUN go mod vendor
RUN CGO_ENABLED=0 go build -o /bin/app
FROM tianon/network-toolbox
COPY --from=build /bin/app /bin/app
COPY assets /assets/
COPY views /views/
ENTRYPOINT ["/bin/app"]
