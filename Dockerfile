FROM golang:1.15.7-alpine3.13

COPY . /app

CMD ["go", "run", "."]