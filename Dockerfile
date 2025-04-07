# Étape 1 : Utiliser une image de base Go officielle
FROM golang:1.24.1-bookworm AS builder

# Définir le répertoire de travail dans le conteneur
WORKDIR /app

# Copier les fichiers de l'application dans le conteneur
COPY . .

# Télécharger les dépendances (si vous utilisez des modules Go)
RUN go mod download

# Compiler l'application Go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o forum .

# Étape 2 : Créer une image finale légère
FROM alpine:latest

# Définir le répertoire de travail
WORKDIR /app

# Copier l'exécutable compilé depuis l'étape "builder"
COPY --from=builder /app/forum .

# Commande à exécuter au lancement du conteneur
CMD ["./forum"]