package gin

import "github.com/gin-gonic/gin"

func setupRouter(engine *gin.Engine) {
	v1 := engine.Group("/apis/v1")
	// add middleware
	v1.Use()
	{

	}
}
