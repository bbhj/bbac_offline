BUILD=dist
DEPLOYMENT=$(BUILD).zip

all: 
	#go get github.com/tools/godep
	#godep get
	# bee generate docs
	bee run -downdoc=true -gendoc=true
	bee run

pack:
	bee pack

test:
	go run test/post.go

.PHONY: clean
clean:
	rm -f $(BUILD) $(DEPLOYMENT)
	rm -f lambda.log
