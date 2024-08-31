
all:
	go buld

install:
	( cd ~/bin; rm -f yaml-to-json )
	( cd ~/bin; ln -s /go/src/github.com/pschlump/yaml-to-json/yaml-to-json . )

