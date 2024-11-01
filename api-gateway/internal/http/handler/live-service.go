package handler

import (
	"api-gateway/logger"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/livepb"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *HandlerST) CreateLiveStream(ctx *gin.Context) {

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		logger.Error("failed to upgrade connection: ", err)
		return
	}

	defer conn.Close()
	c := conn.NetConn()

	for {
		var msg pb.LiveStream
		err := conn.ReadJSON(&msg)
		if err != nil {
			logger.Error("failed to read json:", err)
			break
		}

		logger.Info("Received message: from IP: ", c.RemoteAddr(), &msg)
		h.Service.CreateLive(&msg)

		if err := conn.WriteJSON(&msg); err != nil {
			logger.Error("Failed to write message: ", err)
			return
		}
	}
}

// @Router /live/{eventId} [get]
// @Summary Get Live Stream by Event ID
// @Description This method retrieves a live stream by event ID
// @Security BearerAuth
// @Tags Live Stream
// @Accept json
// @Produce json
// @Param eventId path string true "Event ID"
// @Success 200 {object} models.LiveStream
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) GetLiveStream(ctx *gin.Context) {
	eventId := ctx.Param("eventId")
	resp, err := h.Service.GetLive(&pb.GetStreamRequest{
		Id: eventId,
	})
	if err != nil {
		ctx.JSON(400, err.Error())
		return
	}
	ctx.JSON(200, resp)
}
