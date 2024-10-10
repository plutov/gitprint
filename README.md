# gitprint

Run:

```
make run
```

Test:

```
make test
```

Lint:

```
make lint
```

Deployment:

```
ssh root@195.201.121.112
cd gitprint
git fetch && git pull
docker compose --file compose-prod.yaml up -d --build
```
