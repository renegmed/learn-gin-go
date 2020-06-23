init-project:
	go mod init gin-go-mvansickle

build: 
	if [ -a ./main ]; then  rm ./main; fi;   # remove main if it exists 
	go build -o ./main
	./main
