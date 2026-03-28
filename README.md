<div align="center">
  <img src="ui/public/logo.png" alt="gitprint" width="120" />
  <h3><a href="https://gitprint.me">gitprint.me</a></h3>
  <p>Convert any GitHub repository into a printable PDF.</p>
</div>

---

Sign in with GitHub, paste a repo URL, get a clean PDF — useful for reading code offline, archiving, or just having something physical.

## Stack

- **Go** + [Echo](https://echo.labstack.com) — API server
- **Vue.js** + Tailwind CSS — frontend
- **Gotenberg** — PDF rendering
- **Docker Compose** — orchestration


## Run

```
cp .env.example .env
docker-compose up --build -d
```
