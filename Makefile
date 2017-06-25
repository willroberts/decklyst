testcard:
	curl http://localhost:8000/card/10303

testdeck:
	curl http://localhost:8000/deck/MTo1MDEsMzo1MTcsMjo1MzgsMzo1NDAsMzoxMDMwMiwyOjEwMzAzLDM6MTAzMDUsMjoxMTA5NCwzOjExMDk3LDM6MjAxMzQsMzoyMDEzOSwyOjIwMTQ0LDI6MjAxNDcsMjoyMDIwNywzOjIwMjM3LDM6MjAyNjE=

build:
	@mkdir -p bin
	go build -o bin/server main.go

deploy:
	@mkdir -p bin
	go build -o bin/server main.go
	#scp assets/cards/v1.86.0.json root@decklyst.xyz:/opt/decklyst/assets/cards/
	scp bin/server root@decklyst.xyz:/opt/decklyst/
