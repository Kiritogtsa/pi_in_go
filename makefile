GOCMD:=go

GOBUILD:=$(GOCMD) build
GOCLEAN:=$(GOCMD) clean


BINARY_NAME:=servidor
APP_DIR:= ../cmd/app
all:

startcompose:
	docker compose start
stopcompose:
	docker compose stop
build:
	$(GOBUILD) -o $(BINARY_NAME) $(APP_DIR)
clean:
	rm -f $(BINARY_NAME)
run:
	$(GOCMD) run $(APP_DIR)/main.go
create:
	docker compose up -d
delete:
	docker compose down -v