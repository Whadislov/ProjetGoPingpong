# Étape 1 : Build du backend
FROM golang:1.22.2 AS builder
WORKDIR /src

# Copier le code source
COPY api api
COPY cmd/web/main.go cmd/web/.
COPY internal/my_db internal/my_db
COPY internal/my_functions internal/my_functions
COPY internal/my_types internal/my_types
COPY credentials.env ./
COPY config_app.json ./
COPY config_api.json ./
COPY go.mod go.sum ./
RUN go mod tidy

# Compiler le backend
RUN CGO_ENABLED=0 GOARCH=amd64 go build -o backend ./cmd/web/main.go

# Étape 2 : Copier les fichiers WASM

FROM alpine:latest
#FROM nginx:latest
WORKDIR /app

# Installer Nginx
#RUN apk add --no-cache nginx

# Installer Curl
RUN apk add --no-cache curl

# Copier le backend compilé
COPY --from=builder /src/backend ./backend

# Copier les fichiers WASM + la config
COPY wasm ./wasm
COPY credentials.env ./
COPY config_app.json ./
COPY config_api.json ./

# Copier la configuration de Nginx
#COPY nginx.conf /etc/nginx/nginx.conf
#COPY nginx.conf /etc/nginx/conf.d/default.conf

#COPY ./index.html /usr/share/nginx/html/
#COPY wasm /usr/share/nginx/html/
#COPY credentials.env /usr/share/nginx/html/
#COPY config_app.json /usr/share/nginx/html/
#COPY config_api.json /usr/share/nginx/html/

# Permissions
RUN chmod +x /app/backend

# Exposer les ports
#EXPOSE 8000 8001 8002
EXPOSE 8080 8081

# Lancer le backend et Nginx simultanément
#CMD ["/bin/sh", "-c", "/app/backend & nginx -g 'daemon off;'"]
CMD ["/app/backend"]