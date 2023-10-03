# animalpocketresources

## Use the api on local machine

You have to change your docker_compose.yml by this following code lines after creating a .env file completing the variables below:

```
version: '3.8'

services:
  web:
    build: .
    working_dir: /usr/src/app
    env_file:
     - .env
    ports:
      - "3000:3000"
    volumes:
      - .:/usr/src/app
    command: air ./cmd/main.go -b 0.0.0.0
    depends_on:
      - db
  db:
    image: postgres:alpine
    environment:
      - POSTGRES_USER=${DB_USER}
      - POSTGRES_PASSWORD=${DB_PASSWORD}
      - POSTGRES_DB=${DB_NAME}
    ports:
      - "5432:5432"
    volumes:
      - postgres-db:/var/lib/postgresql/data

volumes:
  postgres-db:
```
