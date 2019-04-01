package app

import (
	"encoding/json"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"github.com/shredx/ep2-golang-graphql-backend/app/models"

	//initializing the mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	revel.OnAppStart(InitDB)
	revel.OnAppStart(StartGraphQL)
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}

//InitDB inits the database and tables.
//This function will migrate the existing tables
func InitDB() {
	/*
		Will connect to the database
		Will automigrate the following tables
	*/
	//Connecting to the database
	driver := revel.Config.StringDefault("db.driver", "mysql")
	connectString := revel.Config.StringDefault("db.connect", "root:root@locahost/test")

	Db, err := gorm.Open(driver, connectString)
	if err != nil {
		revel.AppLog.Error("Error while connecting to the DB", err.Error())
		return
	}

	models.DB = Db

	//automigrating the tables
	models.DB.AutoMigrate(&models.User{})
	models.DB.AutoMigrate(&models.Review{})
	models.DB.AutoMigrate(&models.Tag{})
	models.DB.AutoMigrate(&models.Product{})
}

//StartGraphQL service will start the graphql server
func StartGraphQL() {
	//registering the handler with http
	port := revel.Config.StringDefault("graphql.port", "9090")
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		result := graphql.Do(graphql.Params{
			Schema:        models.Schema,
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)
	})
	revel.AppLog.Info("Now server is running on port", port)
	revel.AppLog.Info("Test with Get      : curl -g 'http://localhost:" + port + "/graphql?query={hero{name}}'")
	go http.ListenAndServe(":"+port, nil)
}
