package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	for i, record := range records {
		if i > 0 {
			for _, row := range record {
				fmt.Println(row[0], row[0])
			}
		}
	}
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func connectToDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", Config("DB_HOST"), Config("DB_PORT"), Config("DB_USER"), Config("DB_PASSWORD"), Config("DB_NAME"), Config("DB_TIMEZONE"))
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error caught while connecting db")
		fmt.Println(err)
	}

	return DB

}

func importFromCSV(DB *gorm.DB, tableName string, filePath string) {
	importFromCSV := `
		COPY %s
		FROM '%s'
		DELIMITER ',' CSV HEADER;
	`

	importFromCSVCommand := fmt.Sprintf(importFromCSV, tableName, filePath)

	DB.Exec(importFromCSVCommand)
}

func createDatabase() {
	fmt.Println("Establishing connection with postgresql")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable TimeZone=%s", Config("DB_HOST"), Config("DB_PORT"), Config("DB_USER"), Config("DB_PASSWORD"), Config("DB_TIMEZONE"))
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error caught while establishing connection with postgresql")
		fmt.Println(err)
	}
	fmt.Println("Established connection with postgresql successfully")
	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", Config("DB_NAME"))
	DB.Exec(createDatabaseCommand)
	fmt.Println("Database created")
}

func migrateDatabase(DB *gorm.DB) {
	dropTableCommands := []string{
		`DROP TABLE IF EXISTS company CASCADE;`,
		`DROP TABLE IF EXISTS customer CASCADE;`,
		`DROP TABLE IF EXISTS ordertable CASCADE;`,
		`DROP TABLE IF EXISTS orderitem CASCADE;`,
		`DROP TABLE IF EXISTS delivery CASCADE;`,
	}

	createTableCommands := []string{
		`
			CREATE TABLE company (
				company_id BIGSERIAL NOT NULL PRIMARY KEY,
				company_name VARCHAR(50) NOT NULL
			);
		`,
		`
			CREATE TABLE customer (
				user_id VARCHAR(4) NOT NULL PRIMARY KEY,
				login VARCHAR(4) NOT NULL,
				password VARCHAR(5) NOT NULL,
				name VARCHAR(100) NOT NULL,
				company_id BIGINT REFERENCES company (company_id) NOT NULL,
				credit_cards VARCHAR(500)
			);
		`,
		`
			CREATE TABLE ordertable (
				id BIGSERIAL NOT NULL PRIMARY KEY,
				created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW () NOT NULL,
				order_name VARCHAR(50) NOT NULL,
				customer_id VARCHAR(4) REFERENCES customer (user_id) NOT NULL
			);
		`,
		`
			CREATE TABLE orderitem (
				id BIGSERIAL NOT NULL PRIMARY KEY,
				order_id BIGINT REFERENCES ordertable (id) NOT NULL,
				price_per_unit NUMERIC(10, 4),
				quantity SERIAL NOT NULL,
				product VARCHAR(100) NOT NULL
			);
		`,
		`
			CREATE TABLE delivery (
				id BIGSERIAL NOT NULL PRIMARY KEY,
				order_item_id BIGINT REFERENCES orderitem (id) NOT NULL,
				delivered_quantity SERIAL NOT NULL
			);
		`,
	}

	fmt.Println("Dropping tables company, customer, ordertable, orderitem, delivery if already exists")
	for _, command := range dropTableCommands {
		DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Exec(command).Error; err != nil {
				fmt.Println(err.Error())
				return err
			}

			return nil
		})
	}
	fmt.Println("Tables dropped company, customer, ordertable, orderitem, delivery if already exists")

	fmt.Println("Creating tables company, customer, ordertable, orderitem, delivery")
	for _, command := range createTableCommands {
		DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Exec(command).Error; err != nil {
				fmt.Println(err.Error())
				return err
			}

			return nil
		})

	}
	fmt.Println("Tables created company, customer, ordertable, orderitem, delivery")

	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`CREATE INDEX ON "ordertable" ("order_name");`).Error; err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println("Applying index on ordertable on order_name column")
		return nil
	})
	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`CREATE INDEX ON "orderitem" ("product");`).Error; err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println("Applying index on orderitem on product column")

		return nil
	})

}

func importDataIntoDatabase(DB *gorm.DB) {
	companyDataFile := "/Users/arsalan/test_data/Test task - Postgres - customer_companies.csv"
	customerDataFile := "/Users/arsalan/test_data/Test task - Postgres - customers.csv"
	orderDataFile := "/Users/arsalan/test_data/Test task - Postgres - orders.csv"
	orderItemsDataFile := "/Users/arsalan/test_data/Test task - Postgres - order_items.csv"
	deliveryDataFile := "/Users/arsalan/test_data/Test task - Postgres - deliveries.csv"

	var allFiles []map[string]string
	allFiles = append(allFiles, map[string]string{"name": "company", "file": companyDataFile})
	allFiles = append(allFiles, map[string]string{"name": "customer", "file": customerDataFile})
	allFiles = append(allFiles, map[string]string{"name": "ordertable", "file": orderDataFile})
	allFiles = append(allFiles, map[string]string{"name": "orderitem", "file": orderItemsDataFile})
	allFiles = append(allFiles, map[string]string{"name": "delivery", "file": deliveryDataFile})

	for _, tableObj := range allFiles {
		DB.Transaction(func(tx *gorm.DB) error {
			importFromCSV(tx, tableObj["name"], tableObj["file"])
			fmt.Printf("Data imported from csv into table %s \n", tableObj["name"])
			return nil
		})
	}
	fmt.Println("All data imported for all tables")
}

func main() {
	createDatabase()
	DB := connectToDatabase()
	migrateDatabase(DB)
	importDataIntoDatabase(DB)
}
