# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest Linux alpine image
FROM alpine:latest

# Add Maintainer Info
LABEL maintainer="Eutenly"

# Enable watchtower
LABEL com.centurylinklabs.watchtower.enable=true

# Set the Current Working Directory inside the container
WORKDIR /app

EXPOSE 8081

# Copy binary files
COPY eutenly.linux ./main
COPY ./static ./static

# Command to run the executable
CMD ["./main"]