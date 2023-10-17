package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-ticketing/app/models"
	"online-ticketing/app/services"
	"online-ticketing/app/validator"
	"online-ticketing/commons/utils"
	"online-ticketing/interfaces/serializers"
)

type SiteController struct {
	service   services.SiteService
	validator validator.CustomValidator
}

func NewSiteController(service services.SiteService, validator validator.CustomValidator) *SiteController {
	return &SiteController{service: service, validator: validator}
}

func (c *SiteController) GetSites(ctx *gin.Context) {
	request := new(models.GetSiteRequest)

	if err := ctx.ShouldBindQuery(request); err != nil {
		errMsg := c.validator.ParseError(err)
		utils.ToJSON(ctx).CustomResponse(http.StatusBadRequest, false, "Failed", "BAD_REQUEST", nil, errMsg)
		return
	}

	data := c.service.WithContext(ctx).GetSites(request)

	utils.ToJSON(ctx).CustomResponse(http.StatusOK, true, "Successfully", "SUCCESS", serializers.SerializeEvents(data), nil)
}

// func (c *SiteController) PatchUser(ctx *gin.Context) {
// 	request := new(models.UpdateUserRequest)

// 	if err := ctx.ShouldBindJSON(request); err != nil {
// 		errMsg := c.validator.ParseError(err)
// 		utils.ToJSON(ctx).CustomResponse(http.StatusBadRequest, false, "Failed", "BAD_REQUEST", nil, errMsg)
// 		return
// 	}

// 	id, err := strconv.Atoi(strings.TrimSpace(ctx.Param("id")))
// 	if err != nil {
// 		utils.ToJSON(ctx).CustomResponse(http.StatusInternalServerError, false, "Failed", "Internal Server Error", nil, err)
// 		return
// 	}

// 	c.service.WithContext(ctx).UpdateUser(id, request)

// 	utils.ToJSON(ctx).CustomResponse(http.StatusOK, true, "Successfully", "SUCCESS", nil, nil)
// }

func (c *SiteController) PostSite(ctx *gin.Context) {
	request := new(models.NewSiteRequest)

	if err := ctx.ShouldBindJSON(request); err != nil {
		errMsg := c.validator.ParseError(err)
		utils.ToJSON(ctx).CustomResponse(http.StatusBadRequest, false, "Failed", "BAD_REQUEST", nil, errMsg)
		return
	}

	c.service.WithContext(ctx).Store(request)
	utils.ToJSON(ctx).CustomResponse(http.StatusCreated, true, "Successfully", "SUCCESS", nil, nil)
}
