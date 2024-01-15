FROM golang:latest
LABEL authors="alexpresso"

WORKDIR /app
COPY /zunivers-webhooks ./

ENTRYPOINT ["./zunivers-webhooks"]