FROM php:7.4-fpm-alpine

RUN apk update \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/*

# Add trusted certificates
#COPY docker/certificates /usr/local/share/ca-certificates
# ... to openssl
#RUN update-ca-certificates

RUN mv "$PHP_INI_DIR/php.ini-production" "$PHP_INI_DIR/php.ini"

COPY ./src /web
#COPY ./vendor /app/vendor
RUN chown -R www-data:www-data /web -R
