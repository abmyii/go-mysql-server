// Copyright 2020-2021 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package expression

import (
	"fmt"
	"github.com/dolthub/go-mysql-server/sql"
)

type ExistsOperator struct {
	subquery sql.Expression
}

var _ sql.Expression = &ExistsOperator{}

func NewExistsOperator(query sql.Expression) *ExistsOperator {
	return &ExistsOperator{subquery: query}
}

func (e *ExistsOperator) String() string {
	return fmt.Sprintf("EXISTS %s", e.subquery.String())
}

func (e *ExistsOperator) Type() sql.Type {
	return sql.Boolean
}

func (e *ExistsOperator) Eval(ctx *sql.Context, row sql.Row) (interface{}, error) {
	subQueryResult, err := e.subquery.Eval(ctx, row)
	if err != nil {
		return nil, err
	}

	if _, ok := subQueryResult.([]interface{}); ok {
		return false, nil
	}

	return true, nil
}

func (e *ExistsOperator) WithChildren(ctx *sql.Context, children ...sql.Expression) (sql.Expression, error) {
	if len(children) != 1 {
		return nil, sql.ErrInvalidChildrenNumber.New(e, len(children), 1)
	}

	return NewExistsOperator(children[0]), nil
}

func (e *ExistsOperator) Resolved() bool {
	return e.subquery.Resolved()
}

func (e *ExistsOperator) IsNullable() bool {
	return false
}

func (e *ExistsOperator) Children() []sql.Expression {
	return []sql.Expression{e.subquery}
}




