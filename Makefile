testcard:
	curl -s http://localhost:8000/card/10303
	@echo

testdeck:
	curl -s http://localhost:8000/deck/MTo1MDEsMzo1MTcsMjo1MzgsMzo1NDAsMzoxMDMwMiwyOjEwMzAzLDM6MTAzMDUsMjoxMTA5NCwzOjExMDk3LDM6MjAxMzQsMzoyMDEzOSwyOjIwMTQ0LDI6MjAxNDcsMjoyMDIwNywzOjIwMjM3LDM6MjAyNjE=
	curl -s http://localhost:8000/deck/MTo0MDEsMzo0MDUsMzo0MDcsMzo0MTAsMzo0MTIsMzo0MTMsMzo0MTUsMzoxMDAxMiwzOjEwOTU5LDM6MTA5ODEsMzoxOTAzNywzOjIwMTEzLDM6MjAxMTYsMzoyMDE1Nw==
	@echo

build:
	@mkdir -p bin
	go build -o bin/server main.go

deploy:
	bash deploy.sh
