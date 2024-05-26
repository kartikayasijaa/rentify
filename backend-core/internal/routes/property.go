package routes

func (r *Routes) PropertyRoutes() {

	router := r.App.Group("/property")
	router.Get("/", r.Middleware.UserMiddleware, r.Controller.PropertyGetController)
	router.Post("/", r.Middleware.SellerMiddleware, r.Controller.PropertyCreateController)
}
