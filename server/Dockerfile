FROM golang

# WORKDIR /app
WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
COPY /ASSETS /app/ASSETS
RUN go build -v -o main .



CMD ["/app/main"]