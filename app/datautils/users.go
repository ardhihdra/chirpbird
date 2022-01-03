package datautils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ardhihdra/chirpbird/db"
	"github.com/ardhihdra/chirpbird/helper"
	"github.com/tidwall/gjson"
)

type User struct {
	ID        string   `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Profile   string   `json:"profile"`
	Interests []string `json:"interests"`
	Country   string   `json:"country"`
	CreatedAt int64    `json:"created_at"`
	UpdatedAt int64    `json:"updated_at"`
}

var oneday int64 = 1000 * 60 * 60 * 6
var Expiry = helper.Timestamp() - oneday

func (u *User) UsernameAvailable() bool {
	var users []User
	expiry := Expiry

	_, err := deleteExpiredUser(expiry)
	if err != nil {
		fmt.Println("error find avail username")
		return false
	}
	// string to interface
	var uname interface{} = strings.ToLower(u.Username)
	query := db.MatchFilterCondition(
		map[string]interface{}{"username": uname},
		map[string]interface{}{
			"created_at": map[string]int64{
				"gt": expiry,
			}},
	)
	err = FindAll(query, db.TypeUsers, &users)
	if err != nil {
		fmt.Println("error find avail username")
		return false
	}
	return len(users) == 0
}

func deleteExpiredUser(expiry int64) ([]string, error) {
	listDel := &[]User{}

	queryDel := map[string]interface{}{
		"query": map[string]interface{}{
			"range": map[string]interface{}{
				"created_at": map[string]int64{
					"lt": expiry,
				},
			},
		},
	}
	err := FindAll(queryDel, db.TypeUsers, listDel)
	if err != nil {
		fmt.Println("error find avail username")
		return nil, err
	}
	var existedIDs []string
	for idx := range *listDel {
		existedIDs = append(existedIDs, (*listDel)[idx].ID)
		db.Elastic.Delete(db.TypeUsers, (*listDel)[idx].ID)
	}
	return existedIDs, nil
}

func (u *User) EmailAvailable() bool {
	query := map[string]interface{}{
		"email": strings.ToLower(u.Email),
	}
	var users []User
	err := FindAll(query, db.TypeUsers, &users)
	if err != nil {
		fmt.Println("error find avail email")
	}
	return len(users) == 0
}

func (u *User) GetByID() {
	query := db.MatchCondition(map[string]interface{}{"id": u.ID})
	err := FindOne(query, db.TypeUsers, &u)
	if err != nil {
		fmt.Println("error find avail email")
	}
}

func (u *User) DeleteByID() {
	db.Elastic.Delete(db.TypeUsers, u.ID)
}

func FindOne(query map[string]interface{}, index string, usr **User) error {
	var b bytes.Buffer
	db.ExecuteQuery(query, index, &b)
	values := gjson.GetManyBytes(b.Bytes(), "hits.hits.0._id", "hits.hits.0._source")
	json.Unmarshal([]byte(values[1].String()), &usr)
	// usr.ID = values[0].String()
	return nil
}

func FindAll(query map[string]interface{}, index string, usrs *[]User) error {
	var b bytes.Buffer
	db.ExecuteQuery(query, index, &b)
	values := gjson.GetManyBytes(b.Bytes(), "hits.hits.0._id", "hits.hits")
	arrVal := values[1].Array()
	var usr User
	for i := range arrVal {
		json.Unmarshal([]byte(arrVal[i].Get("_source").String()), &usr)
		*usrs = append(*usrs, usr)
	}
	// json.Unmarshal([]byte(values[1].String()), &usr)
	return nil
}

func (u *User) CreateUser() error {
	userMarshal, _ := json.Marshal(u)
	res, err := db.Elastic.Index(
		db.IdxMessaging,
		strings.NewReader(string(userMarshal)),    // Document body
		db.Elastic.Index.WithOpType(db.TypeUsers), // Index name
		db.Elastic.Index.WithDocumentID(u.ID),     // Document ID
		db.Elastic.Index.WithRefresh("true"),      // Refresh
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
		return err
	}
	defer res.Body.Close()
	if res.IsError() {
		db.PrintErrorResponse(res)
		return err
	}

	return nil
}
