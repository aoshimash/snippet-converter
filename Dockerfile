##
## Build
##
FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o snippet-converter ./main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/snippet-converter .


USER nonroot:nonroot

ENTRYPOINT [ "/snippet-converter" ]
