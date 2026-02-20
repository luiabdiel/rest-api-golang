package controller

import (
	"net/http"
	"rest-api-golang/model"
	"rest-api-golang/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: usecase,
	}
}

func (pc *productController) GetProducts(ctx *gin.Context) {
	products, err := pc.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (pc *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	err := ctx.BindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := pc.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (pc *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("product_id")
	if id == "" {
		response := model.Response{
			Message: "Id do producto não pode ser nulo",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do producto precisa ser um número",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := pc.productUsecase.GetProductById(productId)
	if id == "" {

		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto não encontrado na base de dados",
		}

		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
