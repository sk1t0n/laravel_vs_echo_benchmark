FROM debian AS builder

RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    nodejs \
    npm \
    && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

WORKDIR /var/www

COPY ./laravel /var/www

RUN npm install && npm run build

FROM nginx:alpine

COPY ./docker/production/nginx/default.conf /etc/nginx/conf.d/default.conf

COPY --from=builder /var/www/public /var/www/public

WORKDIR /var/www/public

EXPOSE 8080
CMD ["nginx", "-g", "daemon off;"]
