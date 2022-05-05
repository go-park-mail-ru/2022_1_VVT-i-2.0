FROM golang:alpine as build

COPY . /project

WORKDIR /project

RUN apk add make git && make build

#CMD ./app -config=./config/serv.toml

#================================

FROM alpine:latest

COPY --from=build /project/order /bin/
COPY --from=build /project/config/order_deploy.toml /

CMD order -config=../order_deploy.toml