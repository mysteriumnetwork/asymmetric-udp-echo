FROM golang AS build

ARG GIT_DESC=undefined

WORKDIR /go/src/github.com/mysteriumnetwork/asymmetric-udp-echo
COPY . .
RUN CGO_ENABLED=0 go build -a -tags netgo -ldflags '-s -w -extldflags "-static" -X main.version='"$GIT_DESC"

FROM scratch
COPY --from=build /go/src/github.com/mysteriumnetwork/asymmetric-udp-echo/asymmetric-udp-echo /
USER 9999:9999
EXPOSE 4589/udp
EXPOSE 4590/udp
ENTRYPOINT ["/asymmetric-udp-echo"]
