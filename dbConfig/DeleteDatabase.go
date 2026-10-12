package dbConfig

func DeleteDatabase() {
	DB.Migrator().DropTable()
}
