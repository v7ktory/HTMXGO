FROM golang:alpine as builder

ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /build
COPY ["go.mod","go.sum", "./"]

RUN go mod download
COPY . .
RUN go build -o ./project ./cmd/app/main.go

FROM scratch

WORKDIR /app
COPY --from=builder /build/project /app/project
CMD ["/app/project"]