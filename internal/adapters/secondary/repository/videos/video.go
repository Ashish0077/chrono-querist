package videos

// "database/sql"
// "fmt"
// "log"
// "os"

type Adapter struct {
	// db *sql.DB
}

func NewAdapter() *Adapter {
	// driverName := os.Getenv("driverName")
	// dbUserName := os.Getenv("userName")
	// dbPassword := os.Getenv("password")
	// dbPort := os.Getenv("port")
	// dbSource := os.Getenv("dataSource")

	// // Connect to DB
	// db, err := sql.Open(
	// 	driverName,
	// 	fmt.Sprintf("%v:%v@tcp(database:%v)/%v", dbUserName, dbPassword, dbPort, dbSource),
	// )
	// if err != nil {
	// 	log.Fatalf("DB Connection Failure: %v", err)
	// }

	// // Ping the DB
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatalf("DB Ping Failure: %v", err)
	// }

	return &Adapter{
		// db: db,
	}
}
