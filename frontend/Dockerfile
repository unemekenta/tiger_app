FROM node:14-alpine

WORKDIR /go/src/tiger_app/frontend

RUN apk --update add --no-cache && \
    npm install -g npm && \
    npm install -g @vue/cli

COPY . .

CMD yarn ${MODE}-serve
