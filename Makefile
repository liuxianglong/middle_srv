dao:
	gf gen dao
service:
	gf gen service
enums:
	$(eval _DIR  = $(shell pwd))
	gf gen enums -s ./internal/consts/ -p $(_DIR)/internal/boot/boot_enums.go
pb:
	gf gen pb -a app/rpc/api -c app/rpc/internal/controller