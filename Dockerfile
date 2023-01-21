FROM golang:1.19 as app-builder

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
# RUN RUN go build -v -o /usr/local/bin/app ./...
RUN go build -ldflags="-s -w" -o /snippetbox ./cmd/web

# Build app image
FROM golang:1.19 

WORKDIR /usr/src/app

COPY --from=app-builder /snippetbox /usr/local/bin/snippetbox 
COPY --from=app-builder /usr/src/app/tls /usr/src/app/tls
COPY --from=app-builder /usr/src/app/ui /usr/src/app/ui

CMD ["snippetbox"]
