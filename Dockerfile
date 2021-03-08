FROM golang:1.15 as builder

WORKDIR /usr/src/app
COPY . .

ENV GO111MODULE=on
ENV CGO_ENABLED=0
# deps
RUN go build -i -o pdfgen

# deploy
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/assets ./assets
COPY --from=builder /usr/src/app/pdfgen .
EXPOSE 9000
RUN ls -la pdfgen
CMD ["./pdfgen"]