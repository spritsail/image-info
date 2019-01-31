FROM golang:alpine as build

WORKDIR /go/src/github.com/spritsail/image-info/
ADD . .

RUN apk --no-cache add git \
 && go get \
 && go install

# ~~~~~~~~~~~~~

FROM spritsail/alpine:3.9

ARG VCS_REF

LABEL maintainer="Spritsail <image-info@spritsail.io>" \
      org.label-schema.name="image-info" \
      org.label-schema.url="https://github.com/spritsail/image-info" \
      org.label-schema.description="An informative badge utility providing SVG badges, inspired by shields.io" \
      org.label-schema.version=${VCS_REF}

COPY --from=build /go/bin/image-info /usr/bin/

RUN apk --no-cache add ca-certificates

CMD ["/usr/bin/image-info"]
