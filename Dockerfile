FROM alpine:3.10.1

ARG version=1.0.0

LABEL maintainer="Haruaki Tamada" \
      description="Notify me later!"

RUN    adduser -D nml \
    && apk --no-cache add --update --virtual .builddeps curl tar \
#    && curl -s -L -O https://github.com/tamada/nml/realeases/download/v${version}/nml-${version}_linux_amd64.tar.gz \
    && curl -s -L -o nml-${version}_linux_amd64.tar.gz https://www.dropbox.com/s/3akz13eepvazu4m/nml-1.0.0_linux_amd64.tar.gz?dl=0 \
    && tar xfz nml-${version}_linux_amd64.tar.gz        \
    && mv nml-${version} /opt                           \
    && ln -s /opt/nml-${version} /opt/nml               \
    && ln -s /opt/nml-${version}/nml /usr/local/bin/nml \
    && rm nml-${version}_linux_amd64.tar.gz             \
    && apk del --purge .builddeps

ENV HOME="/home/nml"

WORKDIR /home/nml
USER    nml

ENTRYPOINT [ "/opt/nml/nml" ]
