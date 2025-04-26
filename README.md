# metricFlow

**metricFlow** — это микросервис на Go для сбора метрик через gRPC.  

---
## Стек технологий

- **Go** (Golang)
- **gRPC** + **Protocol Buffers**
- **Docker / Docker Compose**
- **Prometheus + Grafana**

---

## ⚙️ Как запустить

### 1. Клонируйте репозиторий

```bash
git clone https://github.com/cihgt/metricFlow.git
cd metricFlow
```

### 2. Поднимите сервис через Docker

```bash
docker-compose up --build
```

Сервис будет доступен на порту `50051` по gRPC.

---
