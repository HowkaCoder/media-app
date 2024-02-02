# Используем официальный образ Golang в качестве базового образа
FROM golang:latest

# Установка рабочей директории внутри контейнера
WORKDIR /app

# Копируем все файлы из текущего каталога (.) в рабочую директорию (/app) внутри контейнера
COPY . .

# Копируем файл базы данных внутрь контейнера
COPY ./cmd/database/database.db ./cmd/database/database.db

# Собираем ваше приложение (предполагается, что main.go находится внутри папки cmd)
RUN go build -o main ./cmd

# Определяем порт, который будет использоваться вашим приложением
EXPOSE 8082

# Команда, которая будет запущена при запуске контейнера
CMD ["./main"]

