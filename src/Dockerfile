# Build stage
FROM golang:alpine AS build

ENV PATH="/protoc/bin:${PATH}"
RUN apk add --no-cache make build-base
WORKDIR /go/src/github.com/medson/superheroesAPI/src

COPY . .
RUN go build main.go

# Final stage
FROM alpine

WORKDIR /app

COPY --from=build /go/src/github.com/medson/superheroesAPI/src ./

EXPOSE 3000

CMD [ "./main" ]