package routes

userGroup := router.Group("/users")
{
    userGroup.POST("/register", controllers.RegisterUser)
    userGroup.POST("/login", controllers.LoginUser)
    userGroup.GET("/books", controllers.GetAvailableBooks)
    userGroup.POST("/orders", controllers.CreateOrder)
    // Add more user-related routes as needed
}

adminGroup := router.Group("/admin")
{
    // Use JWT middleware to verify admin role
    adminGroup.Use(middleware.AdminAuthMiddleware())

    adminGroup.POST("/books", controllers.CreateBook)
    adminGroup.PUT("/books/:id", controllers.UpdateBook)
    adminGroup.DELETE("/books/:id", controllers.DeleteBook)
    adminGroup.GET("/orders", controllers.GetOrders)
    adminGroup.PUT("/orders/:id", controllers.UpdateOrderStatus)
    // Add more admin-related routes as needed
}
