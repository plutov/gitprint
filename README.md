# gitprint

## Local Development

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

## Deployment

1. Set environment variables in `.env` file:

```
GITHUB_CLIENT_ID=
GITHUB_CLIENT_SECRET=
```

2. Run:

```
docker-compose up --build -d
```
