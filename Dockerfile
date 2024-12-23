FROM golang:1.23 AS build

WORKDIR /app
COPY . /app
RUN go mod download && go mod verify
RUN go build -v -o release

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /app/release /release
EXPOSE 8080 9091
# USER nonroot:nonroot
ENTRYPOINT [ "./release" ]