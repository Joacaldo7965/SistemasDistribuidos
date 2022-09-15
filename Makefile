HOST=$(shell hostname)

test:
	echo "Test"
	@echo $(HOST)

central:
	go run Central/main.go

laboratorio:
	echo "TODO"
    