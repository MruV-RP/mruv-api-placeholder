# Final Stage
FROM alpine:latest

WORKDIR /service

COPY ./mruv-api-placeholder /service/mruv-api-placeholder
RUN chmod +x /service/mruv-api-placeholder

EXPOSE 8080

CMD ["/service/mruv-api-placeholder"]