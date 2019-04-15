package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	DBName         string
	TableName      string
	DataSourceName string
}

func (conf *MysqlConfig) Default() *MysqlConfig {
	if conf.DBName == "" {
		conf.DBName = "scrapy"
	}

	if conf.TableName == "" {
		conf.TableName = "topic"
	}

	if conf.DataSourceName == "" {
		log.Fatal("dataSourceName is empty")
	}
	return conf
}

//var RedisClient *redis.Client
type Store struct {
	db     *sql.DB
	ctx    context.Context
	config MysqlConfig
}

func (store Store) Close() {
	store.db.Close()
}

func (store *Store) Init(ctx context.Context, config MysqlConfig) error {
	//RedisClient = redis.NewClient(&redis.Options{
	//	//	Addr:     "127.0.0.1:6379",
	//	//	Password: "",
	//	//	DB:       0,
	//	//})

	db, err := sql.Open("mysql", config.DataSourceName)
	if err != nil {
		return err
	}

	// create database if not exist
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + config.DBName)
	if err != nil {
		return err
	}

	_, err = db.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (
	id INT PRIMARY KEY,
	title VARCHAR(255),
	url VARCHAR(255),
	comment_count INT,
	vote_up_count INT,
	excerpt	TEXT,
	type VARCHAR(20),
	author_name VARCHAR(20),
	author_url VARCHAR(255),
	author_avatar_url VARCHAR(255),
	author_type VARCHAR(20),
	author_user_type VARCHAR(20),
	created DATETIME,
	updated DATETIME
	)`, config.DBName, config.TableName))

	if err != nil {
		return err
	}

	store.db = db
	store.config = config
	store.ctx = ctx
	return nil
}

func formatTimeStamp(ts int64) string {
	tm := time.Unix(ts, 0)
	return tm.Format("2006-01-02 15:04:05")
}

//SaveTopics save the topic to mysql
func (store Store) SaveTopics(topic TopicResult) error {
	if store.db == nil {
		return fmt.Errorf("Store is not initialed or initial failed")
	}
	// insert or update
	stmtIns, err := store.db.Prepare(fmt.Sprintf("REPLACE INTO %s.%s VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )", store.config.DBName, store.config.TableName))
	if err != nil {
		return err
	}
	for _, detail := range topic.Content {
		row := detail.Target
		go func(row target) {
			select {
			case <-store.ctx.Done():
				return
			default:
				_, err := stmtIns.Exec(
					row.ID,
					row.Title,
					row.Url,
					row.CommentCount,
					row.VoteUpCount,
					row.Excerpt,
					row.Type,
					row.Author.Name,
					row.Author.Url,
					row.Author.AvatarUrl,
					row.Author.Type,
					row.Author.UserType,
					formatTimeStamp(row.Created),
					formatTimeStamp(row.Updated),
				)

				if err != nil {
					log.Printf("current Stmt execute failed, %v", err)
				}
			}

		}(row)
	}
	return nil
}
