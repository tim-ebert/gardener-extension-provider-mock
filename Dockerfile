############# builder
FROM golang:1.13.4 AS builder

WORKDIR /go/src/github.com/gardener/gardener-extension-provider-mock
COPY . .
RUN make install-requirements && make VERIFY=true all

############# gardener-extension-provider-mock
FROM builder AS gardener-extension-provider-mock

COPY --from=builder /go/bin/gardener-extension-provider-mock /gardener-extension-provider-mock
ENTRYPOINT ["/gardener-extension-provider-mock"]
