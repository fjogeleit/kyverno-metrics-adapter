FROM golang:1.16-buster as builder

WORKDIR /app
COPY . .

RUN go get -d -v \
    && go install -v

RUN make build

FROM alpine:latest
LABEL MAINTAINER "Frank Jogeleit <frank.jogeleit@gweb.de>"

WORKDIR /app

RUN apk add --update --no-cache ca-certificates

RUN addgroup -S kyverno-metrics-adapter && adduser -u 1234 -S kyverno-metrics-adapter -G kyverno-metrics-adapter

USER 1234

COPY --from=builder /app/LICENSE.md .
COPY --from=builder /app/build/kyverno-metrics-adapter /app/kyverno-metrics-adapter

EXPOSE 2112

ENTRYPOINT ["/app/kyverno-metrics-adapter", "run"]