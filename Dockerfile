FROM registry.access.redhat.com/ubi9/go-toolset:1.18.9-14 as build
USER 0
WORKDIR /build

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -ldflags "-s -w" -o queue-debugger 

FROM registry.access.redhat.com/ubi9/ubi-minimal:9.1
COPY --from=build /build/queue-debugger /queue-debugger
ENTRYPOINT ["/queue-debugger"]
