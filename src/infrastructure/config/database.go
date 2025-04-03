package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() (*gorm.DB, error) {
	// Configuración de la base de datos
	username := "root"
	password := "roooooot"
	host := "127.0.0.1"
	port := "3306"
	dbName := "BaseHash"

	// Crear la cadena de conexión
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port)

	// Conectar a MySQL sin especificar la base de datos
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error conectando a MySQL: %w", err)
	}

	// Crear la base de datos si no existe
	err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)).Error
	if err != nil {
		return nil, fmt.Errorf("error creando la base de datos: %w", err)
	}

	// Reconectar especificando la base de datos
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error conectando a la base de datos %s: %w", dbName, err)
	}

	return db, nil
}
