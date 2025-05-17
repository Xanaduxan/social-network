# social-network

Описание:
Это заготовка для социальной сети, включающая:
✅ Ленту постов (Post Service)
✅ Друзья/подписки (User Graph Service)
✅ Аналитика активности (Analytics Service)

Технологии:

Backend: Go (Gin, gRPC)
DB: PostgreSQL (шардирование при необходимости)
Кэш: Redis (счетчики, сессии)
Брокер: Kafka (ивенты лайков, сообщений)
Мониторинг: Prometheus + Grafana
Развертывание: Docker, Kubernetes (опционально)
