FROM golang:alpine AS Builder
WORKDIR /app
COPY go.sum  ./
COPY go.mod ./

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 go test --tags dev ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o  /app/main .

FROM scratch
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]