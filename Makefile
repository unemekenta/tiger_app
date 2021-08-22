migrate:
	miga --config ./db/miga.local.yml all up

start-frontend:
	cd frontend; yarn serve