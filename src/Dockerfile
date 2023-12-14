FROM golang:1-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

FROM alpine

COPY --from=build /app/main /main
COPY --from=build /app/static /static

EXPOSE 8000

ENTRYPOINT ./main