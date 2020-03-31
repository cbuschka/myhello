build:
	docker build -t cbuschka/myhello:latest .

run:	build
	docker run -p 8080:8080 cbuschka/myhello:latest

push:	build
	docker build -t cbuschka/myhello:${VERSION} .
	docker push cbuschka/myhello:${VERSION}
