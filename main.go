package main

import (
	"Recu_ArqSoftware/src/application/usecases"
	"Recu_ArqSoftware/src/domain/entities"
	"Recu_ArqSoftware/src/infrastructure/config"
	"Recu_ArqSoftware/src/infrastructure/controllers"
	infraRepositories "Recu_ArqSoftware/src/infrastructure/repositories"
	infraServices "Recu_ArqSoftware/src/infrastructure/services"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Configurar base de datos
	db, err := config.SetupDatabase()
	if err != nil {
		log.Fatal("Error de base de datos:", err)
	}

	// Migrar modelo de usuario
	err = db.AutoMigrate(&entities.User{})
	if err != nil {
		log.Fatal("Error migrando usuario:", err)
	}

	// Inicializar dependencias (inversión de dependencias)
	userRepo := infraRepositories.NewMySQLUserRepository(db)
	passwordService := infraServices.NewBcryptPasswordService()
	createUserUseCase := usecases.NewCreateUserUseCase(userRepo, passwordService)
	createUserController := controllers.NewCreateUserController(createUserUseCase)

	// Configurar router
	router := gin.Default()
	api := router.Group("/api")
	users := api.Group("/users")
	{
		users.POST("", createUserController.Handle)
	}

	// Iniciar servidor
	log.Println("Servidor iniciado en :8080")
	log.Fatal(router.Run(":8080"))
}
