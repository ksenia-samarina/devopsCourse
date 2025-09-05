include .env

bin-dir:
	mkdir -p ${BINDIR}

build: bin-dir
	#go build -o ${BINDIR}/app ${PACKAGE}
	go build -o ${BINDIR}/app ${CURDIR}/${PACKAGE}

install-grpc-deps: bin-dir
	test -f ${BINDIR}/protoc-gen-go || \
		GOBIN=${BINDIR} go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	test -f ${BINDIR}/protoc-gen-go-grpc || \
		GOBIN=${BINDIR} go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	test -f ${BINDIR}/protoc-gen-grpc-gateway || \
    		GOBIN=${BINDIR} go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

vendor-proto:
		mkdir -p vendor-proto
		@if [ ! -d vendor-proto/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor-proto/googleapis &&\
			mkdir -p  vendor-proto/google/ &&\
			mv vendor-proto/googleapis/google/api vendor-proto/google &&\
			rm -rf vendor-proto/googleapis ;\
		fi
		@if [ ! -d vendor-proto/google/protobuf ]; then\
			git clone https://github.com/protocolbuffers/protobuf vendor-proto/protobuf &&\
			mkdir -p  vendor-proto/google/protobuf &&\
			mv vendor-proto/protobuf/src/google/protobuf/*.proto vendor-proto/google/protobuf &&\
			rm -rf vendor-proto/protobuf ;\
		fi

generate-grpc: bin-dir install-grpc-deps vendor-proto
	mkdir -p pkg/loms_v1
	protoc -I api/loms/v1 -I vendor-proto \
		--go_out=pkg/loms_v1 --go_opt=paths=source_relative \
		--plugin=protoc-gen-go=${BINDIR}/protoc-gen-go \
		--go-grpc_out=pkg/loms_v1 --go-grpc_opt=paths=source_relative \
		--plugin=protoc-gen-go-grpc=${BINDIR}/protoc-gen-go-grpc \
		--grpc-gateway_out=pkg/loms_v1 --grpc-gateway_opt=paths=source_relative \
		--plugin=protoc-gen-grpc-gateway=${BINDIR}/protoc-gen-grpc-gateway \
		api/loms/v1/loms_service.proto

install-minimock: bin-dir
	test -f ${BINDIR}/minimock || \
	GOBIN=${BINDIR} go install github.com/gojuno/minimock/v3/cmd/minimock@v3.1.2

go-generate:
	PATH="${PATH}:${BINDIR}" go generate ./...

install-migrator: bin-dir
	test -f ${BINDIR}/goose || \
		GOBIN=${BINDIR} $ go install github.com/pressly/goose/v3/cmd/goose@v3.10.0

migration-up: bin-dir install-migrator
	${BINDIR}/goose -dir ./db/migrations postgres "host=localhost port=${POSTGRES_PORT} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" up

migration-down: bin-dir install-migrator
	${BINDIR}/goose -dir ./db/migrations postgres "host=localhost port=${POSTGRES_PORT} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" down

migration-status: bin-dir install-migrator
	${BINDIR}/goose -dir ./db/migrations postgres "host=localhost port=${POSTGRES_PORT} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" status

migration-version: bin-dir install-migrator
	${BINDIR}/goose -dir ./db/migrations postgres "host=localhost port=${POSTGRES_PORT} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB} sslmode=disable" version

fill-stocks:
	PGPASSWORD=${POSTGRES_PASSWORD} psql -h "localhost" -p ${POSTGRES_PORT} -U ${POSTGRES_USER} -d ${POSTGRES_DB} \
  -c "\copy stocks(sku, warehouse_id, count) FROM '${CURDIR}/db/init_data/stocks.csv' WITH (FORMAT csv, HEADER)"