package handler

import (
	"api-gateway/logger"
	"api-gateway/models"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/eventpb"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Router /events [post]
// @Summary CREATE EVENT
// @Description This method creates an event
// @Security BearerAuth
// @Tags EVENT
// @Accept json
// @Produce json
// @Param event body models.CreateEventRequest true "Event"
// @Success 200 {object} models.Event
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) CreateEvent(c *gin.Context) {

	req := pb.CreateEventRequest{}
	if err := c.BindJSON(&req); err != nil {
		logger.Error("CreateEvent: Failed to bind JSON: ", err)
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}
	resp, err := h.Service.CreateEvent(&req)
	if err != nil {
		logger.Error("CreateEvent: Failed to create event: ", err)
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("CreateEvent: Event created successfully: ", logrus.Fields{
		"id":   resp.Id,
		"name": resp.Name,
	})
	c.JSON(200, resp)
}

// @Router /events/{id} [get]
// @Summary GET EVENT
// @Description This method gets an event by ID
// @Security BearerAuth
// @Tags EVENT
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.Event
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) GetEvent(c *gin.Context) {

	req := pb.GetEventRequest{}
	req.Id = c.Param("id")
	resp, err := h.Service.GetEvent(&req)
	if err != nil {
		logger.Error("GetEvent: Failed to get event with ID ", logrus.Fields{
			"id": req.Id,
		})
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("GetEvent: Event retrieved successfully: ", logrus.Fields{
		"name": resp.Name,
	})
	c.JSON(200, resp)
}

// @Router /events [get]
// @Summary GET EVENTS
// @Description This method gets a list of events
// @Security BearerAuth
// @Tags EVENT
// @Accept json
// @Produce json
// @Success 200 {object} models.ListOfEventResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) ListOfEvent(c *gin.Context) {

	resp, err := h.Service.ListOfEvent(&pb.ListOfEventRequest{})
	if err != nil {
		logger.Error("ListOfEvent: Failed to list events: ", err)
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("ListOfEvent: Events retrieved successfully")
	c.JSON(200, resp)
}

// @Router /events/{id} [put]
// @Summary UPDATE EVENT
// @Description This method updates an event
// @Security BearerAuth
// @Tags EVENT
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Param event body models.UpdateEventRequest true "Event"
// @Success 200 {object} models.Event
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) UpdateEvent(c *gin.Context) {

	req := pb.UpdateEventRequest{}
	req.Id = c.Param("id")
	if err := c.BindJSON(&req); err != nil {
		logger.Error("UpdateEvent: Failed to bind JSON for event ID ", logrus.Fields{
			"id": req.Id,
		})
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}
	resp, err := h.Service.UpdateEvent(&req)
	if err != nil {
		logger.Error("UpdateEvent: Failed to update event with ID ", logrus.Fields{
			"id": req.Id,
		})
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("UpdateEvent: Event updated successfully: ", logrus.Fields{
		"time": resp.UpdatedAt,
	})
	c.JSON(200, resp)
}

// @Router /events/{id} [delete]
// @Summary DELETE EVENT
// @Description This method deletes an event
// @Security BearerAuth
// @Tags EVENT
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.DeleteEventResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) DeleteEvent(c *gin.Context) {

	req := pb.DeleteEventRequest{}
	req.Id = c.Param("id")
	resp, err := h.Service.DeleteEvent(&req)
	if err != nil {
		logger.Error("DeleteEvent: Failed to delete event with ID ", logrus.Fields{
			"id": req.Id,
		})
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	logger.Info("DeleteEvent: Event deleted successfully: ", resp.Status)
	c.JSON(200, resp)
}
