# Etapa de construção
FROM golang:1.22.5-alpine As build

# Set the working directory to the application root
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the project dependencies using go mod download
RUN go mod download

# Copy the application source code to the working directory
COPY . . ./

# Build the application binary
RUN go build -o app

# Expose port 8000 for the application

# Etapa final
FROM alpine:3.18

WORKDIR /app

# Copie o executável da etapa de construção
COPY --from=build /app/app /app/app

COPY config_docker.json /app/config.json

# Defina o diretório de trabalho e o comando de inicialização
WORKDIR /app
CMD ["/app/app"]