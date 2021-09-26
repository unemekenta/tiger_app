migrate:
	miga --config ./db/miga.local.yml all up

frontend-dev:
	cd frontend; yarn dev-build

start-local:
	cd frontend; yarn local-serve