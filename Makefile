APPNAME = fortunate
GUIAPPNAME = Fortunate

# Default PREFIX to /usr
ifeq ($(PREFIX),)
    PREFIX := /usr
endif

# Default GOPATH to ~/go
ifeq ($(GOPATH),)
    GOPATH := ~/go
endif

all:
	go build .

clean:
	go clean -i

install:
	install -Dm 755 $(APPNAME) $(DESTDIR)$(PREFIX)/bin/$(APPNAME)
	install -Dm 644 $(GUIAPPNAME).desktop $(DESTDIR)$(PREFIX)/share/applications/$(GUIAPPNAME).desktop
	install -Dm 644 Icon.png $(DESTDIR)$(PREFIX)/share/pixmaps/$(GUIAPPNAME).png

install-darwin:
	sudo launchctl config user path "$$(brew --prefix)/bin:${PATH}"
	go install fyne.io/fyne/v2/cmd/fyne@latest
	$(GOPATH)/bin/fyne install --release

uninstall:
	rm $(DESTDIR)$(PREFIX)/bin/$(APPNAME)
	rm $(DESTDIR)$(PREFIX)/share/applications/$(GUIAPPNAME).desktop
	rm $(DESTDIR)$(PREFIX)/share/pixmaps/$(GUIAPPNAME).png
