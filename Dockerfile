# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN apk add -q git
RUN go get gopkg.in/src-d/go-git.v4 
RUN cd /src && go build -o app

# final stage
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build-env /src/app /app/
ENTRYPOINT ./app