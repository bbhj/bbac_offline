BUILD=dist
DEPLOYMENT=$(BUILD).zip

all: 
	#go get github.com/tools/godep
	#godep get
	# bee generate docs
	make test
	#bee run -downdoc=true -gendoc=true
	#bee run

pack:
	bee pack

test:
	go test tests/default_test.go

.PHONY: clean
clean:
	rm -f $(BUILD) $(DEPLOYMENT)
	rm -f lambda.log
