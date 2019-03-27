FROM golang:1.12.0-stretch as build

WORKDIR /code
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go build -o /just-log main.go

FROM gcr.io/distroless/static
COPY --from=build /just-log /bin/just-log