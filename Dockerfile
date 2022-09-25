FROM golang:1.19 AS build
WORKDIR /go/src

COPY . .

RUN go mod download

ARG SERVICE
ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o server ./cmd/$SERVICE


FROM alpine:latest
WORKDIR /app

COPY --from=build /go/src/server .

EXPOSE 4000

CMD ["./server"]
