package main

import (
	"context"
	"github.com/julienschmidt/httprouter"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"net/http"
	"os"
	"routes/accounts"
	"routes/auth"
	"routes/experiences"
	"strconv"
)

var (
	DB_PORT string
	DB_ADDRESS string
	DB_NAME string
	SECRET string
	REGISTER_VERIFY_SECRET string
	APPROVAL_VERIFY_SECRET string
	TOKEN_EXPIRY uint64
	MAIL_USERNAME string
	MAIL_PASSWORD string
	MAIL_SENDER string
	SMTP_HOST string
	SMTP_PORT uint64
	API_DOMAIN string
	SITE_DOMAIN string
	DEBUG bool

	PORT uint64

	mongo_cli *mongo.Client
	mongo_db *mongo.Database

	router *httprouter.Router
)

func init() {
	var err error
	DB_PORT = os.Getenv("DB_PORT")
	DB_ADDRESS = os.Getenv("DB_ADDRESS")
	DB_NAME = os.Getenv("DB_NAME")
	SECRET = os.Getenv("SECRET")
	REGISTER_VERIFY_SECRET = os.Getenv("REGISTER_VERIFY_SECRET")
	APPROVAL_VERIFY_SECRET = os.Getenv("APPROVAL_VERIFY_SECRET")
	TOKEN_EXPIRY, err = strconv.ParseUint(os.Getenv("TOKEN_EXPIRY"), 10, 64)
	if err != nil {
		panic(err)
	}
	MAIL_USERNAME = os.Getenv("MAIL_USERNAME")
	MAIL_PASSWORD = os.Getenv("MAIL_PASSWORD")
	MAIL_SENDER = os.Getenv("MAIL_SENDER")
	SMTP_HOST = os.Getenv("SMTP_HOST")
	SMTP_PORT, err = strconv.ParseUint(os.Getenv("SMTP_PORT"), 10, 64)
	if err != nil {
		panic(err)
	}
	API_DOMAIN = os.Getenv("API_DOMAIN")
	SITE_DOMAIN = os.Getenv("SITE_DOMAIN")
	DEBUG, err = strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Starting ConnectUS API Server...")

	mongoctx := context.Background()
	mongoctx, cancel := context.WithCancel(mongoctx)
	defer cancel()
	connectDB(mongoctx)

	startRouter()
}

func startRouter() {
	router = httprouter.New()
	router.OPTIONS("/*all", func(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
		w.WriteHeader(404)
	})

	router.GET("/", func(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
		w.WriteHeader(200) // Default (root) route: successful ping
	})

	// TODO USE BEST PRACTICES: https://www.owasp.org/index.php/OWASP_Cheat_Sheet_Series

	// v1 Routes
	auth.Routes("/v1/auth", router)
	accounts.AccountRoutes("/v1/accounts", router)
	accounts.PersonalAccountsRoutes("/v1", router)
	experiences.ExperienceRoutes("/v1/experiences", router)
	experiences.OpportunityRoutes("/v1/opportunities", router)

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(int(PORT)), router))
}

func connectDB(ctx context.Context) {
	var err error
	mongo_cli, err = mongo.NewClient("mongodb://" + DB_ADDRESS + ":" + DB_PORT + "/" + DB_NAME)
	if err != nil {
		log.Fatalf("Could not connect to mongo: %v", err)
		os.Exit(1)
	}
	err = mongo_cli.Connect(ctx)
	if err != nil {
		log.Fatalf("Could not connect to mongo: %v", err)
		os.Exit(1)
	}
	mongo_db = mongo_cli.Database(DB_NAME)
}