package database

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

var testWG *sync.WaitGroup

func SessionCleanUp(ticker <-chan time.Time, ctx context.Context, cleanUp models.CleanupFunc) {

	go func() {
		for {
			if testWG != nil {
				testWG.Add(1)
				defer testWG.Done()
			}
			select {
			case <-ctx.Done():
				log.Println("Session clean up shutting down.", ctx.Err())
				return
			case <-ticker:
				err := cleanUp()
				if err != nil {
					log.Println("Error cleaning up old sessions. ", err)
				}
			}
		}
	}()
}

func DeleteOldSessions(db *sql.DB) models.CleanupFunc {
	return func() error {
		query := `DELETE
					FROM sessions
					WHERE julianday('now') - julianday(last_active) > 1`

		_, err := db.Exec(query)
		return err
	}
}
