func ApplicationRoutes(r *gin.RouterGroup, handler *handler.ApplicationHandler) {
	r.POST("/applications", handler.ApplyJob)
}
