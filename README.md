# Go Nakama Apps

**Go Nakama Apps** is a backend application built with Go using the Chi framework. It supports a community-based social platform for sharing moments, interacting, and building connections online.

Tagline: **"Nakama: Connecting Communities, Sharing Moments."**

## Table of Contents
- [Go Nakama Apps](#go-nakama-apps)
  - [Table of Contents](#table-of-contents)
  - [Key Features](#key-features)
  - [Installation](#installation)
    - [Step 1: Requirements](#step-1-requirements)
    - [Step 2: Clone the Repository](#step-2-clone-the-repository)
    - [Step 3: Install dependencies](#step-3-install-dependencies)
    - [Step 4: Set up your environment variables file](#step-4-set-up-your-environment-variables-file)
    - [Step 5: Run Postgres Container](#step-5-run-postgres-container)
    - [Step 6: Run Database Migration](#step-6-run-database-migration)
    - [Step 7: Seeding Our Database with Dummy Data](#step-7-seeding-our-database-with-dummy-data)
    - [Step 8: Build and Push Container Image (Optional)](#step-8-build-and-push-container-image-optional)
    - [Step 9: Run Our Application](#step-9-run-our-application)

## Key Features

1. **User Management**:
   - User registration and login.
   - User activation using invitation token.

2. **Post Interaction**:
   - Create, read, update, delete (CRUD) posts.
   - Add comments to posts.
   - Supports pagination and tagging.

3. **Follower System**:
   - Follow and manage followers.

4. **Security**:
   - Basic authentication for system health check.
   - JWT authentication to secure access private endpoint.

5. **Swagger Integration**:
   - Auto-generated API documentation available at `/docs`.

## Installation

### Step 1: Requirements

1. **Go**: To develop and run our application.
   ```shell
   go version
   # go version go1.22.4 linux/amd64
   ```

2. **PostgreSQL**: The primary database for the application.
   ```shell
   psql \
       -h <your_host> \
       -p <your_port> \
       -d <your_database> \
       -U <your_user> \
       -c "SELECT version();"
   # Password for user <your_user>:
   #                                                        version
   # ---------------------------------------------------------------------------------------------------------------------
   #  PostgreSQL 17.0 (Debian 17.0-1.pgdg120+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 12.2.0-14) 12.2.0, 64-bit
   # (1 row)
   ```

3. **Air**: For hot reloading (Optional).
   ```shell
   air -v
   #   __    _   ___
   #  / /\  | | | |_)
   # /_/--\ |_| |_| \_ 1.61.1, built with Go 1.23.2
   ```

4. **Direnv**: For universal environment variable
   ```shell
   direnv --version
   # 2.35.0
   ```

5. **Migrate**: To migrate our database.

6. **Make**: To simplify our workflow commands.
   ```shell
   make --version
   # GNU Make 4.3
   # Built for x86_64-pc-linux-gnu
   # Copyright (C) 1988-2020 Free Software Foundation, Inc.
   # License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
   # This is free software: you are free to change and redistribute it.
   # There is NO WARRANTY, to the extent permitted by law.
   ```

7. **Swagger**: For API Documentation.
   ```shell
   swag --version
   # swag version v1.16.4
   ```

8. **SQL Client (TablePlus)**: To easily interact with PostgreSQL (Optional).
   ![Alt text](images/03_table-plus.png)

9.  **Docker and Docker Compose**: For containerization.
    ```shell
    docker version
    # Client: Docker Engine - Community
    #  Version:           27.0.2
    #  API version:       1.46
    #  Go version:        go1.21.11
    #  Git commit:        912c1dd
    #  Built:             Wed Jun 26 18:47:25 2024
    #  OS/Arch:           linux/amd64
    #  Context:           default

    # Server: Docker Engine - Community
    #  Engine:
    #   Version:          27.0.2
    #   API version:      1.46 (minimum version 1.24)
    #   Go version:       go1.21.11
    #   Git commit:       e953d76
    #   Built:            Wed Jun 26 18:47:25 2024
    #   OS/Arch:          linux/amd64
    #   Experimental:     false
    #  containerd:
    #   Version:          1.7.18
    #   GitCommit:        ae71819c4f5e67bb4d5ae76a6b735f29cc25774e
    #  runc:
    #   Version:          1.7.18
    #   GitCommit:        v1.1.13-0-g58aa920
    #  docker-init:
    #   Version:          0.19.0
    #   GitCommit:        de40ad0

    docker compose version
    # Docker Compose version v2.28.1
    ```

### Step 2: Clone the Repository
```shell
git clone https://github.com/kandlagifari/go-nakama-apps.git
cd go-nakama-apps
```

### Step 3: Install dependencies
```shell
go mod tidy
```

### Step 4: Set up your environment variables file
1. **Application** environment variables in the `.envrc` file
   ```txt
   # Go App
   export ADDR=<your_port> # ":6969"
   export ENV=<your_environment> # "production"

   # Swagger
   export EXTERNAL_URL=<your_swagger_url> # "domainname.com:6969"

   # Go App Connect to Postgre SQL
   export DB_ADDR=<your_postgres_endpoint> # "postgres://postgres:postgres@localhost:5434/postgres_db?sslmode=disable"
   export DB_MAX_OPEN_CONNS=<your_max_open_connections> # 30
   export DB_MAX_IDLE_CONNS=<your_max_idle_connections> # 30
   export DB_MAX_IDLE_TIME=<your_max_idle_time> # "15m"

   # Basic Authentication
   export AUTH_BASIC_USER=<your_health_endpoit_username> # "username"
   export AUTH_BASIC_PASS=<your_health_endpoit_password> # "password"

   # DockerHub
   export DOCKER_USERNAME=<your_dockerhub_username> # username
   export DOCKER_PASSWORD=<your_dockerhub_token> # dckr_pat_r4nd0m_str1n6
   ```

2. **PostgreSQL** environment variables in the `.env` file
   ```txt
   # Postgre Configuration
   POSTGRES_DB=<your_db_name> # postgres_db
   POSTGRES_USER=<your_db_username> # postgres
   POSTGRES_PASSWORD=<your_db_password> # postgres
   ```

### Step 5: Run Postgres Container
```shell
make compose-up

# WARN[0000] /path/to/file/docker-compose.yaml: `version` is obsolete
# [+] Running 3/3
#  ✔ Network go-nakama-apps_nakama-network  Created                                                                                         0.1s
#  ✔ Volume "go-nakama-apps_db-data"        Created                                                                                         0.0s
#  ✔ Container nakama-postgres-container    Started                                                                                         0.4s
```

### Step 6: Run Database Migration
```shell
make migrate-up

# 1/u create_users (28.640288ms)
# 2/u create_posts (35.998278ms)
# 3/u alter_posts_table (46.984194ms)
# 4/u alter_posts_table_with_tags_updated (52.12477ms)
# 5/u create_comments (63.309066ms)
# 6/u alter_posts_table_by_add_version (72.140296ms)
# 7/u create_followers (82.207365ms)
# 8/u add_indexes (101.956591ms)
# 9/u create_user_invitations (107.105056ms)
# 10/u alter_users_table_by_add_isactive (111.94577ms)
# 11/u alter_user_invitations_table_by_add_expiry (117.976111ms)
```

### Step 7: Seeding Our Database with Dummy Data
```shell
go run cmd/migrate/seed/main.go

# yyyy/MM/dd hh:mm:ss Seeding complete
```

### Step 8: Build and Push Container Image (Optional)
You can build and push a container image, or you can also use my public image which has been pushed to Docker Hub
```shell
make build and push

# [+] Building 27.5s (13/13) FINISHED                                                                                             docker:default
#  => [internal] load build definition from Dockerfile                                                                                      0.0s
#  => => transferring dockerfile: 371B                                                                                                      0.0s
#  => WARN: FromAsCasing: 'as' and 'FROM' keywords' casing do not match (line 2)                                                            0.0s
#  => [internal] load metadata for docker.io/library/golang:1.22                                                                            1.7s
#  => [auth] library/golang:pull token for registry-1.docker.io                                                                             0.0s
#  => [internal] load .dockerignore                                                                                                         0.0s
#  => => transferring context: 2B                                                                                                           0.0s
#  => [builder 1/4] FROM docker.io/library/golang:1.22@sha256:147f428a24c6b80b8afbdaec7f245b9e7ac342601e3aeaffb321a103b7c6b3f4              0.0s
#  => [stage-1 1/3] WORKDIR /app                                                                                                            0.0s
#  => [internal] load build context                                                                                                         0.4s
#  => => transferring context: 30.90MB                                                                                                      0.3s
#  => CACHED [builder 2/4] WORKDIR /app                                                                                                     0.0s
#  => [builder 3/4] COPY . .                                                                                                                1.6s
#  => [builder 4/4] RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o api cmd/api/*.go                           23.7s
#  => CACHED [stage-1 2/3] COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/                                           0.0s
#  => CACHED [stage-1 3/3] COPY --from=builder /app/api .                                                                                   0.0s
#  => exporting to image                                                                                                                    0.0s
#  => => exporting layers                                                                                                                   0.0s
#  => => writing image sha256:23a2e61a4c5bb41a41891cd826ec0d50c92d9a746634c6ee2475476f55b7414e                                              0.0s
#  => => naming to docker.io/kandlagifari/nakama-api:latest                                                                                 0.0s

#  1 warning found (use --debug to expand):
#  - FromAsCasing: 'as' and 'FROM' keywords' casing do not match (line 2)
# WARNING! Using --password via the CLI is insecure. Use --password-stdin.
# WARNING! Your password will be stored unencrypted in /home/user/.docker/config.json.
# Configure a credential helper to remove this warning. See
# https://docs.docker.com/engine/reference/commandline/login/#credential-stores

# Login Succeeded
# The push refers to repository [docker.io/kandlagifari/nakama-api]
# 2c4ea4e601a1: Layer already exists
# 90ea6e8eff85: Layer already exists
# 387e3aa195d2: Layer already exists
# latest: digest: sha256:aed44ebfb14b93f4d97bf879e71fa3c4cd0315155155537950fa5a3c1f8275cf size: 945
```

### Step 9: Run Our Application
```shell
docker run -d -p 6969:6969 \
    --name nakama-api-container \
    --network go-nakama-apps_nakama-network \
    -e ADDR=":6969" \
    -e ENV="production" \
    -e EXTERNAL_URL="domainname.com:6969" \
    -e DB_ADDR="postgres://postgres:postgres@172.20.0.1:5434/postgres_db?sslmode=disable" \
    -e DB_MAX_OPEN_CONNS=30 \
    -e DB_MAX_IDLE_CONNS=30 \
    -e DB_MAX_IDLE_TIME="15m" \
    -e AUTH_BASIC_USER="username" \
    -e AUTH_BASIC_PASS="password" \
    kandlagifari/nakama-api:latest

# e0c5e4836159627b6358be01b4ad89d3b6394c86b083072e41b23aa8d6624530
```

The server will be running on http://{{Public-IP}}:6969