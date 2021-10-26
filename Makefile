BINARY_NAME=masv

run:
	@./${BINARY_NAME} 0apple 1lzru 6gump 8hello 5zyo 9apoiuoqiwuero 1abcdefgh

test:
	@cd solutions && go test -v

benchmark:
	@cd solutions && go test --bench=.

build: clean
	@GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go

clean:
	@rm -rf ${BINARY_NAME}