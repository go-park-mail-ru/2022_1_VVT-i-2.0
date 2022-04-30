FROM golang:alpine as build

COPY . /project

WORKDIR /project

RUN apk add make git && make build

#CMD ./app -config=./config/serv.toml

#================================

FROM alpine:latest

COPY --from=build /project/app /bin/
COPY --from=build /project/config/serv.toml /

CMD app -config=../serv.toml
