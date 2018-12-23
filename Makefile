BUILD=dist
DEPLOYMENT=$(BUILD).zip
MasterBranch = master
DevelopBranch = develop
Plugins = print

all: 
	go build --buildmode=plugin -o plugins/print.so plugins/print.go
	bee run
	#go test tests/default_test.go

run: 
	#make test
	#bee run -downdoc=true -gendoc=true

pack:
	bee pack

test:
	cp -pr conf tests/
	go test tests/default_test.go
	rm -rf tests/conf

merge:
	git checkout $(DevelopBranch)
	git pull
	git checkout $(MasterBranch)
	git merge $(DevelopBranch)
	git push -u origin $(MasterBranch)

remerge:
	git checkout $(MasterBranch)
	git pull
	git checkout $(DevelopBranch)
	git merge $(MasterBranch)
	git push -u origin $(DevelopBranch)
	

.PHONY: clean
clean:
	rm -f plugins/*.so
	rm -f logs/*.log
