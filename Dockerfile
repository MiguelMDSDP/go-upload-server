FROM golang:alpine AS build

RUN apk --no-cache add build-base git gcc
COPY . /src/
RUN cd /src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main
ENTRYPOINT [ "/src/main" ]
VOLUME [ "/src/data" ]
EXPOSE 8080

FROM scratch AS production
COPY --from=build /src/main /
ENTRYPOINT [ "/main" ]
VOLUME [ "/data" ]

FROM production
