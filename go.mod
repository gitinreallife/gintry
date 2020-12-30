module gintry

go 1.15

require (
	github.com/denisenkom/go-mssqldb v0.9.0
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.16
	github.com/satori/go.uuid v1.2.0 // indirect
	modules/models v0.0.0-00010101000000-000000000000
)

replace modules/models => /models
replace modules/controller => /controller
