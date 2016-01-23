FROM kiasaki/alpine-golang

WORKDIR /gopath/src/app
ADD . /gopath/src/app/
RUN go get app

ENV PRODUCT_SERVICE_PORT=3010

CMD []
ENTRYPOINT ["/gopath/bin/app"]