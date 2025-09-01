# morhaat-service

# To Build the application 
go build -o mh-ui-service ./src/main.go

# To Run the Application
go run ./src/main.go

# Docker 

# Docker Build
docker build -t tushark1704/mh-ui-service:latest .

# Docker Run
docker run -d -p 8080:8080 tushark1704/mh-ui-service:latest

# List All Running Docker processes
docker ps
docker ps -a

# Docker Stop
docker stop <<CONTAINER_ID>>