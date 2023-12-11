FROM ubuntu:latest
LABEL authors="gelya"

ENTRYPOINT ["top", "-b"]

# Используем Go образ в качестве базового образа
FROM golang:latest

# Установка рабочей директории внутри контейнера
WORKDIR /app

# Копируем файлы проекта внутрь контейнера
COPY . .

# Запускаем команду сборки проекта
RUN go build -o main .

# Определяем команду, которую будет выполнять контейнер при запуске
CMD ["./main"]
