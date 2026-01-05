FROM alpine:3.23@sha256:865b95f46d98cf867a156fe4a135ad3fe50d2056aa3f25ed31662dff6da4eb62

COPY .go-version .go-version

RUN apk add --no-cache go=$(cat .go-version)-r0 curl

COPY . .

ENTRYPOINT ["/entrypoint.sh"]
