package api_server

import (
	"github.com/globalsign/mgo"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
)

var (
	// ENV
	DB_PORT                string
	DB_ADDRESS             string
	DB_NAME                string
	SECRET                 string
	REGISTER_VERIFY_SECRET string
	APPROVAL_VERIFY_SECRET string
	TOKEN_EXPIRY           uint64
	MAIL_USERNAME          string
	MAIL_PASSWORD          string
	MAIL_SENDER            string
	SMTP_HOST              string
	SMTP_PORT              int
	API_DOMAIN             string
	SITE_DOMAIN            string
	DEBUG                  bool

	PORT uint64

	// Global ref

	Mailer *smtp.Client

	Database *mgo.Database

	IAccountCollection    *mgo.Collection
	IExperienceCollection *mgo.Collection

	router *httprouter.Router
)

func init() {
	// Obtain environment variables
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
	SMTP_PORT, err = strconv.Atoi(os.Getenv("SMTP_PORT"))
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

	ConnectMongoDB()
	InitMailer()
	StartRouter()
}

func StartRouter() {
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

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(int(PORT)), router)) // Start and serve API
}

func ConnectMongoDB() {
	session, err := mgo.Dial(DB_ADDRESS + ":" + DB_PORT)
	if err != nil {
		log.Fatal(err)
	}

	session.SetMode(mgo.Monotonic, true)

	Database = session.DB(DB_NAME)

	IAccountCollection = Database.C("AccountModel")
	IExperienceCollection = Database.C("ExperienceModel")
}