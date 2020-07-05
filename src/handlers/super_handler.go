package handlers

import (
	"github.com/gofiber/fiber"
	external "github.com/medson/superheroesAPI/src/domain/external_services"
	"github.com/medson/superheroesAPI/src/domain/models"
	"github.com/medson/superheroesAPI/src/domain/services"
	"github.com/medson/superheroesAPI/src/helpers"
)

// Handler struct contain the methods available to API
//
type Handler struct {
	superService services.SuperService
}

// Create Supers godoc
// @Summary Create Supers
// @Description Create Supers one or more supers searched in SuperHeroes API at database
// @Accept  json
// @Produce  json
// @Param name body models.SuperCreationReq true "Super hero or villain name"
// @Success 201 {object} []models.Super
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/v1/supers [post]
func (handler Handler) Create(ctx *fiber.Ctx) {
	var data models.SuperAPIResponse
	super := models.Super{}
	supers := []*models.Super{}

	ctx.BodyParser(&super)

	data, err := external.CallSuperAPI(super.Name)

	if data.Response == "error" {
		ctx.Status(400).JSON(err)
		return
	}

	supers, err = services.SuperService{}.Create(data)

	if err.Error != "" {
		ctx.Status(400).JSON(err)
		return
	}

	ctx.Status(201).JSON(&supers)
}

// Index godoc
// @Summary Show all supers
// @Description list out all supers stored in database
// @Accept  json
// @Produce  json
// @Param name query string false "name of one super to search"
// @Param superSide query string false "side of supers that will be listed, good or bad"
// @Param id query string false "id of one super"
// @Success 200 {array} []models.Super
// @Failure 400 {object} models.ErrorResponse
// @Router /api/v1/supers [get]
func (handler Handler) Index(ctx *fiber.Ctx) {
	var filter map[string]string
	filter = helpers.GetQueryParams(ctx.Fasthttp.QueryArgs())

	supers, err := services.SuperService{}.Get(filter)

	if err.Error != "" {
		ctx.Status(400).JSON(err)
		return
	}

	ctx.Status(200).JSON(supers)
}

// Destroy one super by ID godoc
// @Summary Delete one super
// @Description Delete one super stored in database
// @Accept  json
// @Produce  json
// @Param id path int true "Super ID"
// @Success 200 {object} object
// @Failure 404 {object} object
// @Router /api/v1/supers/{id} [delete]
func (handler Handler) Destroy(ctx *fiber.Ctx) {
	res, err := handler.superService.Destroy(ctx.Params("id"))

	if err.Error != "" {
		ctx.Status(404).JSON(err)
		return
	}
	ctx.Status(200).JSON(&fiber.Map{
		"sucess": res,
	})
}
