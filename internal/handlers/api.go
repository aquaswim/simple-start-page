package handlers

import (
	"github.com/gofiber/fiber/v2"
	"simple-start-page/internal/entities"
	"simple-start-page/internal/pkg"
	"simple-start-page/internal/services"
)

type handler struct {
	authService    services.Auth
	settingService services.Setting
}

func (h *handler) UpdateAuth(ctx *fiber.Ctx) error {
	a := new(entities.Auth)
	if err := ctx.BodyParser(a); err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	if err := pkg.ValidateStruct(a); err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	err := h.authService.UpdateAuth(a.Username, a.Password)
	if err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	return pkg.GenerateSuccessResponse(ctx, nil)
}

func (h *handler) GetListLink(ctx *fiber.Ctx) error {
	links, err := h.settingService.ListLink()
	if err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	return pkg.GenerateSuccessResponse(ctx, links)
}

func (h *handler) GetSetting(ctx *fiber.Ctx) error {
	setting, err := h.settingService.GetSetting()
	if err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	return pkg.GenerateSuccessResponse(ctx, setting)
}

func (h *handler) UpdateSetting(ctx *fiber.Ctx) error {
	s := new(entities.Setting)
	if err := ctx.BodyParser(s); err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	if err := pkg.ValidateStruct(s); err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	err := h.settingService.UpdateSetting(s)
	if err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	return pkg.GenerateSuccessResponse(ctx, nil)
}

func (h *handler) Login(ctx *fiber.Ctx) error {
	a := new(entities.Auth)
	if err := ctx.BodyParser(a); err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	if err := pkg.ValidateStruct(a); err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	authResult, err := h.authService.Login(a.Username, a.Password)
	if err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	return pkg.GenerateSuccessResponse(ctx, authResult)
}

func (h *handler) CekAuth(ctx *fiber.Ctx) error {
	a := new(entities.AuthResult)
	a.Token = ctx.Get("Authorization")
	if err := pkg.ValidateStruct(a); err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	if err := h.authService.CekLogin(a); err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	return ctx.Next()
}

type ApiHandler interface {
	GetListLink(ctx *fiber.Ctx) error
	GetSetting(ctx *fiber.Ctx) error
	UpdateSetting(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	CekAuth(ctx *fiber.Ctx) error
	UpdateAuth(ctx *fiber.Ctx) error
	UpdateLinks(ctx *fiber.Ctx) error
}

func (h *handler) UpdateLinks(ctx *fiber.Ctx) error {
	s := &struct {
		Links []entities.Link `json:"links"`
	}{}
	if err := ctx.BodyParser(s); err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	if err := pkg.ValidateStruct(s); err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	err := h.settingService.UpdateListLink(&s.Links)
	if err != nil {
		return pkg.GenerateErrorResponse(ctx, err)
	}
	return pkg.GenerateSuccessResponse(ctx, nil)
}

func NewApiHandler(auth services.Auth, setting services.Setting) ApiHandler {
	return &handler{
		authService:    auth,
		settingService: setting,
	}
}
