package handler

import (
	"log"

	"netprospe/db"
	"netprospe/internal/model"

	"github.com/gofiber/fiber/v2"
)

func ListeAgent(ctx *fiber.Ctx) error {
	db, err := db.Connection()
	if err != nil {
	}

	var agents []model.Agent

	if err := db.Table("dbo.liste_ca_cc_rm").Find(&agents).Error; err != nil {
		log.Println("Erreur :", err)
	}
	return ctx.Render("index", fiber.Map{
		"Agents": agents,
	})
}
