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

func (u *User) UsernameAvailable() bool {
	// res, _ := datastore.DB.Users.Find(bson.M{"username": strings.ToLower(u.Username)}).Count()
	var users []User

	var oneday int64 = 1000 * 60 * 60 * 24
	expiry := helper.Timestamp() - oneday
	_, err := deleteExpiredUser(expiry)
	if err != nil {
		fmt.Println("error find avail username")
		return false
	}
	// query := map[string]interface{}{
	// 	"query": map[string]interface{}{
	// 		"match": map[string]interface{}{
	// 			"username": strings.ToLower(u.Username),
	// 		},
	// 	},
	// }
	query := db.MatchFilterCondition(
		map[string]interface{}{"username": strings.ToLower(u.Username)},
		map[string]interface{}{
			"created_at": map[string]int64{
				"gt": expiry,
			}},
	)
	err = FindAll(query, db.IdxUsers, &users)
	// for idx := range users {
	// 	if helper.SliceContains(existedIDs, users[idx].ID) {
	// 		users[idx] = users[len(users)-1]
	// 		users = users[:len(users)-1]
	// 	}
	// }
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
	err := FindAll(queryDel, db.IdxUsers, listDel)
	if err != nil {
		fmt.Println("error find avail username")
		return nil, err
	}
	var existedIDs []string
	for idx := range *listDel {
		existedIDs = append(existedIDs, (*listDel)[idx].ID)
		db.Elastic.Delete(db.IdxUsers, (*listDel)[idx].ID)
	}
	return existedIDs, nil
}

func (u *User) EmailAvailable() bool {
	query := map[string]interface{}{
		"email": strings.ToLower(u.Email),
	}
	var users []User
	err := FindAll(query, db.IdxUsers, &users)
	if err != nil {
		fmt.Println("error find avail email")
	}
	return len(users) == 0
}

func (u *User) GetByID() {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"_id": u.ID,
			},
		},
	}
	err := FindOne(query, db.IdxUsers, &u)
	if err != nil {
		fmt.Println("error find avail email")
	}
}

func (u *User) DeleteByID() {
	db.Elastic.Delete(db.IdxUsers, u.ID)
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
		db.IdxUsers,                            // Index name
		strings.NewReader(string(userMarshal)), // Document body
		db.Elastic.Index.WithDocumentID(u.ID),  // Document ID
		db.Elastic.Index.WithRefresh("true"),   // Refresh
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
