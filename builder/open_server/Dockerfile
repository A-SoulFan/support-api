FROM golang:1.15.13-alpine as builder

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

WORKDIR /work

COPY go.mod go.sum /work/

ENV GOPROXY=https://goproxy.io,direct

RUN go mod download

COPY ./ /work

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOARM=6 go build -ldflags '-w -s' -o open_server ./cmd/open_server/main.go

FROM alpine as prod

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

COPY --from=builder /work/open_server ./

EXPOSE 9527

CMD ["./open_server"]
