go build -o buildqueuemonitor
docker rm -f buildqueuemonitor
docker build -t buildqueuemonitor ../.
docker run -d --name buildqueuemonitor -p 85:8080  buildqueuemonitor