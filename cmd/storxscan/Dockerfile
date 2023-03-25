ARG DOCKER_ARCH
FROM ${DOCKER_ARCH:-amd64}/alpine

ARG TAG
ARG GOARCH

ENV GOARCH ${GOARCH}

EXPOSE 5999
WORKDIR /app

VOLUME /root/.local/share/storx/storxscan

COPY release/${TAG}/storxscan_linux_${GOARCH:-amd64} /app/storxscan
COPY cmd/storxscan/entrypoint /entrypoint

ENTRYPOINT ["/entrypoint"]

ENV STORX_CONFIG_DIR=/root/.local/share/storx/storxscan
