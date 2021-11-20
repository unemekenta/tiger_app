docker-up:
	docker-compose up --build

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

serve:
	cd backend_src; go run main.go