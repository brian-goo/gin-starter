LOCAL_PORT = localhost:7000
DEPLOY_PORT = :7000

TOKEN := test
CURL_GET_LOCAL := curl --header "Authorization: Bearer $(TOKEN)" $(LOCAL_PORT)
CURL_POST_LOCAL := curl --header "Content-Type: application/json" --header "Authorization: Bearer $(TOKEN)" --request POST $(LOCAL_PORT)

run:
	PORT=$(LOCAL_PORT) go run .

dev: export PORT=$(LOCAL_PORT)
dev:
	nodemon --exec go run . --signal SIGKILL


g-ping:
	$(CURL_GET_LOCAL)/ping

p-ping:
	$(CURL_POST_LOCAL)/ping --data @_test/ping.json