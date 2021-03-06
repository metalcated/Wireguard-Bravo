GOCMD=go

GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BINARY_NAME=bravo

all: test build

build: 
	$(GOBUILD) -v -o ${BINARY_NAME}
	echo "sudo ./$(BINARY_NAME) &" >> run.sh
	chmod +x run.sh

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	sed -i '/sudo \.\/*/d' run.sh

run:
	$(GOBUILD) -v -o ${BINARY_NAME}
	echo "sudo ./$(BINARY_NAME) &" >> run.sh
	chmod +x run.sh
	./run.sh
