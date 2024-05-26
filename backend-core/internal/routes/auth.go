package routes

func (r *Routes) AuthRoutes() {

	router := r.App.Group("/auth")
	router.Post("/register", r.Controller.RegisterUserController)
	router.Post("/login", r.Controller.LoginUserController)
	router.Get("/refresh", r.Middleware.UserMiddleware, r.Controller.RefreshTokenController)
}
