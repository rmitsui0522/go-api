FROM golang:1.17

WORKDIR /app
# COPY ./app /app

# ホットリロードツール
RUN go install github.com/cosmtrek/air@v1.27.3

CMD ["/app/start.sh"]
