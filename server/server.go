package server

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"

	"github.com/z0mi3ie/recipes_back/handlers"
)

// Server will hold all server specific information
type Server struct {
	Router *gin.Engine
}

// Configure Routes
// route: /recipe
// POST   Add a new recipe            [x]
// GET    Get an existing recipe(s)   [x]
// PUT    Update an existing recipe
// DELETE Delete an existing recipe   [x]
func (s *Server) ConfigureRouters() {
	s.Router.POST("/recipe", handlers.AddRecipe)
	s.Router.GET("/recipe", handlers.GetAllRecipes)
	s.Router.GET("/recipe/:name", handlers.GetRecipe)
	s.Router.DELETE("/recipe/:name", handlers.DeleteRecipe)
}

func (s *Server) ConfigureMiddleWare() {
	// Apply cors middleware
	corsMiddleware := cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	})

	s.Router.Use(corsMiddleware)
}

func New() (s *Server) {
	s = &Server{Router: gin.Default()}

	s.ConfigureMiddleWare()
	s.ConfigureRouters()

	return
}

func (s *Server) Start() {
	s.Router.Run()
}
