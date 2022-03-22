# Town Hall

Run a demo web application
Takes the name from a http header (istio integration)
Takes the verb from an environment variable (hcv integration)

## Build and Push

```
echo $PAT | docker login ghcr.io --username schuhu --password-stdin
docker build --tag townhall .
docker tag townhall ghcr.io/schuhu/townhall:1.0.21
docker push ghcr.io/schuhu/townhall:1.0.21
```


## Run in docker

```
docker run -p 0.0.0.0:8080:8080/tcp  --env VERB=delighted townhall
```
