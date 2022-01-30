# M1チップでないmacユーザ，Windowsユーザ用のコマンド
start-dev:
	docker-compose -f docker-compose.dev.yml up --build -d
start-dev-api:
	docker-compose -f docker-compose.dev.yml up -d mysql phpmyadmin api
down-dev:
	docker-compose -f docker-compose.dev.yml down --rmi all


# M1チップmacユーザ用のコマンド
start-m1:
	docker-compose -f docker-compose.M1.yml up --build -d
start-m1-api:
	docker-compose -f docker-compose.M1.yml up -d mysql phpmyadmin api
down-m1:
	docker-compose -f docker-compose.M1.yml down --rmi all


# 本番環境用のコマンド
start-prod:
	sudo chmod -R 777 db && docker-compose -f docker-compose.prod.yml up --build -d --remove-orphans

down-prod:
	docker-compose -f docker-compose.prod.yml down --rmi all