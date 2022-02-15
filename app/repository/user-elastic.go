package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/ardhihdra/chirpbird/app/datautils"
	"github.com/ardhihdra/chirpbird/app/db"
	"github.com/ardhihdra/chirpbird/app/helper"
	"github.com/tidwall/gjson"
)

type userRepo struct{}

func NewUserElasticRepository() UserRepository {
	return &userRepo{}
}

var (
	oneday int64 = 1000 * 60 * 60 * 6
	Expiry       = helper.Timestamp() - oneday
)

func (r userRepo) FindByUsername(username string, exactmatch bool) (*[]datautils.User, error) {
	var u []datautils.User
	var query map[string]interface{}
	if exactmatch {
		query = db.MatchCondition(map[string]interface{}{
			"username": strings.ToLower(username),
		})
	} else {
		query = db.QueryString(map[string]interface{}{
			"fields": []string{"username"},
			"query":  fmt.Sprintf("*%s*", username),
		})
	}
	return &u, FindAll(query, db.IdxUsers, &u)
}

func (r userRepo) FindByEmail(email string) (*datautils.User, error) {
	var u *datautils.User
	query := db.MatchCondition(map[string]interface{}{"email": strings.ToLower(email)})
	return u, FindOne(query, db.IdxUsers, &u)
}

func (r userRepo) FindByID(ID string) (*datautils.User, error) {
	var u *datautils.User
	query := db.MatchCondition(map[string]interface{}{"id": strings.ToLower(ID)})
	return u, FindOne(query, db.IdxUsers, &u)
}

func (r userRepo) CheckExpiry(id string) (*datautils.User, error) {
	var u *datautils.User
	var i_id interface{} = id
	query := db.MatchFilterCondition(
		map[string]interface{}{"id": i_id},
		map[string]interface{}{
			"created_at": map[string]int64{
				"gt": Expiry,
			}},
	)
	return u, FindOne(query, db.IdxUsers, &u)
}

func (r userRepo) EmailAvailable(user datautils.User) bool {
	query := map[string]interface{}{
		"email": strings.ToLower(user.Email),
	}
	var users []datautils.User
	err := FindAll(query, db.IdxUsers, &users)
	if err != nil {
		fmt.Println("error find avail email")
	}
	return len(users) == 0
}

func (r userRepo) UsernameAvailable(u datautils.User) bool {
	var users []datautils.User
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
	err = FindAll(query, db.IdxUsers, &users)
	if err != nil {
		fmt.Println("error find avail username")
		return false
	}
	return len(users) == 0
}

func (r userRepo) CreateUser(u datautils.User) error {
	userMarshal, _ := json.Marshal(u)
	res, err := db.Elastic.Index(
		db.IdxUsers,
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

func (r userRepo) DeleteByID(id string) {
	db.Elastic.Delete(db.IdxUsers, id)
}

func FindOne(query map[string]interface{}, index string, usr **datautils.User) error {
	var b bytes.Buffer
	db.ExecuteQuery(query, index, &b)
	values := gjson.GetManyBytes(b.Bytes(), "hits.hits.0._id", "hits.hits.0._source")
	json.Unmarshal([]byte(values[1].String()), &usr)
	// usr.ID = values[0].String()
	return nil
}

func FindAll(query map[string]interface{}, dtype string, usrs *[]datautils.User) error {
	var b bytes.Buffer
	db.ExecuteQuery(query, dtype, &b)
	values := gjson.GetManyBytes(b.Bytes(), "hits.hits.0._id", "hits.hits")
	arrVal := values[1].Array()
	var usr datautils.User
	for i := range arrVal {
		json.Unmarshal([]byte(arrVal[i].Get("_source").String()), &usr)
		*usrs = append(*usrs, usr)
	}
	// json.Unmarshal([]byte(values[1].String()), &usr)
	return nil
}

func deleteExpiredUser(expiry int64) ([]string, error) {
	listDel := &[]datautils.User{}

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
