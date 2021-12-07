FROM node:14-alpine AS builder

LABEL maintainer="Ardhi <ardhi.rofi@gmail.com>"

WORKDIR /app

COPY package.json .
COPY package-lock.json .

RUN npm install

COPY . .

RUN npm run build

FROM nginx:1.21.0-alpine as production
ENV NODE_ENV production
ENV PORT 3000

COPY --from=builder /app/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx","-g","daemon off;"]