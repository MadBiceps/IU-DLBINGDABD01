# Starte von einem offiziellen Golang-Image
FROM golang:1.18 as builder

# Setze das Arbeitsverzeichnis im Container
WORKDIR /app

# Kopiere den Go-Quellcode
COPY . .

# Lade Abhängigkeiten herunter und kompiliere das Go-Programm
RUN go mod init temp-publisher && go get -d -v && go build -o /temp-publisher

# Startphase
FROM gcr.io/distroless/base-debian10

COPY --from=builder /temp-publisher /

# Starte das Programm
CMD ["/temp-publisher"]