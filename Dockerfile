FROM node:14.17 as vuebuilder
WORKDIR /home/node/app
RUN npm config set registry https://registry.npm.taobao.org/
COPY ./frontend /home/node/app
RUN npm i
RUN yarn build


FROM golang:latest as gobuilder

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /go/src/github.com/huchengbei/for-my-girl/backend
COPY ./backend /go/src/github.com/huchengbei/for-my-girl/backend
COPY --from=vuebuilder /home/node/app/dist/ /go/src/github.com/huchengbei/for-my-girl/backend/html
RUN go get github.com/rakyll/statik \
    && statik -src=html
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:3.13 AS dist
WORKDIR /for-my-girl
LABEL maintainer="huchengbei <huchengbei@gmail.com>"

ARG TZ="Asia/Shanghai"

ENV TZ ${TZ}

COPY --from=gobuilder /go/src/github.com/huchengbei/for-my-girl/backend/app /for-my-girl/for-my-girl
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk upgrade \
    && apk add bash tzdata \
    && ln -s /for-my-girl/for-my-girl /usr/bin/for-my-girl \
    && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && rm -rf /var/cache/apk/*

RUN mkdir ./conf
COPY ./conf /for-my-girl/conf

ENTRYPOINT ["for-my-girl"]
