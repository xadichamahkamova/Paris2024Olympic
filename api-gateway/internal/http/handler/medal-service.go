package handler

import (
	"context"
	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/medalspb"
	pbCountry "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/countrypb"
	pbEvent "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/eventpb"
	pbAthlete "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/athletepb"
	"api-gateway/logger"
	"api-gateway/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Router /medals [post]
// @Summary CREATE MEDAL
// @Description This method creates a medal
// @Security BearerAuth
// @Tags MEDAL
// @Accept json
// @Produce json
// @Param medal body models.CreateMedalRequest true "Medal"
// @Success 200 {object} models.Medal
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) CreateMedal(c *gin.Context) {

	req := pb.CreateMedalRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("CreateMedal: Failed to bind JSON: ", err)
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}
	
	//Check Country Id
	if _, err := h.Service.GetCountry(&pbCountry.GetCountryRequest{Id: req.CountryId}); err != nil {
		logger.Error("CreateMedal: Failed to get country: ", err)
		c.JSON(500, models.Message{Err: "Country with the provided ID does not exist or has been deleted"})
		return
	}
	//Check Event Id
	if _, err := h.Service.GetEvent(&pbEvent.GetEventRequest{Id: req.EventId}); err != nil {
		logger.Error("CreateMedal: Failed to get event: ", err)
		c.JSON(500, models.Message{Err: "Event with the provided ID does not exist or has been deleted"})
		return
	}
	//Check Athelete Id
	if _, err := h.Service.GetAthlete(&pbAthlete.GetAthleteRequest{Id: req.AthleteId}); err != nil {
		logger.Error("CreateMedal: Failed to get athlete: ", err)
		c.JSON(500, models.Message{Err: "Athlete with the provided ID does not exist or has been deleted"})
		return
	}

	resp, err := h.Service.CreateMedal(context.Background(), &req)
	if err != nil {
		logger.Error("CreateMedal: Failed to create medal: ", err)
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("CreateMedal: Medal created successfully: ", logrus.Fields{
		"id":   resp.Id,
		"type": resp.Type,
	})
	c.JSON(200, resp)
}

// @Router /medals/{id} [put]
// @Summary UPDATE MEDAL
// @Description This method updates a medal
// @Security BearerAuth
// @Tags MEDAL
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param medal body models.UpdateMedalRequest true "Medal"
// @Success 200 {object} models.Medal
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) UpdateMedal(c *gin.Context) {

	req := pb.UpdateMedalRequest{}
	req.Id = c.Param("id")
	if err := c.BindJSON(&req); err != nil {
		logger.Error("UpdateMedal: Failed to bind JSON for medal ID ", logrus.Fields{
			"id": req.Id,
		})
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}
	resp, err := h.Service.UpdateMedal(context.Background(), &req)
	if err != nil {
		logger.Error("UpdateMedal: Failed to update medal with ID ", logrus.Fields{
			"id": req.Id,
		})
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("UpdateMedal: Medal updated successfully: ", logrus.Fields{
		"id":   resp.Id,
		"name": resp.Type,
	})
	c.JSON(200, resp)
}

// @Router /medals/{id} [delete]
// @Summary DELETE MEDAL
// @Description This method deletes a medal
// @Security BearerAuth
// @Tags MEDAL
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.DeleteMedalResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) DeleteMedal(c *gin.Context) {

	req := pb.DeleteMedalRequest{}
	req.Id = c.Param("id")
	resp, err := h.Service.DeleteMedal(context.Background(), &req)
	if err != nil {
		logger.Error("DeleteMedal: Failed to delete medal with ID ", logrus.Fields{
			"id": req.Id,
		})
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}

	logger.Info("DeleteMedal: Medal deleted successfully")
	c.JSON(200, resp)
}

// @Router /medals/{id} [get]
// @Summary GET MEDAL BY ID
// @Description This method gets a medal by ID
// @Security BearerAuth
// @Tags MEDAL
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.GetMedalByIdResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) GetMedalById(c *gin.Context) {

	req := pb.GetMedalByIdRequest{}
	req.Id = c.Param("id")
	resp, err := h.Service.GetMedalById(context.Background(), &req)
	if err != nil {
		logger.Error("GetMedalById: Failed to get medal with ID ", logrus.Fields{
			"id": req.Id,
		})
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("GetMedalById: Medal retrieved successfully: ", logrus.Fields{
		"id":   resp.Id,
		"name": resp.Type,
	})
	c.JSON(200, resp)
}

// @Router /medals [get]
// @Summary GET MEDALS
// @Description This method gets all medals
// @Security BearerAuth
// @Tags MEDAL
// @Accept json
// @Produce json
// @Success 200 {object} models.GetMedalsResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) GetMedals(c *gin.Context) {

	resp, err := h.Service.GetMedals(context.Background(), &pb.VoidMedal{})
	if err != nil {
		logger.Error("GetMedals: Failed to get medals: ", err)
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}

	logger.Info("GetMedals: Medals retrieved successfully")
	c.JSON(200, resp)
}

// @Router /medals/filter [get]
// @Summary GET MEDALS BY FILTER
// @Description This method gets medals by filter
// @Security BearerAuth
// @Tags MEDAL
// @Accept json
// @Produce json
// @Param filter body models.GetMedalByFilterRequest true "Filter"
// @Success 200 {object} models.GetMedalByFilterResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) GetMedalByFilter(c *gin.Context) {
	
	req := pb.GetMedalByFilterRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("GetMedalByFilter: Failed to bind JSON: ", err)
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}
	resp, err := h.Service.GetMedalByFilter(context.Background(), &req)
	if err != nil {
		logger.Error("GetMedalByFilter: Failed to get medals by filter: ", err)
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("GetMedalByFilter: Medals retrieved successfully by filter")
	c.JSON(200, resp)
}
