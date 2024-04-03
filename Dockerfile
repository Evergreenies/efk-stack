FROM golang:1.22.1-alpine3.18

WORKDIR /app

COPY go.mod go.sum ./ 
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /efk-stack

EXPOSE 8080
CMD [ "/efk-stack" ]







