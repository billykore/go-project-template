package server

func (srv *Server) registerRoutes() {
	v1 := srv.router.Group("/api/v1")
	v1.GET("/example/:id", srv.exampleHandler.Get)
}
