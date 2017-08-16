// Package dao contains the types for schema 'mj'.
package dao

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
	"time"
)

// AgAccount represents a row from 'mj.ag_account'.
type AgAccount struct {
	IndexID      int32     `json:"index_id"`      // index_id
	AgUpperID    int32     `json:"ag_upper_id"`   // ag_upper_id
	AgID         int32     `json:"ag_id"`         // ag_id
	AgLevel      int32     `json:"ag_level"`      // ag_level
	Password     string    `json:"password"`      // password
	Telephone    string    `json:"telephone"`     // telephone
	Realname     string    `json:"realname"`      // realname
	Weixin       string    `json:"weixin"`        // weixin
	Alipay       string    `json:"alipay"`        // alipay
	Email        string    `json:"email"`         // email
	Hongbao      int32     `json:"hongbao"`       // hongbao
	TotalBalance int32     `json:"total_balance"` // total_balance
	CreateTime   time.Time `json:"create_time"`   // create_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AgAccount exists in the database.
func (aa *AgAccount) Exists() bool {
	return aa._exists
}

// Deleted provides information if the AgAccount has been deleted from the database.
func (aa *AgAccount) Deleted() bool {
	return aa._deleted
}

// Insert inserts the AgAccount to the database.
func (aa *AgAccount) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if aa._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO mj.ag_account (` +
		`ag_upper_id, ag_id, ag_level, password, telephone, realname, weixin, alipay, email, hongbao, total_balance, create_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, aa.AgUpperID, aa.AgID, aa.AgLevel, aa.Password, aa.Telephone, aa.Realname, aa.Weixin, aa.Alipay, aa.Email, aa.Hongbao, aa.TotalBalance, aa.CreateTime)
	res, err := db.Exec(sqlstr, aa.AgUpperID, aa.AgID, aa.AgLevel, aa.Password, aa.Telephone, aa.Realname, aa.Weixin, aa.Alipay, aa.Email, aa.Hongbao, aa.TotalBalance, aa.CreateTime)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	aa.IndexID = int32(id)
	aa._exists = true

	return nil
}

// Update updates the AgAccount in the database.
func (aa *AgAccount) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !aa._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if aa._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE mj.ag_account SET ` +
		`ag_upper_id = ?, ag_id = ?, ag_level = ?, password = ?, telephone = ?, realname = ?, weixin = ?, alipay = ?, email = ?, hongbao = ?, total_balance = ?, create_time = ?` +
		` WHERE index_id = ?`

	// run query
	XOLog(sqlstr, aa.AgUpperID, aa.AgID, aa.AgLevel, aa.Password, aa.Telephone, aa.Realname, aa.Weixin, aa.Alipay, aa.Email, aa.Hongbao, aa.TotalBalance, aa.CreateTime, aa.IndexID)
	_, err = db.Exec(sqlstr, aa.AgUpperID, aa.AgID, aa.AgLevel, aa.Password, aa.Telephone, aa.Realname, aa.Weixin, aa.Alipay, aa.Email, aa.Hongbao, aa.TotalBalance, aa.CreateTime, aa.IndexID)
	return err
}

// Save saves the AgAccount to the database.
func (aa *AgAccount) Save(db XODB) error {
	if aa.Exists() {
		return aa.Update(db)
	}

	return aa.Insert(db)
}

// Delete deletes the AgAccount from the database.
func (aa *AgAccount) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !aa._exists {
		return nil
	}

	// if deleted, bail
	if aa._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM mj.ag_account WHERE index_id = ?`

	// run query
	XOLog(sqlstr, aa.IndexID)
	_, err = db.Exec(sqlstr, aa.IndexID)
	if err != nil {
		return err
	}

	// set deleted
	aa._deleted = true

	return nil
}

// AgAccountByIndexID retrieves a row from 'mj.ag_account' as a AgAccount.
//
// Generated from index 'ag_account_index_id_pkey'.
func AgAccountByIndexID(db XODB, indexID int32) (*AgAccount, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, ag_upper_id, ag_id, ag_level, password, telephone, realname, weixin, alipay, email, hongbao, total_balance, create_time ` +
		`FROM mj.ag_account ` +
		`WHERE index_id = ?`

	// run query
	XOLog(sqlstr, indexID)
	aa := AgAccount{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, indexID).Scan(&aa.IndexID, &aa.AgUpperID, &aa.AgID, &aa.AgLevel, &aa.Password, &aa.Telephone, &aa.Realname, &aa.Weixin, &aa.Alipay, &aa.Email, &aa.Hongbao, &aa.TotalBalance, &aa.CreateTime)
	if err != nil {
		return nil, err
	}

	return &aa, nil
}

// AgAccountsByAgUpperID retrieves a row from 'mj.ag_account' as a AgAccount.
//
// Generated from index 'idx_ag_upper_id'.
func AgAccountsByAgUpperID(db XODB, agUpperID int32) ([]*AgAccount, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, ag_upper_id, ag_id, ag_level, password, telephone, realname, weixin, alipay, email, hongbao, total_balance, create_time ` +
		`FROM mj.ag_account ` +
		`WHERE ag_upper_id = ?`

	// run query
	XOLog(sqlstr, agUpperID)
	q, err := db.Query(sqlstr, agUpperID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AgAccount{}
	for q.Next() {
		aa := AgAccount{
			_exists: true,
		}

		// scan
		err = q.Scan(&aa.IndexID, &aa.AgUpperID, &aa.AgID, &aa.AgLevel, &aa.Password, &aa.Telephone, &aa.Realname, &aa.Weixin, &aa.Alipay, &aa.Email, &aa.Hongbao, &aa.TotalBalance, &aa.CreateTime)
		if err != nil {
			return nil, err
		}

		res = append(res, &aa)
	}

	return res, nil
}

// AgAccountByAgID retrieves a row from 'mj.ag_account' as a AgAccount.
//
// Generated from index 'uidx_ag_id'.
func AgAccountByAgID(db XODB, agID int32) (*AgAccount, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`index_id, ag_upper_id, ag_id, ag_level, password, telephone, realname, weixin, alipay, email, hongbao, total_balance, create_time ` +
		`FROM mj.ag_account ` +
		`WHERE ag_id = ?`

	// run query
	XOLog(sqlstr, agID)
	aa := AgAccount{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, agID).Scan(&aa.IndexID, &aa.AgUpperID, &aa.AgID, &aa.AgLevel, &aa.Password, &aa.Telephone, &aa.Realname, &aa.Weixin, &aa.Alipay, &aa.Email, &aa.Hongbao, &aa.TotalBalance, &aa.CreateTime)
	if err != nil {
		return nil, err
	}

	return &aa, nil
}
