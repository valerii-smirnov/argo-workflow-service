FROM golang:1.16.3-buster AS go-builder

#RUN apk update && apk add --no-cache make

WORKDIR /app

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /bin/entrypoint ./cmd/main.go

FROM scratch

COPY --from=go-builder /bin/entrypoint /bin/entrypoint

CMD ["/bin/entrypoint"]