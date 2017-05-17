all: build

INSTALLDATA = /usr/bin/install -c -m 644
INSTALLDIR = /usr/local/bin

build:
	go build .


install:
	$(INSTALLDATA) ./find-replace $(INSTALLDIR)
