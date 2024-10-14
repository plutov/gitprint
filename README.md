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

Request initial cert:

```
docker compose --file compose-prod.yaml run --rm certbot certonly --webroot --webroot-path /var/www/certbot/ --dry-run -d gitprint.me

docker compose --file compose-prod.yaml run --rm certbot certonly --webroot --webroot-path /var/www/certbot/ --dry-run -d api.gitprint.me
```
