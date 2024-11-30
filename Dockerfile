FROM golang:1.22.4 AS builder 

WORKDIR /app 
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C ./cmd/ -o goexperts_lab_cloud_run

FROM scratch
WORKDIR /app
COPY --from=builder /app/cmd/goexperts_lab_cloud_run ./
ENTRYPOINT ["./goexperts_lab_cloud_run"]
