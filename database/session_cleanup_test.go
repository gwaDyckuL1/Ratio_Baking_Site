package database

import (
	"context"
	"database/sql"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSessionCleanupRunsAndStops(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var callCount atomic.Int32
	cleanup := func() error {
		callCount.Add(1)
		return nil
	}

	ticks := make(chan time.Time)

	var wg sync.WaitGroup
	testWG = &wg

	SessionCleanUp(ticks, ctx, cleanup)

	ticks <- time.Now()
	time.Sleep(5 * time.Millisecond)

	if callCount.Load() != 1 {
		t.Fatalf("Expected cleanup to run once, got %d", callCount.Load())
	}

	cancel()

	close(ticks)
	//time.Sleep(5 * time.Millisecond)

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		//good
	case <-time.After(100 * time.Millisecond):
		t.Fatalf("goroutine leak: cleanup goroutine did not exit after cancel.")
	}

	if callCount.Load() != 1 {
		t.Fatalf("Clean up continued to run after shutdown: %d", callCount.Load())
	}
}

func TestDeleteOldSession(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Database failed to open: %v", err)
	}

	createTable := `
		CREATE TABLE sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		session_token TEXT NOT NULL,
		last_active DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}

	_, err = db.Exec(`
		INSERT INTO sessions (user_id, session_token, last_active)
		VALUES (1, "old session", DATETIME("now", "-2 days")),
			   (2, "new session", DATETIME("now"))
	`)
	if err != nil {
		t.Fatalf("Failed to insert data into db: %v", err)
	}

	cleanUp := DeleteOldSessions(db)
	err = cleanUp()
	if err != nil {
		t.Fatalf("Clean up function failed: %v", err)
	}

	row, err := db.Query(`SELECT session_token FROM sessions`)
	if err != nil {
		t.Fatalf("Failed to get session_token after clean up: %v", err)
	}

	var remaining []string
	for row.Next() {
		var token string
		err := row.Scan(&token)
		if err != nil {
			t.Fatalf("Error scanning remaining rows: %v", err)
		}
		remaining = append(remaining, token)
	}

	if len(remaining) != 1 || remaining[0] != "new session" {
		t.Fatalf("Expected only 'new session' to remain, got: %v", err)
	}
}
