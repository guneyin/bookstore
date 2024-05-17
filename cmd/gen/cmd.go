package gen

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/guneyin/bookstore/database"
	"github.com/guneyin/bookstore/util"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"time"
)

var Cmd = &cobra.Command{
	Use: "gen",
}

var testData = &cobra.Command{
	Use: "test",
	Run: func(cmd *cobra.Command, args []string) {
		testDataGenerator()
	},
}

func init() {
	Cmd.AddCommand(testData)
}

func testDataGenerator() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	err := database.Connect()
	if err != nil {
		panic(err)
	}

	r := resty.New().R().SetContext(ctx)
	db := database.DB.WithContext(ctx)

	err = util.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	err = generateUserData(ctx, r, db)
	if err != nil {
		panic(err)
	}
}

func truncateTable(db *gorm.DB, model any) {
	stmt := &gorm.Statement{DB: db}
	_ = stmt.Parse(model)
	tableName := stmt.Schema.Table

	db.Raw(fmt.Sprintf("TRUNCATE TABLE  %s;", tableName))
}
