fmt:
	go fmt github.com/fuwn/space...

run: fmt validate
	go run github.com/fuwn/space

build: fmt
	go build

validate: .space/.certificates/space.crt .space/.certificates/space.key

ssl:
	mkdir -p .space/.certificates
	openssl req -new -newkey rsa:4096 -x509 -sha256 -days 365 -nodes \
	  -out .space/.certificates/space.crt \
	  -keyout .space/.certificates/space.key \
	  -subj "/CN=fuwn.me"

docker: fmt
	docker build -t fuwn/space:latest .

# https://stackoverflow.com/a/49022012
dangling:
	sudo docker rmi $(sudo docker images -f "dangling=true" -q) --force
