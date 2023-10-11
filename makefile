build:
	protoc api/*.proto \
		--go_out=. \
		--go-grpc_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.

deploy-otelcol:
	docker run --name otelcol \
		-d -p 4317:4317 \
		-v ./otel-config.yaml:/etc/otelcol-contrib/config.yaml \
		otel/opentelemetry-collector-contrib

deploy-zipkin:
	docker run --name zipkin -d -p 9411:9411 openzipkin/zipkin

remove-otelcol:
	docker rm -f otelcol

remove-zipkin:
	docker rm -f zipkin

logs-otelcol:
	docker logs -f otelcol
