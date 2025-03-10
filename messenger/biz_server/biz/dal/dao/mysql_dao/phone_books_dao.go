// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)

package mysql_dao

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/jmoiron/sqlx"
	"github.com/fedigram/fedigram-server/messenger/biz_server/biz/dal/dataobject"
	"github.com/fedigram/fedigram-server/mtproto"
)

type PhoneBooksDAO struct {
	db *sqlx.DB
}

func NewPhoneBooksDAO(db *sqlx.DB) *PhoneBooksDAO {
	return &PhoneBooksDAO{db}
}

// insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on duplicate key update phone = values(phone), first_name = values(first_name), last_name = values(last_name)
// TODO(@benqi): sqlmap
func (dao *PhoneBooksDAO) InsertOrUpdate(do *dataobject.PhoneBooksDO) int64 {
	var query = "insert into phone_books(auth_key_id, client_id, phone, first_name, last_name) values (:auth_key_id, :client_id, :phone, :first_name, :last_name) on duplicate key update phone = values(phone), first_name = values(first_name), last_name = values(last_name)"
	r, err := dao.db.NamedExec(query, do)
	if err != nil {
		errDesc := fmt.Sprintf("NamedExec in InsertOrUpdate(%v), error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}

	id, err := r.LastInsertId()
	if err != nil {
		errDesc := fmt.Sprintf("LastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		glog.Error(errDesc)
		panic(mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_DBERR), errDesc))
	}
	return id
}
