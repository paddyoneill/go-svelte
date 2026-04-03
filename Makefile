.PHONY: oapi-gen
oapi-gen: openapi/server.go openapi/types.go ui/src/lib/api.d.ts

openapi/server.go: openapi/spec.yaml
	go tool oapi-codegen -generate std-http -package api -o openapi/server.go openapi/spec.yaml

openapi/types.go: openapi/spec.yaml
	go tool oapi-codegen -generate types -package api -o openapi/types.go openapi/spec.yaml

ui/src/lib/api.d.ts: openapi/spec.yaml
	npm --prefix ui run oapi:gen:client
