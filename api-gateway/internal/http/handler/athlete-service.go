package handler

import (
	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/athletepb"
	pbCountry "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/countrypb"
	"api-gateway/logger"
	"api-gateway/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Router /athletes [post]
// @Summary CREATE ATHLETE
// @Description This method creates athlete
// @Security BearerAuth
// @Tags ATHLETE
// @Accept json
// @Produce json
// @Param athlete body models.CreateAthleteRequest true "Athlete"
// @Success 200 {object} models.Athlete
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) CreateAthlete(c *gin.Context) {

	req := pb.CreateAthleteRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("CreateAthlete: Failed to bind JSON: ", err)
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}

	//Check Country Id
	if _, err := h.Service.GetCountry(&pbCountry.GetCountryRequest{Id: req.CountryId}); err != nil {
		logger.Error("CreateAthlete: Failed to get country: ", err)
		c.JSON(500, models.Message{Err: "Country with the provided ID does not exist or has been deleted"})
		return
	}

	resp, err := h.Service.CreateAthlete(&req)
	if err != nil {
		logger.Error("CreateAthlete: Failed to create athlete: ", err)
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("CreateAthlete: Athlete created successfully: ", logrus.Fields{
		"id":resp.Id,
		"name":resp.Name,
	})
	c.JSON(200, resp)
}

// @Router /athletes/{id} [get]
// @Summary GET ATHLETE
// @Description This method gets athlete
// @Security BearerAuth
// @Tags ATHLETE
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.GetAthleteResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) GetAthlete(c *gin.Context) {

	req := pb.GetAthleteRequest{}
	req.Id = c.Param("id")
	resp, err := h.Service.GetAthlete(&req)
	if err != nil {
		logger.Error("GetAthlete: Failed to get athlete with ID ", logrus.Fields{
			"id":req.Id,
		})
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("GetAthlete: Athlete retrieved successfully: ", logrus.Fields{
		"name":resp.Name,
	})
	c.JSON(200, resp)
}

// @Router /athletes [get]
// @Summary GET ATHLETES
// @Description This method gets athletes
// @Security BearerAuth
// @Tags ATHLETE
// @Accept json
// @Produce json
// @Success 200 {object} models.ListOfAthleteResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) ListOfAthlete(c *gin.Context) {

	resp, err := h.Service.ListOfAthlete(&pb.ListOfAthleteRequest{})
	if err != nil {
		logger.Error("ListOfAthlete: Failed to list athletes: ", err)
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}

	logger.Info("ListOfAthlete: Athletes retrieved successfully")
	c.JSON(200, resp)
}

// @Router /athletes/{id} [put]
// @Summary UPDATE ATHLETE
// @Description This method updates athlete
// @Security BearerAuth
// @Tags ATHLETE
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param athlete body models.UpdateAthleteRequest true "Athlete"
// @Success 200 {object} models.Athlete
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) UpdateAthlete(c *gin.Context) {

	req := pb.UpdateAthleteRequest{}
	req.Id = c.Param("id")
	if err := c.BindJSON(&req); err != nil {
		logger.Error("UpdateAthlete: Failed to bind JSON for athlete ID ", logrus.Fields{
			"id":req.Id,
		})
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}
	resp, err := h.Service.UpdateAthlete(&req)
	if err != nil {
		logger.Error("UpdateAthlete: Failed to update athlete with ID ", logrus.Fields{
			"id":req.Id,
		})
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("UpdateAthlete: Athlete updated successfully: ", logrus.Fields{
		"time":resp.UpdatedAt,
	})
	c.JSON(200, resp)
}

// @Router /athletes/{id} [delete]
// @Summary DELETE ATHLETE
// @Description This method deletes an athlete
// @Security BearerAuth
// @Tags ATHLETE
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.DeleteAthleteResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) DeleteAthlete(c *gin.Context) {

	req := pb.DeleteAthleteRequest{}
	req.Id = c.Param("id")
	resp, err := h.Service.DeleteAthlete(&req)
	if err != nil {
		logger.Error("DeleteAthlete: Failed to delete athlete with ID ", logrus.Fields{
			"id":req.Id,
		})
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}

	logger.Info("DeleteAthlete: Athlete deleted successfully: ", resp.Status)
	c.JSON(200, resp)
}
