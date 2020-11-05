docker build -t aad-golang-webapp .
docker run --env-file .env -p 3000:3000 -it aad-golang-webapp
