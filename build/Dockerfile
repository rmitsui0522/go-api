FROM golang:1.17

WORKDIR /app
COPY . /app

# ホットリロードツール
RUN go install github.com/cosmtrek/air@v1.27.3

CMD ["air"]