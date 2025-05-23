# Social Network Platform

Заготовка для социальной сети с базовым функционалом.

## 📌 Основные компоненты

### ✅ Сервис постов (Post Service)

- Создание/редактирование/удаление постов
- Лента пользователя (подписки)

### 📊 Сервис аналитики (Analytics Service)

- Статистика активности пользователей
- Топ контента
- Аналитика роста платформы
- Кастомные отчеты (для админов)

## 🛠 Технологический стек

### Backend

| Компонент                   | Технологии          |
| --------------------------- | ------------------- |
| Язык                        | Go                  |
| Межсервисное взаимодействие | gRPC + REST Gateway |
| Асинхронные задачи          | Kafka               |

### Базы данных

| Назначение         | Технология                   |
| ------------------ | ---------------------------- |
| Основное хранилище | PostgreSQL (с шардированием) |
| Кэширование        | Redis (счетчики, сессии)     |

### Инфраструктура

| Компонент     | Технологии                        |
| ------------- | --------------------------------- |
| Мониторинг    | Prometheus + Grafana              |
| Логирование   | Zerolog                           |
| Развертывание | Docker + Kubernetes (опционально) |

## 🚀 Запуск проекта

```bash
# Клонировать репозиторий
git clone https://github.com/Xanaduxan/social-network.git

```
