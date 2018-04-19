NAME=main
TAG=$(NAME)
VER=v1.0

all: clean build 

build:
	go build main.go

clean:
	rm -rf $(NAME)
