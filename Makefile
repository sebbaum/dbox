NGINX_IMAGE = cbox_nginx
PHP_IMAGE = cbox_php
TAG = latest

.PHONY: rmi
rmi:
	docker rmi -f ${NGINX_IMAGE}
	docker rmi -f ${PHP_IMAGE}
