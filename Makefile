gopkgs= ./cmd/url-shortener/ ./design/ ./server/ ./stores/

export GOBIN=$(shell pwd)/bin
export GOPATH=$(shell pwd)/.go
orgurl=github.com/oosidat
orgpath=$(GOPATH)/src/$(orgurl)
projecturl=$(orgurl)/go-url-shortener
projectpath=$(orgpath)/go-url-shortener

.PHONY: gen build
all: clean build
build: $(GOBIN)/url-shortener
$(GOBIN)/url-shortener: $(projectpath)
	go get -d $(projecturl)/cmd/url-shortener
	go install $(projecturl)/cmd/url-shortener

$(projectpath):
	mkdir -p $(orgpath)
	ln -sf $(shell pwd) $(projectpath)

gen: $(projectpath) $(GOBIN)/goagen
	cd $(projectpath) && $(GOBIN)/goagen bootstrap -d $(projecturl)/design

clean:
	rm -rf $(projectpath)/test/client $(projectpath)/test/models $(projectpath)/client $(GOBIN) $(GOPATH) $(GOTOOL)


$(GOBIN)/goagen:
	go get -d github.com/goadesign/goa/goagen
	go get -d github.com/goadesign/goa/goagen/gen_swagger
	go install github.com/goadesign/goa/goagen

$(GOBIN)/golint:
	go get github.com/golang/lint/golint
	go get -d github.com/golang/lint/golint
	go install github.com/golang/lint/golint
