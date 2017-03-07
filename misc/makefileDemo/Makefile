BIN_DIR=./bin/
OWN_BIN=$(BIN_DIR)/myCmd

all:$(OWN_BIN)

SRC_DIR=./src/
SRCS=$(wildcard $(SRC_DIR)*.go)

$(OWN_BIN):$(SRCS)
#	go build -o $(OWN_BIN) $(SRC_DIR)
	go build -o $(OWN_BIN)  -gcflags "-N -l" -ldflags "-s" $(SRC_DIR)



clean:
	rm -f $(OWN_BIN)

run:
	make
	echo ________________
	$(OWN_BIN)