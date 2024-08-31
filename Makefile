
all:
	go buld

install:
	( cd ~/bin; rm -f yaml-to-json )
	( cd ~/bin; ln -s /go/src/github.com/pschlump/yaml-to-json/yaml-to-json . )

.PHONY: test
test:
	go build
	mkdir -p test out ref
	yaml-to-json --input test/api.yaml --output out/api.json
	diff out/api.json ref
