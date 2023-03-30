package routes

import (
	"github.com/Apouzi/golang-shop/app/api/authorization"
	"github.com/go-chi/chi"
)

type Routes struct{}

func RouteDigest(digest *chi.Mux) *chi.Mux{
	r := Routes{}

	
	digest.Group(func(digest chi.Router){
		digest.Use(authorization.EnsureValidToken())
		digest.Post("/login", r.Login)
	})


	

	digest.Get("/", r.Index)
	

	return digest
}