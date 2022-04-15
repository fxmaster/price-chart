
protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ${arg}

docker-up:
	docker-compose -f ./resource/docker-compose.yaml up -d

docker-down:
	docker-compose -f ./resource/docker-compose.yaml down --remove-orphans

run:
	find ./cmd -not -path "./.*" -name 'main.go' -maxdepth 2 -mindepth 1 -exec sh -c 'go run {} &' \;


# find ./cmd -not -path "./.*" -name 'main.go' -maxdepth 2 -mindepth 1 -exec echo {} \;
