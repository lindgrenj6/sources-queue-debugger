FROM docker.io/golang:1.18 as build

RUN mkdir /build
WORKDIR /build

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o queue-debugger && strip queue-debugger

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.6
COPY --from=build /build/queue-debugger /queue-debugger
ENTRYPOINT ["/queue-debugger"]
