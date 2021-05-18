FROM golang:latest as gobuilder

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /go/src/github.com/huchengbei/for-my-girl/backend
COPY ./backend /go/src/github.com/huchengbei/for-my-girl/backend
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM node:14.17 as vuebuilder
WORKDIR /home/node/app
RUN npm config set registry https://registry.npm.taobao.org/
COPY ./frontend /home/node/app
RUN npm i
RUN yarn build

FROM nginx:stable
WORKDIR /usr/share/nginx/html
COPY --from=gobuilder /go/src/github.com/huchengbei/for-my-girl/backend/app ./app
COPY --from=vuebuilder /home/node/app/dist ./
RUN mkdir ./conf
COPY ./conf/app.ini.sample ./conf/app.ini.sample
COPY ./conf/nginx.conf /etc/nginx/conf.d/default.conf
COPY ./entrypoint.sh /docker-entrypoint.d/40-go-moment.sh

# ENTRYPOINT ["/entrypoint.sh"]
