# Build Stage
FROM golang:alpine3.12 AS build-stage

ADD . /build
WORKDIR /build

RUN go build .

# Final Stage
FROM alpine:3.12

WORKDIR /service

COPY --from=build-stage /build/mruv-api-placeholder /service/mruv-api-placeholder
RUN chmod +x /service/mruv-api-placeholder

EXPOSE 8080
EXPOSE 8081

CMD ["/service/mruv-api-placeholder"]