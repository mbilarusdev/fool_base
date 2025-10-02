# Fool_base/v2

**Fool_base** — это библиотека, содержащая общие утилиты инфраструктуры для проекта **fool_game**.

Она включает следующие полезные инструменты:

- **db/postgres.go**: метод `ConnPGX`, использующий переменную окружения `"PGX"` для соединения с PostgreSQL. Формат строки подключения: `postgres://admin:secret@postgres:5432/mydb?sslmode=disable&pool_max_conns=4&pool_max_conn_lifetime=10m`.

- **cache/redis.go**: метод `PingRedis`, использующий переменную окружения `"RDB"` для проверки связи с Redis-сервером. Формат: `redis:6379&pass`.

- **log/logger.go**: метод `Init`, принимающий уровень логирования из переменной окружения `"LOG_LEVEL"` ("debug", "info", "warn", "error"). Дополнительно предоставляются удобные методы логирования: `Info`, `Err`, `Warn`.

- **net/network.go**: метод `Run`, принимающий объект роутинга (`mux.Router`), имя сервиса и слушающий адрес из переменной окружения `"ADDR"` (пример значения: `":8080"`).

---

📌 **Использование**:

\```go
import (
    "github.com/your-organization-name/fool_base/db"
    "github.com/your-organization-name/fool_base/cache"
    "github.com/your-organization-name/fool_base/log"
    "github.com/your-organization-name/fool_base/net"
)

// Example Usage:
pgxPool := db.ConnPGX()
rdb = cache.PingRedis()
err = net.Run(mux.NewRouter(), "MyService")
\```
