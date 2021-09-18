LOCAL_PORT = localhost:7000
DEPLOY_PORT = :7000

AUTH0_ISS = 
AUTH0_AUD = 

TOKEN := test
TOKEN_AUTH0 := 

CURL_GET_LOCAL := curl --header "Authorization: Bearer $(TOKEN)" $(LOCAL_PORT)
CURL_POST_LOCAL := curl --header "Content-Type: application/json" --header "Authorization: Bearer $(TOKEN)" --request POST $(LOCAL_PORT)
CURL_GET_LOCAL_AUTH0 := curl --header "Authorization: Bearer $(TOKEN_AUTH0)" $(LOCAL_PORT)
CURL_POST_LOCAL_AUTH0 := curl --header "Content-Type: application/json" --header "Authorization: Bearer $(TOKEN_AUTH0)" --request POST $(LOCAL_PORT)

run:
	PORT=$(LOCAL_PORT) API_KEY=$(TOKEN) ISS=$(AUTH0_ISS) AUD=$(AUTH0_AUD) go run .

dev: export PORT=$(LOCAL_PORT)
dev: export API_KEY=$(TOKEN)
dev: export ISS=$(AUTH0_ISS)
dev: export AUD=$(AUTH0_AUD)
dev:
	nodemon --exec go run . --signal SIGKILL


g-ping:
	$(CURL_GET_LOCAL)/ping
g-pingA:
	$(CURL_GET_LOCAL_AUTH0)/ping

p-ping:
	$(CURL_POST_LOCAL)/ping --data @_test/ping.json