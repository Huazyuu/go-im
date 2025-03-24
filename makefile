.PHONY: build buildwin buildlinux \
        run clean \
        gotool \
        swag \
        air-init air-run \
        dbcreate dbload dbdump \
        escreate esload esdump \
        avatar \
        help

BINARY="./build/yublog"

build :
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe

buildwin :
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BINARY}.exe

buildlinux :
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run : 
	go run main.go


clean :
	@if [ -f ${BINARY} ]; then rm ${BINARY}; fi
	@if [ -f ${BINARY}.exe ]; then rm ${BINARY}.exe; fi


gotool :
	go fmt 
	go vet 

swag :
	swag init

air-init:
	air init

air-run:
	air run 


dbcreate :
	go run main.go -dbcreate

dbload :
    # 假设需要指定导入文件路径，这里以占位符表示，实际使用时替换
	go run main.go -dbload = your_database_file.sql 

dbdump :
	go run main.go -dbdump


escreate:
	go run main.go -escreate

esload :
    # 假设需要指定导入 ES 索引文件路径，这里以占位符表示，实际使用时替换
	go run main.go -esload = your_es_index_file.json 

esdump:
	go run main.go -esdump

avatar:	
	go run main.go -avatar

help :
	@echo "make build - Compile the Go code for Windows and generate the binary"
	@echo "make buildwin - Compile the Go code for Windows and generate the binary"
	@echo "make buildlinux - Compile the Go code for Linux and generate the binary"
	@echo "make run - Run main.go"
	@echo "make dbcreate - Initialize the database"
	@echo "make dbload - Load data into the database from a file (replace 'your_database_file.sql')"
	@echo "make dbdump - Dump the database to a file"
	@echo "make escreate - Create an ES index"
	@echo "make esload - Load data into an ES index from a file (replace 'your_es_index_file.json')"
	@echo "make esdump - Dump data from an ES index"
	@echo "make clean - Remove binary code files"
	@echo "make gotool - Run Go tools: 'fmt' and 'vet'"
	@echo "make swag - Generate API doc"
	@echo "make air-init - Initialize air configuration"
	@echo "make air-run - Run the project with air for auto - restart"