package swagger

import (
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var HandlerFunc = ginSwagger.WrapHandler(swaggerFiles.Handler)
