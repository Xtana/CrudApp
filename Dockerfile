FROM golang

WORKDIR /crudapp

COPY go.mod go.sum ./
RUN go mod download

# Копируем конфигурационный файл
COPY config/config.yaml ./config.yaml

# Копируем остальные файлы
COPY . .

RUN go build -o /crudapp/crudapp ./cmd/crudapp

CMD ["/crudapp/crudapp"]