FROM golang:1.17-alpine
RUN apk --no-cache add ca-certificates git

ENV APP_NAME go-joke
ENV PORT 8082

COPY . "/go/src/${APP_NAME}"
WORKDIR "/go/src/${APP_NAME}"
RUN go mod download
RUN go build -o "cmd/${APP_NAME}" cmd/main.go

CMD "./cmd/${APP_NAME}"

EXPOSE ${PORT}
