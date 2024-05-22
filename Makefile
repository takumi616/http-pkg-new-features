# Build image
build:
	docker build --no-cache --tag go-http .

# Run container
run:
	docker run --publish 8080:8080 go-http

# Send GET request
get:
	curl -i -X GET localhost:8080/messages/5

# Send POST request
post:
	curl -i -X POST -d '{"content":"sample message."}' localhost:8080/messages 