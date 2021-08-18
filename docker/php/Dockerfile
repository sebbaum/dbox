FROM php:7.4-fpm-alpine

RUN apk update \
    && apk add ca-certificates zip \
    && rm -rf /var/cache/apk/*

# Add trusted certificates
#COPY docker/certificates /usr/local/share/ca-certificates
# ... to openssl
#RUN update-ca-certificates

COPY docker/php/provision/composer.sh /provision/composer.sh
RUN /provision/composer.sh

RUN mv "$PHP_INI_DIR/php.ini-production" "$PHP_INI_DIR/php.ini"

COPY ./src /web
RUN chown -R www-data:www-data /web -R
