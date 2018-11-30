all:
	$(MAKE) -C proto all
	go build --buildmode=plugin -o lib/echo.so pkg/echo/plugin/*.go
