build:
	sudo docker build . -t rerible

run:
	sudo docker run -p 8080:8080 \
	-e API_KEY=${API_KEY} \
 	-e PORT=${PORT} \
 	-e RARIBLE_ROOT_URL=${RARIBLE_ROOT_URL}  \
 	-e GIN_MODE=${GIN_MODE} rerible \