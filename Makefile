docker-local-up:
	docker-compose -f docker-compose.yml -f docker-compose.local.yml up --build

docker-dev-up:
	docker-compose -f docker-compose.yml -f docker-compose.dev.yml up --build

docker-api:
	docker-compose exec backend /bin/ash

docker-frontend:
	docker-compose exec frontend /bin/ash

docker-postgres:
	docker-compose exec postgres /bin/ash

migrate:
	miga --config ./db/miga.local.yml all up

frontend-dev:
	cd frontend; yarn dev-build

frontend-local:
	cd frontend; yarn local-serve

air:
	cd backend; air -c .air.toml

reset-frontend:
	cd frontend; rm -rf node_modules && rm package-lock.json && npm cache clear --force && npm cache clean --force && npm i