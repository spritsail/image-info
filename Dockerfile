FROM spritsail/alpine:3.7

ARG VCS_REF

LABEL maintainer="Spritsail <image-info@spritsail.io>" \
      org.label-schema.name="image-info" \
      org.label-schema.url="https://github.com/spritsail/image-info" \
      org.label-schema.description="An informative badge utility providing SVG badges, inspired by shields.io" \
      org.label-schema.version=${VCS_REF}

WORKDIR /app

COPY *.js \
     package.json \
     Verdana.ttf \
     /app/

RUN apk --no-cache add nodejs \
 && npm install --production

CMD ["node", "/app/image-info.js"]
