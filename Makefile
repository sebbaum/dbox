NGINX_IMAGE = cbox_nginx
PHP_IMAGE = cbox_php
TAG = latest

.PHONY: rm
rm:
	docker rm -f ${NGINX_IMAGE}
	docker rm -f ${PHP_IMAGE}

.PHONY: rmi
rmi:
	docker rmi -f ${NGINX_IMAGE}
	docker rmi -f ${PHP_IMAGE}
