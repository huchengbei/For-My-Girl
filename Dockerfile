FROM node:14-buster as vuebuilder
WORKDIR /home/node/app
RUN npm config set registry https://registry.npm.taobao.org/
RUN git clone https://github.com/huchengbei/For-My-Girl-Frontend.git /home/node/app --depth 1
RUN npm i
RUN yarn build


FROM golang:1.17.3-alpine3.14 as gobuilder

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /go/src/github.com/huchengbei/for-my-girl/
COPY ./ /go/src/github.com/huchengbei/for-my-girl/
COPY --from=vuebuilder /home/node/app/dist/ /go/src/github.com/huchengbei/for-my-girl/html
RUN go install github.com/rakyll/statik \
    && statik -src=html
RUN CGO_ENABLED=0 go install


FROM alpine:3.15 AS dist
LABEL maintainer="huchengbei <huchengbei@gmail.com>"

ARG TZ="Asia/Shanghai"

ENV TZ ${TZ}

WORKDIR /for-my-girl
COPY --from=gobuilder /go/bin/for-my-girl /for-my-girl/for-my-girl
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk upgrade \
    && apk add bash tzdata \
    && ln -s /for-my-girl/for-my-girl /usr/bin/for-my-girl \
    && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && rm -rf /var/cache/apk/*

RUN mkdir ./conf
COPY ./conf /for-my-girl/conf

EXPOSE 8000/tcp

ENTRYPOINT ["for-my-girl"]
