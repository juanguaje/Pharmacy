# Build Stage
FROM golang:1.19-buster:1.13 AS build-stage

LABEL app="build-api-template-juanguaje"
LABEL REPO="https://github.com/juanguaje/api-template-juanguaje"

ENV PROJPATH=/go/src/github.com/juanguaje/api-template-juanguaje

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/juanguaje/api-template-juanguaje
WORKDIR /go/src/github.com/juanguaje/api-template-juanguaje

RUN make build-alpine

# Final Stage
FROM golang:1.19-buster

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/juanguaje/api-template-juanguaje"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/api-template-juanguaje/bin

WORKDIR /opt/api-template-juanguaje/bin

COPY --from=build-stage /go/src/github.com/juanguaje/api-template-juanguaje/bin/api-template-juanguaje /opt/api-template-juanguaje/bin/
RUN chmod +x /opt/api-template-juanguaje/bin/api-template-juanguaje

# Create appuser
RUN adduser -D -g '' api-template-juanguaje
USER api-template-juanguaje

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/api-template-juanguaje/bin/api-template-juanguaje"]
