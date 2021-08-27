LOCAL_PORT = localhost:7000
DEPLOY_DEV_PORT = :7000

run:
	PORT=$(LOCAL_PORT) go run .

dev: export PORT=$(LOCAL_PORT)
dev:
	nodemon --exec go run . --signal SIGKILL


g-ping:
	curl $(LOCAL_PORT)/ping
