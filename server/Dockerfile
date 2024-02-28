# use official Golang image
# FROM golang:1.16.3-alpine3.13

# FROM golang

# # set working directory
# WORKDIR /app

# # Copy the source code
# COPY . . 

# # Download and install the dependencies
# RUN go get -d -v ./...

# # Build the Go app
# RUN go build -o api .

# #EXPOSE the port
# EXPOSE 8080

# # Run the executable
# CMD ["go","run","./api"]

FROM golang

RUN mkdir /app  
COPY . /app   
WORKDIR /app   
RUN go get ./ && go build && go mod download
EXPOSE 8080

CMD ["go", "run", "/main.go"]