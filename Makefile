migrate:
	miga --config ./db/miga.local.yml all up

frontend-dev:
	cd frontend; yarn dev-build

frontend-local:
	cd frontend; yarn local-serve

serve:
	cd backend_src; go run main.go