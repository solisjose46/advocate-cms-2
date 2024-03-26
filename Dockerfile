FROM golang:latest
WORKDIR /app
COPY advocate-2 .
CMD ["go", "run", "main.go"]