PROJECT?=github.com/donskova1ex/effective_mobile
API_NAME?=songs
API_VERSION?=0.0.1
API_CONTAINER_NAME?=docker.io/donskova1ex/${API_NAME}


clean_api:
	rm -rf bin/${API_NAME}

api_docker_build:
	docker build --no-cache -t ${API_CONTAINER_NAME}:${API_VERSION} -t ${API_CONTAINER_NAME}:latest -f dockerfile.api .