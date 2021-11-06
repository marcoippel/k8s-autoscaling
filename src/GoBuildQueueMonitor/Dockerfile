FROM golang:alpine as builder

RUN mkdir app
RUN mkdir -p /src/GoBuildQueueMonitor/

# First create a root directory
WORKDIR /app

# Create the go.mod file
RUN go mod init buildqueuemonitor

# Copy all the files in the same folder structure so the package include files are correct
COPY . /app/src/GoBuildQueueMonitor/

# Switch to the source folder
WORKDIR /app/src/GoBuildQueueMonitor/

# RUN GOOS=linux go build -o GoBuildQueueMonitor .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o GoBuildQueueMonitor .


FROM gcr.io/distroless/static:latest
COPY --from=builder /app/src/GoBuildQueueMonitor/GoBuildQueueMonitor .

CMD [ "/GoBuildQueueMonitor" ]