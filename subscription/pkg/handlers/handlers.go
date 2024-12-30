package handlers

import (
	"net/http"
	"tiny-letter/subscription/pkg/db"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *db.Repository
}

func NewHandler(repo *db.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) HandlerSubscribePublisher(c *gin.Context) {
	var req db.PublisherSubscriptionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.SubscribePublisherPlan(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Publisher subscribed successfully")
}

func (h *Handler) HandlerSubscribePublication(c *gin.Context) {
	var req db.SubscriberSubscriptionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	if err := h.repo.SubscribeSubscriberPlan(req); err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, "Subscriber subscribed successfully")
}

func (h *Handler) HandlerChangeSubscriberSubscriptionPlan(c *gin.Context) {
	var req db.SubscriberChangePlanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	if err := h.repo.ChangeSubscriberSubscriptionPlan(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "Subscriber subscription plan changed successfully")
}

func (h *Handler) HandlerChangePublisherSubscriptionPlan(c *gin.Context) {
	var req db.PublisherChangePlanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	if err := h.repo.ChangePublisherSubscriptionPlan(req); err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, "Publisher subscription plan changed successfully")
}

func (h *Handler) HandleUnsubscribePublication(c *gin.Context) {
	var req db.UnsubscribeSubscriberRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	if err := h.repo.UnsubscriptionSubscriberPlan(req); err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, "Subscriber unsubscribed successfully")
}

func (h *Handler) HandleUnsubscribeSubscriber(c *gin.Context) {
	var req db.UnsubscribePublisherRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}

	if err := h.repo.UnsubscriptionPublisherPlan(req); err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, "Publisher unsubscribed successfully")
}
