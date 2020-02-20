// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"database/sql"
	secret "github.com/uber/athenadriver/examples/constants"
	drv "github.com/uber/athenadriver/go"
	"log"
)

// main will query Athena and print all columns and rows information in csv format
func main() {
	// 1. Set AWS Credential in Driver Config.
	conf, err := drv.NewDefaultConfig(secret.OutputBucket, secret.Region,
		secret.AccessID, secret.SecretAccessKey)
	if err != nil {
		log.Fatal(err)
		return
	}
	// 2. Open Connection.
	dsn := conf.Stringify()
	db, _ := sql.Open(drv.DriverName, dsn)
	// 3. Query and print results
	rows, err := db.Query("select map(array['alice'], array['has a cat']);")
	if err != nil {
		println(err.Error())
		return
	}
	defer rows.Close()
	drv.ColsToCSV(rows)
	// array, map, binary, structure are returned as string type.
	var aMap string
	for rows.Next() {
		if err := rows.Scan(&aMap); err != nil {
			log.Fatal(err)
		}
	}
	println(aMap)
}

/*
Sample output:
_col0
{alice=has a cat}
*/
