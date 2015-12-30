// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_goto2.go (add link), as a start
// of the Traffic Ops golang data model

package todb

import (
	"encoding/json"
	"fmt"
	"gopkg.in/guregu/null.v3"
	"time"
)

type Parameter struct {
	Id          int64       `db:"id" json:"id"`
	Name        string      `db:"name" json:"name"`
	ConfigFile  string      `db:"config_file" json:"configFile"`
	Value       null.String `db:"value" json:"value"`
	LastUpdated time.Time   `db:"last_updated" json:"lastUpdated"`
}

func handleParameter(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getParameter(id)
	} else if method == "POST" {
		return postParameter(payload)
	} else if method == "PUT" {
		return putParameter(id, payload)
	} else if method == "DELETE" {
		return delParameter(id)
	}
	return nil, nil
}

func getParameter(id int) (interface{}, error) {
	ret := []Parameter{}
	if id >= 0 {
		err := globalDB.Select(&ret, "select * from parameter where id=$1", id)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	} else {
		queryStr := "select * from parameter"
		err := globalDB.Select(&ret, queryStr)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	return ret, nil
}

func postParameter(payload []byte) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	if err != nil {
		fmt.Println(err)
	}
	sqlString := "INSERT INTO parameter("
	sqlString += "name"
	sqlString += ",config_file"
	sqlString += ",value"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ",:config_file"
	sqlString += ",:value"
	sqlString += ")"
	result, err := globalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func putParameter(id int, payload []byte) (interface{}, error) {
	var v Asn
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwirte the id in the payload
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE parameter SET "
	sqlString += "name = :name"
	sqlString += ",config_file = :config_file"
	sqlString += ",value = :value"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := globalDB.NamedExec(sqlString, v)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}

func delParameter(id int) (interface{}, error) {
	result, err := globalDB.Exec("DELETE FROM parameter WHERE id=$1", id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result, err
}
