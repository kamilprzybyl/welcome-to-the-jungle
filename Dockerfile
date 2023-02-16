FROM golang:1.19

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod ./
RUN go mod download

COPY . .

# run a server
CMD ["air"]
# ENTRYPOINT ["go", "run", "main.go"]
