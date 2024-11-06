APPNAME = fortunate

# Default PREFIX to /usr
ifeq ($(PREFIX),)
    PREFIX := /usr
endif

all:
	go build .

clean:
	go clean -i

install:
	install -Dm 755 $(APPNAME) $(DESTDIR)$(PREFIX)/bin/$(APPNAME)
	install -Dm 644 $(APPNAME).desktop $(DESTDIR)$(PREFIX)/share/applications/$(APPNAME).desktop
	install -Dm 644 icon/$(APPNAME).png $(DESTDIR)$(PREFIX)/share/icons/$(APPNAME).png

uninstall:
	rm $(DESTDIR)$(PREFIX)/bin/$(APPNAME)
	rm $(DESTDIR)$(PREFIX)/share/applications/$(APPNAME).desktop
	rm $(DESTDIR)$(PREFIX)/share/icons/$(APPNAME).png
