FROM golang:1.15 as build

ENV PORT=8080

WORKDIR /build
COPY hello.go /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o hello hello.go

FROM scratch
WORKDIR /app
COPY --from=build /build/hello /app

EXPOSE ${PORT}

CMD [ "/app/hello" ]

