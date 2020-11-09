# Build & Copy frontend
# echo "Building frontend..."
# yarn --cwd ../eutenly-website/ build
# rm -rf static/
# cp -r ../eutenly-website/public .
# mv ./public ./static

# Build binary for linux
echo "Building backend..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o eutenly.linux

# Build image & tag latest
echo "Building container..."
docker build -t registry.eutenly.com/eutenly/backend:${version} -t registry.eutenly.com/eutenly/backend:latest .

# Push image
# WARNING: In prod, this will automatically update the live website
echo "Pushing container..."
docker push registry.eutenly.com/eutenly/backend