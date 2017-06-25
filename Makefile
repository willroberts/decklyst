build:
	@mkdir -p bin
	go build -o bin/server main.go

deploy:
	scp assets/cards/v1.86.0.json root@decklyst.xyz:/opt/decklyst/assets/cards/
	scp bin/server root@decklyst.xyz:/opt/decklyst/
