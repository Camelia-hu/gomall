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
	@cp proto/cart.proto cart && cd cart && kitex -module github.com/Camelia-hu/gomall -service cart cart.proto && rm -rf cart.proto

.PHONY: docker

docker:
	@sudo docker compose up -d