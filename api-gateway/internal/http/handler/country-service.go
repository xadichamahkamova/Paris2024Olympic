package handler

import (
	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/countrypb"
	"api-gateway/logger"
	"api-gateway/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Router /countries [post]
// @Summary CREATE COUNTRY
// @Description This method creates a country
// @Security BearerAuth
// @Tags COUNTRY
// @Accept json
// @Produce json
// @Param country body models.CreateCountryRequest true "Country"
// @Success 200 {object} models.Country
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) CreateCountry(c *gin.Context) {

	req := pb.CreateCountryRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("CreateCountry: Failed to bind JSON: ", err)
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}
	resp, err := h.Service.CreateCountry(&req)
	if err != nil {
		logger.Error("CreateCountry: Failed to create country: ", err)
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("CreateCountry: Country created successfully: ", logrus.Fields{
		"id":   resp.Id,
		"name": resp.Name,
	})
	c.JSON(200, resp)
}

// @Router /countries/{id} [get]
// @Summary GET COUNTRY
// @Description This method gets a country by ID
// @Security BearerAuth
// @Tags COUNTRY
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.Country
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) GetCountry(c *gin.Context) {

	req := pb.GetCountryRequest{}
	req.Id = c.Param("id")
	resp, err := h.Service.GetCountry(&req)
	if err != nil {
		logger.Error("GetCountry: Failed to get country with ID ", logrus.Fields{
			"id": req.Id,
		})
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("GetCountry: Country retrieved successfully: ", logrus.Fields{
		"name": resp.Name,
	})
	c.JSON(200, resp)
}

// @Router /countries [get]
// @Summary GET COUNTRIES
// @Description This method gets a list of countries
// @Security BearerAuth
// @Tags COUNTRY
// @Accept json
// @Produce json
// @Success 200 {object} models.ListOfCountryResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) ListOfCountry(c *gin.Context) {

	resp, err := h.Service.ListOfCountry(&pb.ListOfCountryRequest{})
	if err != nil {
		logger.Error("ListOfCountry: Failed to list countries: ", err)
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("ListOfCountry: Countries retrieved successfully")
	c.JSON(200, resp)
}

// @Router /countries/{id} [put]
// @Summary UPDATE COUNTRY
// @Description This method updates a country
// @Security BearerAuth
// @Tags COUNTRY
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param country body models.UpdateCountryRequest true "Country"
// @Success 200 {object} models.Country
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) UpdateCountry(c *gin.Context) {

	req := pb.UpdateCountryRequest{}
	req.Id = c.Param("id")
	if err := c.BindJSON(&req); err != nil {
		logger.Error("UpdateCountry: Failed to bind JSON for country ID ", logrus.Fields{
			"id": req.Id,
		})
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}
	resp, err := h.Service.UpdateCountry(&req)
	if err != nil {
		logger.Error("UpdateCountry: Failed to update country with ID ", logrus.Fields{
			"id": req.Id,
		})
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("UpdateCountry: Country updated successfully: ", logrus.Fields{
		"time": resp.UpdatedAt,
	})
	c.JSON(200, resp)
}

// @Router /countries/{id} [delete]
// @Summary DELETE COUNTRY
// @Description This method deletes a country
// @Security BearerAuth
// @Tags COUNTRY
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.DeleteCountryResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) DeleteCountry(c *gin.Context) {

	req := pb.DeleteCountryRequest{}
	req.Id = c.Param("id")
	resp, err := h.Service.DeleteCountry(&req)
	if err != nil {
		logger.Error("DeleteCountry: Failed to delete country with ID ", logrus.Fields{
			"id": req.Id,
		})
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("DeleteCountry: Country deleted successfully: ", resp.Status)
	c.JSON(200, resp)
}
