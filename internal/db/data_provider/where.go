package data_provider

import (
	"strconv"
	"strings"

	"likemind/internal/db/op"
)

type condition struct {
	field string
	op    op.Operator
	value any
}

func writeWhere(query *strings.Builder, conds []condition, argShift int) []any {
	if len(conds) == 0 {
		return nil
	}

	args := make([]any, 0, len(conds))

	query.WriteString(" WHERE ")

	for i, cond := range conds {
		if i != 0 {
			query.WriteString(" AND ")
		}

		query.WriteString(cond.field)
		query.WriteRune(' ')
		query.WriteString(cond.op.String())
		query.WriteString(" $")
		query.WriteString(strconv.Itoa(argShift + i))

		args = append(args, cond.value)
	}

	return args
}
