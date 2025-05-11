package routes

import (
	"GoProjects/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.NovoAluno)
	r.GET("/aluno/:id", controllers.BuscaAlunoPorId)
	r.DELETE("/aluno/:id", controllers.DeletaAluno)
	r.PUT("/aluno/:id", controllers.AtualizaAluno)
	r.GET("/aluno/cpf/:cpf", controllers.BuscaAlunoPorCpf)
	r.Run(":8035")
}
