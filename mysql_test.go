package main

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"testing"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

func TestMysql(t *testing.T) {
	db, err := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/Scrapy")
	defer db.Close()
	if err != nil {
		t.Errorf("connect mysql failed err: %s\n", err.Error())
	}
	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)
	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)
	go func() {
		select {
		case <-appSignal:
			stop()
		}
	}()

	err = Ping(ctx, db)
	if err != nil {
		t.Errorf("cannot ping err:%s\n", err.Error())
	}
}

func Ping(ctx context.Context, db *sql.DB) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return err
	}
	return nil
}

func TestStore(t *testing.T) {
	store := &Store{}
	ctx, _ := context.WithCancel(context.Background())
	err := store.Init(ctx, "scrapy", "topic")
	if err != nil {
		t.Fatal(err)
	}
}
