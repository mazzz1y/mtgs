###############################################################################
# BUILD STAGE

FROM golang:1.12-alpine

RUN set -x \
  && apk --no-cache --update add \
    bash \
    ca-certificates \
    curl \
    git \
    make \
    upx

COPY . /go/src/mtgs/

RUN set -x \
  && cd /go/src/mtgs \
  && make -j 4 static \
  && upx --ultra-brute -qq ./mtgs


###############################################################################
# PACKAGE STAGE

FROM scratch

ENTRYPOINT ["/mtgs"]
ENV MTG_IP=0.0.0.0 \
    MTG_PORT=3128 \
    MTG_STATS_IP=0.0.0.0 \
    MTG_STATS_PORT=3129
EXPOSE 3128 3129

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=0 /go/src/mtgs/mtgs /mtgs
