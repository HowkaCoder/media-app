# media-app

Backend for a multilingual electronics store: product catalog with categories, characteristics and translations, image processing, customer accounts with SMS verification, orders and sales metrics.

Built for a real client and used in production in 2024.

![Go](https://img.shields.io/badge/Go-1.21-00ADD8?logo=go&logoColor=white)
![Fiber](https://img.shields.io/badge/Fiber-v2-00ACD7)
![SQLite](https://img.shields.io/badge/SQLite-GORM-003B57?logo=sqlite&logoColor=white)
![Redis](https://img.shields.io/badge/Redis-OTP-DC382D?logo=redis&logoColor=white)
[![ci](https://github.com/HowkaCoder/media-app/actions/workflows/ci.yml/badge.svg)](https://github.com/HowkaCoder/media-app/actions/workflows/ci.yml)

## Features

- Catalog: main categories → subcategories → products, filters by price, discount, brand and characteristics, sorting and pagination
- Product content in several languages (translations for products and characteristics)
- Image upload as base64, converted to WebP on the server
- Registration with SMS code (Eskiz.uz), JWT access and refresh tokens, `user` and `admin` roles
- Orders with stock check: stock is decreased atomically inside a transaction, so two buyers can't take the last item
- Daily, weekly, monthly and yearly sales metrics calculated in the background

## Architecture

```
main.go                  wiring, routes, graceful shutdown
internal/
  config/                configuration from environment variables
  app/
    entity/              GORM models
    repository/          database access
    usecase/             business rules
    service/             JWT and password hashing
    handler/             HTTP handlers (Fiber)
  otp/                   one-time codes in Redis: TTL, resend cooldown, attempt limit
  sms/                   Eskiz.uz client with token refresh
```

Request flow: `handler → usecase → repository → database`.

## Security

- Passwords are hashed with bcrypt and never returned by the API
- JWT secret and SMS credentials come from the environment
- Write endpoints require an `admin` token; users can read and change only their own profile
- Public registration always creates a `user`, the role can be changed only by an admin
- SMS codes expire after 5 minutes, can be requested once a minute and allow 5 attempts

## Getting started

```bash
cp .env.example .env      # set JWT_SECRET
go run .
```

SMS verification is turned on when `ESKIZ_EMAIL` and `ESKIZ_PASSWORD` are set (Redis is required in that case).

With Docker:

```bash
docker compose up --build
```

Requires `libwebp` for local builds (`apt install libwebp-dev` / `brew install webp`).

## Configuration

| Variable | Default | Description |
|---|---|---|
| `PORT` | `8082` | HTTP port |
| `DB_PATH` | `database/database.db` | SQLite file |
| `IMAGES_DIR` | `images` | Where uploaded images are stored |
| `PUBLIC_URL` | `http://localhost:8082` | Base URL used in image links |
| `JWT_SECRET` | — | Required |
| `REDIS_ADDR` | `localhost:6379` | Redis for SMS codes |
| `ESKIZ_EMAIL`, `ESKIZ_PASSWORD` | — | Eskiz.uz credentials, SMS is disabled if empty |

## API

| Method | Path | Access |
|---|---|---|
| `POST` | `/api/users/phone` | public, sends SMS code |
| `POST` | `/register` | public |
| `POST` | `/login` | public |
| `GET` | `/api/users-profile` | user |
| `GET`, `PATCH` | `/api/users/:id` | owner or admin |
| `GET` | `/api/users` | admin |
| `GET` | `/api/:lang/products` | public (`limit`, `offset`, `min`, `max`, `value`, `description`) |
| `GET` | `/api/:lang/products/:id` | public |
| `GET` | `/api/:lang/categories/:id/products` | public (`sort=cheap\|expensive`) |
| `POST` | `/api/:lang/products-filter` | public |
| `POST`, `PATCH`, `DELETE` | `/api/:lang/products[/:id]` | admin |
| `GET` | `/:lang/api/maincategories[/:id]` | public |
| `GET` | `/api/subcategories[/:id]` | public |
| `GET` | `/api/languages[/:id]` | public |
| `POST` | `/api/orders` | user |
| `GET`, `PATCH`, `DELETE` | `/api/orders[/:id]` | admin |
| `GET` | `/metrics` | admin |

Write endpoints for categories, subcategories, languages and translations are admin only.

## Tests

```bash
go test ./...
```

Covers password hashing and JWT, the OTP flow (expiry, cooldown, attempt limit) and HTTP tests for login, roles and access to user profiles.
