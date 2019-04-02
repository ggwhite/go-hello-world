FROM alpine:3.7

COPY hello /hello

EXPOSE 3000

CMD ["/hello"]