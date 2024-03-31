FROM debian:bookworm
ARG TARGETARCH

WORKDIR /app

RUN apt-get update && \
    apt-get install -y ca-certificates

COPY openmeteo-exporter-$TARGETARCH /usr/bin/openmeteo-exporter

CMD ["/usr/bin/openmeteo-exporter"]
