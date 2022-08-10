FROM registry.access.redhat.com/ubi8/ubi:8.4 as build

RUN mkdir /build
WORKDIR /build

RUN dnf -y --disableplugin=subscription-manager install go

COPY go.mod .
RUN go mod download

COPY . .
RUN go build

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.4
COPY --from=build /build/queue-debugger /queue-debugger
ENTRYPOINT ["/queue-debugger"]
