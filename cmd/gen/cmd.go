package gen

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/guneyin/bookstore/database"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"os"
	"time"
)

var log *slog.Logger

var Cmd = &cobra.Command{
	Use: "gen",
}

var testData = &cobra.Command{
	Use: "mock",
	Run: func(cmd *cobra.Command, args []string) {
		mockDataGenerator()
	},
}

func init() {
	Cmd.AddCommand(testData)

	log = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func mockDataGenerator() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	err := database.Connect()
	if err != nil {
		panic(err)
	}

	r := resty.New().R().SetContext(ctx)
	db := database.GetDB(ctx, false)

	err = generateUserData(ctx, r, db)
	if err != nil {
		panic(err)
	}

	err = generateBookData(ctx, r, db)
	if err != nil {
		panic(err)
	}
}

func truncateTable(db *gorm.DB, model any) {
	stmt := &gorm.Statement{DB: db}
	_ = stmt.Parse(model)
	tableName := stmt.Schema.Table

	db.Exec(fmt.Sprintf("DELETE FROM %s;", tableName))
}

func fetchData[V any](ctx context.Context, r *resty.Request, url string) ([]V, error) {
	var list []V

	res, err := r.
		SetResult(&list).
		SetQueryParam("limit", "10").
		Get(url)
	if err != nil {
		log.ErrorContext(ctx, "error on api call", slog.String("msg", err.Error()))

		return nil, err
	}

	if res.StatusCode() >= http.StatusBadRequest {
		log.ErrorContext(ctx, "api returned error",
			slog.Int("status_code", res.StatusCode()),
			slog.String("status", res.Status()),
		)

		return nil, errors.New(res.Status())
	}

	return list, nil
}
