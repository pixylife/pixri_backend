package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type UserACC struct {
	ID               int       `json:"id"`
	Login            string    `json:"login"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	Email            string    `json:"email"`
	ImageURL         string    `json:"imageUrl"`
	Activated        bool      `json:"activated"`
	LangKey          string    `json:"langKey"`
	CreatedBy        string    `json:"createdBy"`
	CreatedDate      time.Time `json:"createdDate"`
	LastModifiedBy   string    `json:"lastModifiedBy"`
	LastModifiedDate time.Time `json:"lastModifiedDate"`
	Authorities      []string  `json:"authorities"`
}

func AccountController(e *echo.Group, contextRoot string) {

	e.GET(contextRoot+"/account", getAccount)
}

func getAccount(c echo.Context) error {

	user := UserACC{3, "admin", "Administrator", "Administrator", "admin@localhost",
		"", true, "en", "system", time.Now(), "system",
		time.Now(), []string{"ROLE_USER", "ROLE_ADMIN"}}

	return c.JSON(http.StatusOK, user)
}
