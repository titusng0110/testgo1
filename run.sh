docker rmi multiplayer-image
docker build -t multiplayer-image multiplayer
docker run -d -p 8080:8080 --name multiplayer-container multiplayer-image