package handler

import (
	"net/http"
	grpc_auth "tiny-letter/coordinator/cmd/grpc/client/auth"
	grpc_subscription "tiny-letter/coordinator/cmd/grpc/client/subscription"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

// HandlePublisherSubscription handles publisher subscription with 2PC
func (h *Handler) HandleConfirmPublisherSubscription(c *gin.Context) {
	var req struct {
		UserID int `json:"user_id" binding:"required"`
		PlanID int `json:"plan_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscriptionId, err := grpc_subscription.ConfirmPublisherSubscription(req.UserID, req.PlanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Subscription service prepare failed")
		return
	}

	err = grpc_auth.ConfirmPublisherSubscription(req.UserID, req.PlanID)
	if err != nil {
		grpc_subscription.RollbackConfirmPublisherSubscription(int(subscriptionId))
		c.JSON(http.StatusInternalServerError, "Auth service prepare failed")
		return
	}

	c.JSON(http.StatusOK, "Publisher subscription confirmed")
}

// HandlePublisherUnsubscription handles publisher unsubscription with 2PC
func (h *Handler) HandleRevokePublisherSubscription(c *gin.Context) {
	var req struct {
		UserID int `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	planId, err := grpc_subscription.RevokePublisherSubscription(req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Subscription service prepare failed")
		return
	}

	err = grpc_auth.RevokePublisherSubscription(req.UserID, int(planId))
	if err != nil {
		grpc_subscription.RollbackRevokePublisherSubscription(req.UserID, int(planId))
		c.JSON(http.StatusInternalServerError, "Auth service prepare failed")
		return
	}

	c.JSON(http.StatusOK, "Publisher unsubscription confirmed")
}

// HandleChangePublisherPlan handles plan changes for publishers with 2PC
func (h *Handler) HandleChangePublisherSubscription(c *gin.Context) {
	var req struct {
		UserID int `json:"user_id" binding:"required"`
		PlanID int `json:"plan_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscriptionID, oldPlanId, err := grpc_subscription.ChangePublisherSubscription(req.UserID, req.PlanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Subscription service prepare failed")
		return
	}

	err = grpc_auth.ChangePublisherSubscription(req.UserID, req.PlanID)
	if err != nil {
		grpc_subscription.RollbackChangePublisherSubscription(int(subscriptionID), int(oldPlanId))
		c.JSON(http.StatusInternalServerError, "Auth service prepare failed")
		return
	}

	// Phase 2: Commit
	c.JSON(http.StatusOK, "Publisher plan change confirmed")
}

// HandleJoinPublication handles joining a publication with 2PC
func (h *Handler) HandleJoinPublication(c *gin.Context) {
	var req struct {
		UserID        int  `json:"user_id" binding:"required"`
		PublicationID int  `json:"publication_id" binding:"required"`
		IsPremium     bool `json:"is_premium" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscriptionID, err := grpc_subscription.JoinPublication(req.UserID, req.PublicationID, req.IsPremium)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Subscription service prepare failed")
		return
	}

	err = grpc_auth.JoinPublication(req.UserID, req.PublicationID, req.IsPremium)
	if err != nil {
		grpc_subscription.RollbackJoinPublication(int(subscriptionID))
		c.JSON(http.StatusInternalServerError, "Auth service prepare failed")
		return
	}

	c.JSON(http.StatusOK, "Join publication confirmed")
}

// HandleLeavePublication handles leaving a publication with 2PC
func (h *Handler) HandleLeavePublication(c *gin.Context) {
	var req struct {
		UserID        int `json:"user_id" binding:"required"`
		PublicationID int `json:"publication_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isPremium, err := grpc_subscription.LeavePublication(req.UserID, req.PublicationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Subscription service prepare failed")
		return
	}

	err = grpc_auth.LeavePublication(req.UserID, req.PublicationID)
	if err != nil {
		grpc_subscription.RollbackLeavePublication(req.UserID, req.PublicationID, bool(isPremium))
		c.JSON(http.StatusInternalServerError, "Auth service prepare failed")
		return
	}

	c.JSON(http.StatusOK, "Leave publication confirmed")
}

// HandleChangeSubscriberPlan handles plan changes for subscribers with 2PC
func (h *Handler) ChangeSubscriberSubscription(c *gin.Context) {
	var req struct {
		UserID        int `json:"user_id" binding:"required"`
		PublicationID int `json:"publication_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscriptionID, err := grpc_subscription.ChangeSubscriberSubscription(req.UserID, req.PublicationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Subscription service prepare failed")
		return
	}

	err = grpc_auth.ChangeSubscriberSubscription(req.UserID, req.PublicationID)
	if err != nil {
		grpc_subscription.RollbackChangeSubscriberSubscription(int(subscriptionID))
		c.JSON(http.StatusInternalServerError, "Auth service prepare failed")
		return
	}

	c.JSON(http.StatusOK, "Subscriber plan change confirmed")
}
