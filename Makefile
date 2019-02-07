
build: 
	if [ -a ./main ]; then  rm ./main; fi;   # remove main if it exists 
	go build -o ./main
	./main
