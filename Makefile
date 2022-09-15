HOST=$(shell hostname)

test:
	echo "Test"
	@echo $(HOST)

central:
	go run Central/main.go

laboratorio:
	ifeq ($(HOST), dist145)
		go run Lab1/main.go
	ifeq ($(HOST), dist146)
		go run Lab2/main.go
    ifeq ($(HOST), dist147)
		go run Lab3/main.go
	ifeq ($(HOST), dist148)
		go run Lab4/main.go