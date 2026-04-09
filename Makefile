PNAM = detoxtube
PREFIX = /usr/local

all: build

build:
	go build -o $(PNAM) main.go

install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f $(PNAM) $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/$(PNAM)

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/$(PNAM)

clean:
	rm -f $(PNAM)

.PHONY: all build install uninstall clean
