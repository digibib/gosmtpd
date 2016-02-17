FROM alpine:3.3
MAINTAINER Soren Mathiasen <sorenm@mymessages.dk>
WORKDIR /app
COPY ./build/app /app/gosmtpd

EXPOSE 2525
EXPOSE 8080
CMD ["/app/gosmtpd"]
