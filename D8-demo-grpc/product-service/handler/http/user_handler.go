package handler

import (
	"net/http"

	"hacktiv/model"
	"hacktiv/usecase"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productUsecase usecase.IProductUsecase
}

func NewProductHandler(productUsecase usecase.IProductUsecase) productHandler {
	return productHandler{
		productUsecase: productUsecase,
	}
}

func (u productHandler) RegisterProductRoutes(e *echo.Echo) {
	e.GET("/products", u.GetAllProductHandler)
	e.POST("/products", u.CreateProductHandler)
}

// ShowAccount godoc
// @Summary      Create Account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Product
// @Failure      400  {object}  utils.HTTPError
// @Failure      404  {object}  utils.HTTPError
// @Failure      500  {object}  utils.HTTPError
// @Router       /accounts/{id} [get]
func (u productHandler) CreateProductHandler(c echo.Context) error {
	var product model.Product
	err := c.Bind(&product)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// fmt.Println("ACTIVE USER: ", product.ActiveProduct)

	err = u.productUsecase.CreateProduct(c.Request().Context(), product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

func (u productHandler) GetAllProductHandler(c echo.Context) error {
	products, err := u.productUsecase.GetAllProducts(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"products": products,
	})
}
