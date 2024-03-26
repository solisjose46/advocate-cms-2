# advocate-cms-2 uses docker go sdk to do this progrmatically
FROM golang:latest
WORKDIR /app
COPY advocate-2 .
EXPOSE 8090
ENTRYPOINT ["go", "run", "main.go"]