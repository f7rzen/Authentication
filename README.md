# Authentication System

## Описание

Система аутентификации на основе Golang с использованием фреймворка Gin и базы данных PostgreSQL. Система позволяет пользователям регистрироваться, входить в систему и отображать приветственное сообщение после успешного входа или регистрации.

### Установка

1. Клонируйте репозиторий на ваше локальное устройство:

```bash
git clone https://github.com/f7rzen/Authentication.git
```

2. Перейдите в директорию проекта:

```bash
cd Authentication
```

3. Настройте базу данных PostgreSQL в файле main.go:

```bash
connStr := "user=postgres password=yourpassword dbname=users sslmode=disable"
```

4. Установите зависимости

```bash
go mod tidy
```

### Запуск
1. Запустите программу:

```bash
go run . 
```

### Лицензия
Этот проект лицензирован в соответствии с условиями [MIT License](https://opensource.org/license/MIT).
