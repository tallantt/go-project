package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/golang-migrate/migrate/v4"
	migratepg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *gorm.DB

func InitDB() {
	// Читаем настройки из переменных окружения с дефолтами
	dbHost := getEnv("DB_HOST", "localhost")
	dbName := getEnv("DB_NAME", "postgres")
	dbUser := getEnv("DB_USER", "postgres")
	dbPass := getEnv("DB_PASSWORD", "123")
	dbPort := getEnv("DB_PORT", "5432")
	sslmode := getEnv("DB_SSLMODE", "disable")

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, sslmode)

	log.Printf("Connecting to DB at %s:%s...", dbHost, dbPort)

	// Открытие подключения к базе данных
	sqlDB, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	// Проверяем соединение (Ping)
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Не удалось подключиться к базе: %v", err)
	}

	// Инициализация драйвера для миграции
	driver, err := migratepg.WithInstance(sqlDB, &migratepg.Config{})
	if err != nil {
		log.Fatalf("Ошибка создания экземпляра миграции: %v", err)
	}

	// Создание миграции (путь к миграциям оставьте актуальным)
	m, err := migrate.NewWithDatabaseInstance("file://internal/db/migrations", "postgres", driver)
	if err != nil {
		log.Fatalf("Ошибка создания миграции: %v", err)
	}

	// Применение миграций
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("Миграции уже применены (нет изменений)")
		} else {
			log.Fatalf("Ошибка применения миграций: %v", err)
		}
	} else {
		log.Println("Миграции успешно применены")
	}

	// Открытие подключения через GORM
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка при открытии подключения с GORM: %v", err)
	}

	DB = gormDB
}

// Вспомогательная функция: получить переменную окружения с дефолтом
func getEnv(key string, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}
