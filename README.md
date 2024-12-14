# golang-minify-url
URL Shortener Rest API Example written in Golang

## Dependencies
- [gocron](https://github.com/go-co-op/gocron)
- [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)
- [gin](https://github.com/gin-gonic/gin)
- [wire](https://github.com/google/wire)
- [viper](https://github.com/spf13/viper)
- [uber-zap](https://github.com/uber-go/zap)

## Project Structure

```
golang-minify-url
├── cmd/                            # Uygulamanın giriş noktaları
│   └── api/
│       ├── main.go                 # Uygulamanın ana dosyası
│       ├── wire.go                 # Wire setup
│       └── Dockerfile              # Dockerfile
├── internal/
│   ├── app/                        # İş mantığı
│   │   ├── controllers/            # HTTP handler'lar (controller katmanı)
│   │   │   └── controller.go
│   │   ├── services/               # İş mantığı (service katmanı)
│   │   │   └── service.go
│   │   ├── repositories/           # Veritabanı erişimi
│   │   │   └── repository.go
│   │   ├── entities/               # Domain entity'leri
│   │   │   └── entity.go
│   │   ├── models/                 # Genel modeller
│   │   │   └── model.go
│   │   ├── jobs/                   # Cron Jobs
│   │   │   └── scheduler.go
│   │   └── wire/                   # Wire bağımlılıkları
│   │       ├── provider.go         # Wire provider'lar
│   │       └── injector.go         # Wire injector
│   ├── middlewares/                # Middleware'lar
│   │   └── middleware.go
│   ├── database/
│   │   └── database.go             # Veritabanı erişim kodları
│   └── logger/                     # Loglama kodları
│       └── logger.go               # Loglama ayarlaması
├── configs/
│   ├── config.json                 # Config dosyası
│   ├── config.development.json             # Config dosyası - development
│   └── app_config.go               # Config kodları
├── docs/                           # API docs
├── pkg/                            # Tekrar kullanılabilir bağımsız paketler
│   └── common/                     # Ortak modeller ve yardımcı fonksiyonlar
├── go.mod                          # Go modül dosyası
├── go.sum                          # Go sum dosyası
├── README.md                       # Proje hakkında bilgiler
├── .gitignore                      # .gitignore dosyası
├── docker-compose.yml              # Docker-compose dosyası
└── .env                            # .env dosyası docker compose ve env değişkenlerini tanımlar
```
