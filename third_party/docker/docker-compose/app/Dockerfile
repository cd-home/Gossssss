FROM golang:alpine AS Builder
WORKDIR /app/build/
ENV GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GOPROXY=https://goproxy.cn,direct
COPY . .
RUN go build -o app -ldflags="-s -w" main.go

FROM scratch
WORKDIR /app/release/
COPY --from=0 /app/build/app /app/release/
EXPOSE 8999
ENTRYPOINT ["./app"]