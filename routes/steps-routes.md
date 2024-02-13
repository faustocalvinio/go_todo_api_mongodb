1) definir el paquete

2) importar : net/http , chi, controllers, middleware

3) definir funcion Router
    - crear router, usar middleware, asignar home handler a la ruta get en / 
    - router.mount en la ruta /todo con el todohandler definido aca abajo
    - router.delete /todo/delete-completed con su handler
    - servir archivos estaticos
    - fileServer := http.FileServer(http.Dir("./static/"))
	- router.Handle("/static/*", http.StripPrefix("/static", fileServer))
    - retornar


3) definir funcion todoHandler
    - rter = new router
    - rter.group()
    - rg.Group(func(r chi.Router) {
		r.Get("/", controllers.FetchTodos)
		r.Post("/", controllers.CreateTodo)
		r.Put("/{id}", controllers.UpdateTodo)
		r.Delete("/{id}", controllers.DeleteOneTodo)
	})
    - retornar