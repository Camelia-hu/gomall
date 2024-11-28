
.PHONY: kitex-auth

kitex-auth:
	@cp proto/auth.proto auth && cd auth && kitex -module github.com/Camelia-hu/gomall -service auth auth.proto && rm -rf auth.proto


.PHONY: kitex-product

kitex-product:
	@cp proto/product.proto product && cd product && kitex -module github.com/Camelia-hu/gomall -service product product.proto && rm -rf product.proto

.PHONY: kitex-user

kitex-user:
	@cp proto/user.proto user && cd user && kitex -module github.com/Camelia-hu/gomall -service user user.proto && rm -rf user.proto

.PHONY: kitex-cart

kitex-cart:
	@cp cart.proto cart && cd cart && kitex -module github.com/Camelia-hu/gomall -service cart cart.proto && rm -rf cart.proto

.PHONY: kitex-order

kitex-order:
	@cp order.proto order && cd order && kitex -module github.com/Camelia-hu/gomall -service order order.proto && rm -rf order.proto

.PHONY: kitex-payment

kitex-payment:
	@cp payment.proto payment && cd payment && kitex -module github.com/Camelia-hu/gomall -service payment payment.proto && rm -rf payment.proto

.PHONY: kitex-checkout

kitex-checkout:
	@cp checkout.proto checkout && cd checkout && kitex -module github.com/Camelia-hu/gomall -service checkout checkout.proto && rm -rf checkout.proto

.PHONY: docker

docker:
	@sudo docker compose up -d



