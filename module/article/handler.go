package article

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/test_cache_CQRS/infrastructure/httplib"
	logger "github.com/test_cache_CQRS/infrastructure/log"
	"github.com/test_cache_CQRS/infrastructure/validator"
	"github.com/test_cache_CQRS/module/primitive"
	"github.com/test_cache_CQRS/utils"

	"github.com/labstack/echo/v4"
)

type Http struct {
	serviceArticle InterfaceService
}

func NewHttp(serviceHealth InterfaceService) InterfaceHttp {
	return &Http{
		serviceArticle: serviceHealth,
	}
}

type InterfaceHttp interface {
	GroupArticle(group *echo.Group)
}

func (h *Http) GroupArticle(g *echo.Group) {
	g.GET("", h.GetListArticle)
	g.GET("/:id", h.DetailArticle)
	g.POST("", h.CreateArticle)
}

func (h *Http) GetListArticle(c echo.Context) error {
	logCtx := fmt.Sprintf("handler.GetListArticle")
	ctx := context.Background()

	paginationQuery, err := httplib.GetPaginationFromCtx(c)
	if err != nil {
		logger.Error(ctx, utils.ErrorLogFormat, err.Error(), logCtx, "httplib.GetPaginationFromCtx")
		return httplib.SetErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	query := c.QueryParam("query")
	if query != "" {
		if !utils.IsValidSanitizeSQL(query) {
			err = errors.New(primitive.QueryIsSuspicious)
			logger.Error(ctx, utils.ErrorLogFormat, err.Error(), logCtx, "utils.IsValidSanitizeSQL")
			return httplib.SetErrorResponse(c, http.StatusBadRequest, primitive.QueryIsSuspicious)
		}
	}

	author := c.QueryParam("author")
	if author != "" {
		if !utils.IsValidSanitizeSQL(author) {
			err = errors.New(primitive.QueryIsSuspicious)
			logger.Error(ctx, utils.ErrorLogFormat, err.Error(), logCtx, "utils.IsValidSanitizeSQL")
			return httplib.SetErrorResponse(c, http.StatusBadRequest, primitive.QueryIsSuspicious)
		}
	}

	param := primitive.ParameterArticleHandler{
		Query:  query,
		Author: author,
	}

	data, count, err := h.serviceArticle.GetListArticle(ctx, param, paginationQuery)
	if err != nil {
		logger.Error(ctx, utils.ErrorLogFormat, err.Error(), logCtx, "h.serviceArticle.GetListArticle")
		return httplib.SetErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return httplib.SetPaginationResponse(c,
		http.StatusOK,
		primitive.SuccessGetArticle,
		data,
		uint64(count),
		paginationQuery)

}

func (h *Http) CreateArticle(c echo.Context) error {
	logCtx := fmt.Sprintf("handler.CreateArticle")
	ctx := context.Background()

	var requestBody primitive.ArticleReq
	if err := c.Bind(&requestBody); err != nil {
		logger.Error(ctx, utils.ErrorLogFormat, err.Error(), logCtx, primitive.ErrorBindBodyRequest)
		return httplib.SetErrorResponse(c, http.StatusBadRequest, primitive.SomethingWrongWithTheBodyRequest)
	}

	errValidateStruct := validator.ValidateStructResponseSliceString(requestBody)
	if errValidateStruct != nil {
		logger.Error(ctx, logCtx, "validator.ValidateStructResponseSliceString got err : %v", errValidateStruct)
		return httplib.SetCustomResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil, errValidateStruct)
	}

	data, err := h.serviceArticle.RecordArticle(ctx, requestBody)
	if err != nil {
		logger.Error(ctx, utils.ErrorLogFormat, err.Error(), logCtx, "h.serviceArticle.GetListArticle")
		return httplib.SetErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return httplib.SetSuccessResponse(c, http.StatusOK, primitive.SuccessCreateArticle, data)

}

func (h *Http) DetailArticle(c echo.Context) error {
	logCtx := fmt.Sprintf("handler.CreateArticle")
	ctx := context.Background()

	idParam := c.Param("id")
	if idParam == "" {
		err := errors.New(primitive.ParamIdIsZeroOrNullString)
		logger.Error(ctx, utils.ErrorLogFormat, err.Error(), logCtx, "c.Param")
		return httplib.SetErrorResponse(c, http.StatusBadRequest, primitive.ParamIdIsZeroOrNullString)
	}

	idInt64, err := strconv.Atoi(idParam)
	if err != nil || idInt64 == 0 {
		err := errors.New(primitive.ParamIdIsZeroOrNullString)
		logger.Error(ctx, utils.ErrorLogFormat, err.Error(), logCtx, "strconv.Atoi")
		return httplib.SetErrorResponse(c, http.StatusBadRequest, primitive.ParamIdIsZeroOrNullString)
	}

	data, err := h.serviceArticle.GetDetailArticle(ctx, int64(idInt64))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error(ctx, utils.ErrorLogFormat, err.Error(), logCtx, "h.serviceArticle.GetDetailArticle")
			return httplib.SetErrorResponse(c, http.StatusNotFound, primitive.RecordArticleNotFound)
		}
		logger.Error(ctx, utils.ErrorLogFormat, err.Error(), logCtx, "h.serviceArticle.GetDetailArticle")
		return httplib.SetErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return httplib.SetSuccessResponse(c, http.StatusOK, primitive.SuccessCreateArticle, data)

}
