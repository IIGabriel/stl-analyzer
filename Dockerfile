FROM golang:1.22-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init --dir ./cmd/api --output ./docs --parseDependency --parseInternal

ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -a -installsuffix cgo -o app cmd/api/main.go

FROM --platform=linux/amd64 alpine

WORKDIR /app

COPY --from=build /app/docs/swagger.json /app/docs/swagger.json
COPY --from=build /app/app /app

EXPOSE 8080

ENTRYPOINT [ "/app/app" ]
