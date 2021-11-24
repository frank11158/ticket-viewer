
install:
	go build -v .
ifneq ($(findstring $(ENV),$(ENV_LIST)),)
	cp ./env/$(ENV).env .env
endif