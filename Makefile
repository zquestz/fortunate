APPNAME = fortunate

# Default PREFIX to /usr/local
ifeq ($(PREFIX),)
    PREFIX := /usr/local
endif

all:
	go build .

clean:
	go clean -i

install:
	install -d $(DESTDIR)$(PREFIX)/bin/
	install -m 755 $(APPNAME) $(DESTDIR)$(PREFIX)/bin/
	install -d $(DESTDIR)$(PREFIX)/share/applications
	install -m 644 $(APPNAME).desktop $(DESTDIR)$(PREFIX)/share/applications/
	install -d $(DESTDIR)$(PREFIX)/share/icons
	install -m 644 icon/$(APPNAME).png $(DESTDIR)$(PREFIX)/share/icons/

uninstall:
	rm $(DESTDIR)$(PREFIX)/bin/$(APPNAME)
	rm $(DESTDIR)$(PREFIX)/share/applications/$(APPNAME).desktop
	rm $(DESTDIR)$(PREFIX)/share/icons/$(APPNAME).png
