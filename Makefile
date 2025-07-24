build:
	sudo docker build . -t 7753547891/rerible:0.0.1

run:
	sudo docker run -p 8080:8080 \
	-e API_KEY=${API_KEY} \
 	-e PORT=${PORT} \
 	-e RARIBLE_ROOT_URL=${RARIBLE_ROOT_URL}  \
 	-e GIN_MODE=${GIN_MODE} 7753547891/rerible:0.0.1 \

deploy:
	helm upgrade rerible ./deploy --set env.RARIBLE_API_KEY=${API_KEY}

push:
	sudo docker push 7753547891/rerible:0.0.1

expose:
	 kubectl port-forward svc/rerible 8080:8080
	 