docker-login:
	docker login 94.139.247.63:5000

# sudo nano /etc/docker/daemon.json - добавить хост vm, чтобы подключаться по http
# sudo systemctl restart docker

docker-build:
	docker-compose build

docker-push:
	docker-compose push

docker-pull:
	docker-compose pull

test-images:
	curl -u myuser:mypassword http://localhost:5000/v2/_catalog