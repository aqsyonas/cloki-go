FROM golang:alpine as builder
ENV BUILD 20200111-004
RUN apk --update add git make
COPY . /cloki-go
WORKDIR /cloki-go
RUN make modules && make all

FROM alpine
WORKDIR /
RUN apk --update add bash sed
# Create default directories
RUN mkdir -p /usr/local/cloki-go
COPY --from=builder /cloki/cloki-go .
COPY --from=builder /cloki/docker/webapp_config.json /usr/local/cloki/webapp_config.json
# Configure entrypoint
COPY ./docker/docker-entrypoint.sh /
COPY ./docker/docker-entrypoint.d/* /docker-entrypoint.d/
RUN chmod +x /docker-entrypoint.d/* /docker-entrypoint.sh

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["/cloki-go", "-config-path", "/usr/local/cloki-go"]
