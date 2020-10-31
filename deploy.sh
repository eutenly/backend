# Build binary for linux
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o eutenly.linux

# Build image & tag latest
docker build -t registry.eutenly.com/eutenly/backend:${version} -t registry.eutenly.com/eutenly/backend:latest .

# Push image
# WARNING: In prod, this will automatically update the live website
docker push registry.eutenly.com/eutenly/backend