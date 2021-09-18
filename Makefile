LOCAL_PORT = localhost:7000
DEPLOY_PORT = :7000

AUTH0_ISS = https://dev-h54azlcu.jp.auth0.com/
AUTH0_AUD = https://trade.api.dev

TOKEN := test
TOKEN_AUTH0 := eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6ImJKVHhTSWlpTTc4SjllczB3Qm1HdyJ9.eyJpc3MiOiJodHRwczovL2Rldi1oNTRhemxjdS5qcC5hdXRoMC5jb20vIiwic3ViIjoiZ29vZ2xlLW9hdXRoMnwxMDc4NDk3MDA4ODM0NjM0NzQ4NzgiLCJhdWQiOlsiaHR0cHM6Ly90cmFkZS5hcGkuZGV2IiwiaHR0cHM6Ly9kZXYtaDU0YXpsY3UuanAuYXV0aDAuY29tL3VzZXJpbmZvIl0sImlhdCI6MTYzMTkzMTgxNCwiZXhwIjoxNjMyMDE4MjE0LCJhenAiOiI0MWVqTVMydERhVk81czVRVHEzdGIwQXhWdFRLcExnWSIsInNjb3BlIjoib3BlbmlkIHByb2ZpbGUgZW1haWwifQ.JV_0AIfcUfq9gwDF5-zartDdgKRjehvslC0XZRIgUlMOvCgZ1zybZMxWyxfwpzHU2ERV5YXtaRy4ibgjH1sA08RXdkXh1GYeaL5xctRpqpyikcPOndrDhBZQArFIUAyV9uJeIvU-inOiHfSZk7bsUEna9DzybA5HLJnyyIkqcn8qaF3CWqwTjtUXrxEnhXTtCLOojDvQqpc1pLmEpt3Zh5PGL1w8WTLbfnJ8bMQRPv05DcfzUhYohWv0HimycHN8aWYExLeHcLpye9xoaUSrL5t9YbIMEVdANDvUB8SeJM2B5D64e5SDJMspTOAVMv3R5VSxp8zFrJn749-v7KhNNA

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