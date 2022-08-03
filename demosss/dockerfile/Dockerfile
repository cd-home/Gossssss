ARG v
FROM golang:${v} AS Builder
LABEL stage=Build
WORKDIR /app/build/
COPY . .
ENV GOOS=linux CGO_ENABLED=0 GOARCH=amd64 GOPROXY=https://goproxy.cn,direct
RUN go build -o app -ldflags="-s -w" main.go

FROM scratch
LABEL stage=Run
WORKDIR /app/release/
COPY --from=Builder /app/build/app /app/release/
EXPOSE 8999
ENTRYPOINT ["./app"]
