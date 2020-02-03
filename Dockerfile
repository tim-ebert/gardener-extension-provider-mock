############# builder
FROM golang:1.13.4 AS builder

WORKDIR /go/src/github.com/gardener/gardener-extension-provider-mock
COPY . .
RUN make install-requirements

ARG VERIFY=true

RUN make VERIFY=$VERIFY all

############# base
FROM alpine:3.11.3 AS base

############# gardener-extension-provider-mock
FROM base AS gardener-extension-provider-mock

COPY charts /charts
COPY --from=builder /go/bin/gardener-extension-provider-mock /gardener-extension-provider-mock
ENTRYPOINT ["/gardener-extension-provider-mock"]
