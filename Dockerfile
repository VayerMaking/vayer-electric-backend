FROM golang:alpine as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o bin/vayer-electric-backend .

FROM scratch
COPY --from=builder /app/bin/vayer-electric-backend ./vayer-electric-backend
# copy migrations
COPY --from=builder /app/migrations ./migrations
ENTRYPOINT ["/vayer-electric-backend"]
