package handler

import (
	"goth-countries/components"
	"goth-countries/models"
	"net/http"
	"strings"
)

func CountryList(rw http.ResponseWriter, r *http.Request) {
	component := components.Countries(models.Countries)
	component.Render(r.Context(), rw)
}

func Country(rw http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	for _, country := range models.Countries {
		if country.Name == name {
			component := components.CountryDetail(country)
			component.Render(r.Context(), rw)
			return
		}
	}
	http.NotFound(rw, r)
}

func Search(rw http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("search"))
	var results []models.Country

	for _, country := range models.Countries {
		if strings.Contains(strings.ToLower(country.Name), query) {
			results = append(results, country)
		}
	}

	components.CountryList(results).Render(r.Context(), rw)
}
