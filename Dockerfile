FROM golang:alpine as build

LABEL maintainer="Muhammad Luthfi <muhammadluthfi059@gmail.com>"

ARG APP_NAME=api-sftp-client

RUN mkdir /app
ADD . /app/

COPY .env /app

WORKDIR /app
RUN go build -mod=vendor -o ${APP_NAME} .

FROM alpine
WORKDIR /app
COPY --from=build /app/${APP_NAME}  /app/${APP_NAME}
COPY --from=build /app/.env         /app/.env
EXPOSE 8081

ENTRYPOINT ["/app/api-sftp-client"]
