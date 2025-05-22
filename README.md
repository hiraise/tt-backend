# TaskTrail Backend API

![MIT License](https://img.shields.io/github/license/hiraise/tt-backend)
[![codecov](https://codecov.io/gh/hiraise/tt-backend/graph/badge.svg?token=WD2LRZ5R1I)](https://codecov.io/gh/hiraise/tt-backend)

Backend of the [TaskTrail](https://github.com/hiraise/task-trail) project, built with Go.

## 🚀 Features

_Coming soon_

## 🧱 Technologies

- [Gin](https://github.com/gin-gonic/gin) — High-performance HTTP web framework
- [pgx](https://github.com/jackc/pgx) — PostgreSQL driver and toolkit
- [Docker](https://www.docker.com/) — Containerized environment
- [golangci-lint](https://github.com/golangci/golangci-lint) — Linting

## 🛠 Installation

_Coming soon_

## 🛠 Environment

| Environment Variable                | Example Value          | Description |
| ------------------------------------ | --------------------- | ----------- |
| **APP SETTINGS**                     |                       |             |
| `APP_DEBUG`                          | `true`                | Enable debug mode |
| `APP_ROOT_PATH`                      | `8080`                | Port on which the app will run. Can be empty; defaults to 8080 |
| `APP_ACC_VERIFICATION_ENABLED`       | `true`                | When enabled, service sends confirmation email after user registration. Requires SMTP settings. Can be empty; defaults to true |
| **DATABASE SETTINGS**                |                       |             |
| `PG_MIGRATION_ENABLED`               | `true`                | When enabled, automatically applies all migrations to DB. Can be empty; defaults to false |
| `PG_MIGRATION_PATH`                  | `"file://migrations"` | Migration folder path. Can be empty; defaults to `"file://../../migrations"` (for local run) |
| `PG_CONNECTION_STRING`               | `postgresql://login:password@address:5432/db_name` | PostgreSQL connection string |
| `PG_MAX_POOL_SIZE`                   | `10`                  | Maximum number of simultaneous PostgreSQL connections |
| **AUTHENTICATION SETTINGS**          |                       |             |
| `AUTH_ACCESS_TOKEN_SECRET`           | `s3cr3tK3y!@#2025$%^&*()_+aBcDeFgHiJkLmNoPqRsTuVwXyZ1234567890` | Secret for creating access tokens |
| `AUTH_ACCESS_TOKEN_LIFETIME_MIN`     | `5`                   | Access token lifetime in minutes |
| `AUTH_REFRESH_TOKEN_SECRET`          | `s3cr3tK3y!@#2025$%^&*()_+aBcDeFgHiJkLmNoPqRsTuVwXyZ1234567890` | Secret for creating refresh tokens (should differ from access token secret) |
| `AUTH_REFRESH_TOKEN_LIFETIME_MIN`    | `1440`                | Refresh token lifetime in minutes |
| `AUTH_TOKEN_ISSUER`                  | `example.com`         | [Issuer claim](https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.1) |
| **SMTP SETTINGS**                    |                       |             |
| `SMTP_HOST`                          | `mail.example.com`    | SMTP server host. Required if `APP_ACC_VERIFICATION_ENABLED` is true |
| `SMTP_PORT`                          | `587`                 | SMTP server port. Required if `APP_ACC_VERIFICATION_ENABLED` is true |
| `SMTP_USER`                          | `noreply@example.com` | SMTP server authentication email. Required if `APP_ACC_VERIFICATION_ENABLED` is true |
| `SMTP_PASSWORD`                      | `password123`         | SMTP server authentication password. Required if `APP_ACC_VERIFICATION_ENABLED` is true |
| `SMTP_SENDER`                        | `TaskTrail <noreply@example.com>` | Sender email and name. Can be empty; defaults to `SMTP_USER` |
