// Package dao contains the types for schema 'mj'.
package dao

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
	"time"
)

// Treasure represents a row from 'mj.treasure'.
type Treasure struct {
	IndexID    int32     `json:"index_id"`    // index_id
	PlayerID   int32     `json:"player_id"`   // player_id
	Reason     int32     `json:"reason"`      // reason
	Coins      int32     `json:"coins"`       // coins
	Cards      int32     `json:"cards"`       // cards
	ChangeTime time.Time `json:"change_time"` // change_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Treasure exists in the database.
func (t *Treasure) Exists() bool {
	return t._exists
}

// Deleted provides information if the Treasure has been deleted from the database.
func (t *Treasure) Deleted() bool {
	return t._deleted
}

// Insert inserts the Treasure to the database.
func (t *Treasure) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO mj.treasure (` +
		`player_id, reason, coins, cards, change_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, t.PlayerID, t.Reason, t.Coins, t.Cards, t.ChangeTime)
	res, err := db.Exec(sqlstr, t.PlayerID, t.Reason, t.Coins, t.Cards, t.ChangeTime)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	t.IndexID = int32(id)
	t._exists = true

	return nil
}

// Update updates the Treasure in the database.
func (t *Treasure) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if t._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE mj.treasure SET ` +
		`player_id = ?, reason = ?, coins = ?, cards = ?, change_time = ?` +
		` WHERE index_id = ?`

	// run query
	XOLog(sqlstr, t.PlayerID, t.Reason, t.Coins, t.Cards, t.ChangeTime, t.IndexID)
	_, err = db.Exec(sqlstr, t.PlayerID, t.Reason, t.Coins, t.Cards, t.ChangeTime, t.IndexID)
	return err
}

// Save saves the Treasure to the database.
func (t *Treasure) Save(db XODB) error {
	if t.Exists() {
		return t.Update(db)
	}

	return t.Insert(db)
}

// Delete deletes the Treasure from the database.
func (t *Treasure) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return nil
	}

	// if deleted, bail
	if t._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM mj.treasure WHERE index_id = ?`

	// run query
	XOLog(sqlstr, t.IndexID)
	_, err = db.Exec(sqlstr, t.IndexID)
	if err != nil {
		return err
	}

	// set deleted
	t._deleted = true

	return nil
}

// TreasuresByChangeTime retrieves a row from 'mj.treasure' as a Treasure.
//
// Generated from index 'idx_change_time'.
func TreasuresByChangeTime(db XODB, changeTime time.Time) ([]*Treasure, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, player_id, reason, coins, cards, change_time ` +
		`FROM mj.treasure ` +
		`WHERE change_time = ?`

	// run query
	XOLog(sqlstr, changeTime)
	q, err := db.Query(sqlstr, changeTime)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Treasure{}
	for q.Next() {
		t := Treasure{
			_exists: true,
		}

		// scan
		err = q.Scan(&t.IndexID, &t.PlayerID, &t.Reason, &t.Coins, &t.Cards, &t.ChangeTime)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}

	return res, nil
}

// TreasuresByPlayerID retrieves a row from 'mj.treasure' as a Treasure.
//
// Generated from index 'idx_player_id'.
func TreasuresByPlayerID(db XODB, playerID int32) ([]*Treasure, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, player_id, reason, coins, cards, change_time ` +
		`FROM mj.treasure ` +
		`WHERE player_id = ?`

	// run query
	XOLog(sqlstr, playerID)
	q, err := db.Query(sqlstr, playerID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Treasure{}
	for q.Next() {
		t := Treasure{
			_exists: true,
		}

		// scan
		err = q.Scan(&t.IndexID, &t.PlayerID, &t.Reason, &t.Coins, &t.Cards, &t.ChangeTime)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}

	return res, nil
}

// TreasuresByReason retrieves a row from 'mj.treasure' as a Treasure.
//
// Generated from index 'idx_reason'.
func TreasuresByReason(db XODB, reason int32) ([]*Treasure, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, player_id, reason, coins, cards, change_time ` +
		`FROM mj.treasure ` +
		`WHERE reason = ?`

	// run query
	XOLog(sqlstr, reason)
	q, err := db.Query(sqlstr, reason)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Treasure{}
	for q.Next() {
		t := Treasure{
			_exists: true,
		}

		// scan
		err = q.Scan(&t.IndexID, &t.PlayerID, &t.Reason, &t.Coins, &t.Cards, &t.ChangeTime)
		if err != nil {
			return nil, err
		}

		res = append(res, &t)
	}

	return res, nil
}

// TreasureByIndexID retrieves a row from 'mj.treasure' as a Treasure.
//
// Generated from index 'treasure_index_id_pkey'.
func TreasureByIndexID(db XODB, indexID int32) (*Treasure, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, player_id, reason, coins, cards, change_time ` +
		`FROM mj.treasure ` +
		`WHERE index_id = ?`

	// run query
	XOLog(sqlstr, indexID)
	t := Treasure{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, indexID).Scan(&t.IndexID, &t.PlayerID, &t.Reason, &t.Coins, &t.Cards, &t.ChangeTime)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
