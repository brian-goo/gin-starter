LOCAL_PORT = localhost:7000
DEV_PORT = :7000

run:
	PORT=$(LOCAL_PORT) go run .

wat: export PORT=$(LOCAL_PORT)
wat:
	nodemon --exec go run . --signal SIGKILL


g-ping:
	curl $(LOCAL_PORT)/ping
