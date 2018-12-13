package main

import (
	"github.com/globalsign/mgo"
	"github.com/julienschmidt/httprouter"
	"interfaces-internal"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	// ENV
	DB_PORT                string
	DB_ADDRESS             string
	DB_NAME                string
	SECRET                 string
	REGISTER_VERIFY_SECRET string
	APPROVAL_VERIFY_SECRET string
	TOKEN_EXPIRY           int64
	MAIL_USERNAME          string
	MAIL_PASSWORD          string
	MAIL_SENDER            string
	SMTP_HOST              string
	SMTP_PORT              int
	SMTP_TLS               bool
	API_DOMAIN             string
	SITE_DOMAIN            string
	DEBUG                  bool

	PORT                      int
	BCRYPT_COST               int
	API_READ_TIMEOUT_SECONDS  int
	API_WRITE_TIMEOUT_SECONDS int

	// Global ref

	Database *mgo.Database

	IAccountCollection     *mgo.Collection
	IOpportunityCollection *mgo.Collection
	IPostCollection        *mgo.Collection

	router *httprouter.Router
)

func getEnv(key string, def string) string {
	e := os.Getenv(key)
	if e == "" {
		return def
	}
	return e
}

func init() {
	// Obtain environment variables
	var err error
	DB_PORT = getEnv("DB_PORT", "27017")
	DB_ADDRESS = getEnv("DB_ADDRESS", "localhost")
	DB_NAME = getEnv("DB_NAME", "api-server")
	SECRET = getEnv("SECRET", "secret")
	REGISTER_VERIFY_SECRET = getEnv("REGISTER_VERIFY_SECRET", "secret")
	APPROVAL_VERIFY_SECRET = getEnv("APPROVAL_VERIFY_SECRET", "secret")
	TOKEN_EXPIRY, err = strconv.ParseInt(getEnv("TOKEN_EXPIRY", "86400"), 0, 64)
	if err != nil {
		panic(err)
	}
	MAIL_USERNAME = getEnv("MAIL_USERNAME", "test@test.com")
	MAIL_PASSWORD = getEnv("MAIL_PASSWORD", "pass")
	MAIL_SENDER = getEnv("MAIL_SENDER", "test@test.com")
	SMTP_HOST = getEnv("SMTP_HOST", "test.com")
	SMTP_PORT, err = strconv.Atoi(getEnv("SMTP_PORT", "587"))
	if err != nil {
		panic(err)
	}
	SMTP_TLS = getEnv("SMTP_TLS", "true") == "true"
	API_DOMAIN = getEnv("API_DOMAIN", "localhost/api")
	SITE_DOMAIN = getEnv("SITE_DOMAIN", "localhost")
	DEBUG, err = strconv.ParseBool(getEnv("DEBUG", "false"))
	if err != nil {
		panic(err)
	}
	PORT, err = strconv.Atoi(getEnv("PORT", "3000"))
	if err != nil {
		panic(err)
	}
	BCRYPT_COST, err = strconv.Atoi(getEnv("BCRYPT_COST", "10"))
	if err != nil {
		panic(err)
	}
	API_READ_TIMEOUT_SECONDS, err = strconv.Atoi(getEnv("API_READ_TIMEOUT_SECONDS", "30"))
	if err != nil {
		panic(err)
	}
	API_WRITE_TIMEOUT_SECONDS, err = strconv.Atoi(getEnv("API_WRITE_TIMEOUT_SECONDS", "50"))
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Starting ConnectUS API Server...")
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Println("Received signal " + sig.String() + " from host")
		done <- true
		time.Sleep(30*time.Second)
		os.Exit(1)
	}()

	ConnectMongoDB()
	InitMailer(true)
	go StartRouter() // i love goroutines a lot

	log.Println("Completed initialization of api-server.")
	<-done
	log.Println("Exiting api-server...")
}

func StartRouter() {
	log.Println("Starting API router...")

	router = httprouter.New()
	router.OPTIONS("/*all", func(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
		w.WriteHeader(404)
	})

	router.GET("/", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		w.WriteHeader(200) // Default (root) route: successful ping
	})

	// TODO USE BEST PRACTICES: https://www.owasp.org/index.php/OWASP_Cheat_Sheet_Series

	// v1 API Routes
	AuthRoutes("/v1/auth", router)
	AccountRoutes("/v1/accounts", router)
	PersonalAccountsRoutes("/v1", router)
	ExperienceRoutes("/v1/experiences", router)
	OpportunityRoutes("/v1/opportunities", router)

	server := &http.Server{
		Addr: ":"+strconv.Itoa(int(PORT)),
		Handler: router,
		ReadTimeout:  time.Duration(API_READ_TIMEOUT_SECONDS) * time.Second,
		WriteTimeout: time.Duration(API_WRITE_TIMEOUT_SECONDS) * time.Second,
	}

	log.Println("Initialized router on port :" + strconv.Itoa(int(PORT)) + ".")
	log.Fatal(server.ListenAndServe()) // Start and serve API
}

func ConnectMongoDB() {
	log.Println("Connecting to MongoDB at " + DB_ADDRESS + ":" + DB_PORT)
	session, err := mgo.Dial(DB_ADDRESS + ":" + DB_PORT)
	if err != nil {
		log.Fatal(err)
	}

	session.SetMode(mgo.Monotonic, true)

	log.Println("Successfully connected.")

	Database = session.DB(DB_NAME)

	// Store collections

	IAccountCollection = Database.C("AccountModel")
	IOpportunityCollection = Database.C("OpportunityModel")
	IPostCollection = Database.C("PostModel")

	// Initialize indexes
	interfaces_internal.InitIAccountIndexes(IAccountCollection)
	interfaces_internal.InitIUserIndexes(IAccountCollection)
	interfaces_internal.InitIOrganizationIndexes(IAccountCollection)

	interfaces_internal.InitIOpportunityIndexes(IOpportunityCollection)

	interfaces_internal.InitIPostIndexes(IPostCollection)

}
