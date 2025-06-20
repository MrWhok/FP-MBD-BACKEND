package configuration

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/MrWhok/FP-MBD-BACKEND/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config Config) *gorm.DB {
	username := config.Get("SUPABASE_USERNAME")
	password := config.Get("SUPABASE_PASSWORD")
	host := config.Get("SUPABASE_HOST")
	port := config.Get("SUPABASE_PORT")
	dbName := config.Get("SUPABASE_DB_NAME")
	sslMode := config.Get("SUPABASE_SSLMODE")
	maxPoolOpen, err := strconv.Atoi(config.Get("SUPABASE_POOL_MAX_CONN"))
	maxPoolIdle, err := strconv.Atoi(config.Get("SUPABASE_POOL_IDLE_CONN"))
	maxPollLifeTime, err := strconv.Atoi(config.Get("SUPABASE_POOL_LIFE_TIME"))
	exception.PanicLogging(err)

	// Logger GORM
	loggerDb := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	// DSN string
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		username, password, host, port, dbName, sslMode,
	)

	// ✅ Gunakan PreferSimpleProtocol & matikan prepared statement caching
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // ✅ Hindari named prepared statement (stmtcache)
	}), &gorm.Config{
		Logger:      loggerDb,
		PrepareStmt: false, // ✅ Hindari caching statement di sisi GORM
	})
	exception.PanicLogging(err)

	// Konfigurasi koneksi pool
	sqlDB, err := db.DB()
	exception.PanicLogging(err)
	sqlDB.SetMaxOpenConns(maxPoolOpen)
	sqlDB.SetMaxIdleConns(maxPoolIdle)
	sqlDB.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)

	return db
}
