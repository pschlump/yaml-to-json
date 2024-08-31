
all:
	go buld

install:
	( cd ~/bin; rm -f yaml-to-json )
	( cd ~/bin; ln -s /go/src/github.com/pschlump/yaml-to-json/yaml-to-json . )

.PHONY: test
test: test_setup test01 test02 test_done

test_setup:
	go build
	mkdir -p test out ref

test_done:
	@echo PASS

test01:
	yaml-to-json --input test/api.yaml --output out/api01.json
	diff out/api01.json ref

test02:
	yaml-to-json --input test/api.yaml --output out/api02.json --no-indent
	diff out/api02.json ref
