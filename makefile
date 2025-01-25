CMD = ./cmd/str_calc
TEST = ./tests
NAME = str_calc


build:
	mkdir -p build
	go build -o ./build/$(NAME) $(CMD)

test:
	go test -v $(TEST)

clean:
	rm -rf build