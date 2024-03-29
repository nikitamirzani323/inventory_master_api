package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/controllers"
	"github.com/nikitamirzani323/BTANGKAS_SUPER_API/middleware"
)

func Init() *fiber.App {
	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		// Set some security headers:
		// c.Set("Content-Security-Policy", "frame-ancestors 'none'")
		// c.Set("X-XSS-Protection", "1; mode=block")
		// c.Set("X-Content-Type-Options", "nosniff")
		// c.Set("X-Download-Options", "noopen")
		// c.Set("Strict-Transport-Security", "max-age=5184000")
		// c.Set("X-Frame-Options", "SAMEORIGIN")
		// c.Set("X-DNS-Prefetch-Control", "off")

		// Go to next middleware:
		return c.Next()
	})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(compress.New())
	app.Get("/ipaddress", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":      fiber.StatusOK,
			"message":     "Success",
			"record":      "data",
			"BASEURL":     c.BaseURL(),
			"HOSTNAME":    c.Hostname(),
			"IP":          c.IP(),
			"IPS":         c.IPs(),
			"OriginalURL": c.OriginalURL(),
			"Path":        c.Path(),
			"Protocol":    c.Protocol(),
			"Subdomain":   c.Subdomains(),
		})
	})
	app.Get("/dashboard", monitor.New())

	app.Post("/api/login", controllers.CheckLogin)
	app.Post("/api/valid", middleware.JWTProtected(), controllers.Home)
	app.Post("/api/alladmin", middleware.JWTProtected(), controllers.Adminhome)
	app.Post("/api/detailadmin", middleware.JWTProtected(), controllers.AdminDetail)
	app.Post("/api/saveadmin", middleware.JWTProtected(), controllers.AdminSave)

	app.Post("/api/alladminrule", middleware.JWTProtected(), controllers.Adminrulehome)
	app.Post("/api/saveadminrule", middleware.JWTProtected(), controllers.AdminruleSave)

	app.Post("/api/employee", middleware.JWTProtected(), controllers.Employeehome)
	app.Post("/api/employeeshare", middleware.JWTProtected(), controllers.Employeeshare)
	app.Post("/api/employeesave", middleware.JWTProtected(), controllers.EmployeeSave)
	app.Post("/api/departement", middleware.JWTProtected(), controllers.Departementhome)
	app.Post("/api/departementsave", middleware.JWTProtected(), controllers.DepartementSave)
	app.Post("/api/catevendor", middleware.JWTProtected(), controllers.Catevendorhome)
	app.Post("/api/catevendorsave", middleware.JWTProtected(), controllers.CatevendorSave)
	app.Post("/api/vendor", middleware.JWTProtected(), controllers.Vendorhome)
	app.Post("/api/vendorshare", middleware.JWTProtected(), controllers.Vendorshare)
	app.Post("/api/vendorsave", middleware.JWTProtected(), controllers.VendorSave)
	app.Post("/api/curr", middleware.JWTProtected(), controllers.Currhome)
	app.Post("/api/currsave", middleware.JWTProtected(), controllers.CurrSave)
	app.Post("/api/uom", middleware.JWTProtected(), controllers.Uomhome)
	app.Post("/api/uomshare", middleware.JWTProtected(), controllers.Uomshare)
	app.Post("/api/uomsave", middleware.JWTProtected(), controllers.UomSave)
	app.Post("/api/branch", middleware.JWTProtected(), controllers.Branchhome)
	app.Post("/api/branchsave", middleware.JWTProtected(), controllers.BranchSave)
	app.Post("/api/warehouse", middleware.JWTProtected(), controllers.Warehousehome)
	app.Post("/api/warehousesave", middleware.JWTProtected(), controllers.WarehouseSave)
	app.Post("/api/warehousestorage", middleware.JWTProtected(), controllers.Warehousestoragehome)
	app.Post("/api/warehousestoragesave", middleware.JWTProtected(), controllers.WarehouseStorageSave)
	app.Post("/api/warehousestoragebin", middleware.JWTProtected(), controllers.WarehousestorageBinhome)
	app.Post("/api/warehousestoragebinsave", middleware.JWTProtected(), controllers.WarehouseStorageBinSave)

	app.Post("/api/merek", middleware.JWTProtected(), controllers.Merekhome)
	app.Post("/api/merekshare", middleware.JWTProtected(), controllers.Merekshare)
	app.Post("/api/mereksave", middleware.JWTProtected(), controllers.MerekSave)
	app.Post("/api/cateitem", middleware.JWTProtected(), controllers.Cateitemhome)
	app.Post("/api/cateitemsave", middleware.JWTProtected(), controllers.CateitemSave)
	app.Post("/api/item", middleware.JWTProtected(), controllers.Itemhome)
	app.Post("/api/itemshare", middleware.JWTProtected(), controllers.Itemshare)
	app.Post("/api/itemuom", middleware.JWTProtected(), controllers.Itemuom)
	app.Post("/api/itemsave", middleware.JWTProtected(), controllers.ItemSave)
	app.Post("/api/itemuomsave", middleware.JWTProtected(), controllers.ItemuomSave)
	app.Post("/api/itemuomdelete", middleware.JWTProtected(), controllers.ItemuomDelete)

	app.Post("/api/purchaserequest", middleware.JWTProtected(), controllers.Purchaserequesthome)
	app.Post("/api/purchaserequestdetail", middleware.JWTProtected(), controllers.Purchaserequestdetail)
	app.Post("/api/purchaserequestdetailview", middleware.JWTProtected(), controllers.Purchaserequestdetailview)
	app.Post("/api/purchaserequestsave", middleware.JWTProtected(), controllers.PurchaserequestSave)
	app.Post("/api/purchaserequeststatussave", middleware.JWTProtected(), controllers.PurchaserequeststatusSave)
	app.Post("/api/rfq", middleware.JWTProtected(), controllers.Rfqhome)
	app.Post("/api/rfqdetail", middleware.JWTProtected(), controllers.Rfqdetail)
	app.Post("/api/rfqsave", middleware.JWTProtected(), controllers.RfqSave)
	app.Post("/api/rfqstatussave", middleware.JWTProtected(), controllers.RfqstatusSave)
	app.Post("/api/po", middleware.JWTProtected(), controllers.Pohome)
	app.Post("/api/podetail", middleware.JWTProtected(), controllers.Pohome)
	app.Post("/api/posave", middleware.JWTProtected(), controllers.PoSave)
	app.Post("/api/postatussave", middleware.JWTProtected(), controllers.Pohome)

	return app
}
