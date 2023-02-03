# syntax=docker/dockerfile:1

FROM golang:1.18-alpine AS build

RUN apk update && apk add --no-cache build-base git
ADD . /app
WORKDIR /app

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /deva-fiber

EXPOSE 8000

FROM alpine

COPY conf /app/conf
WORKDIR /app
COPY --from=build /deva-fiber /app/deva-fiber 
RUN apk update && apk add nano 

CMD [ "/app/deva-fiber" ]
