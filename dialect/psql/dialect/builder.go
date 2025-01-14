package dialect

import (
	"strings"

	"github.com/stephenafamo/bob/expr"
)

//nolint:gochecknoglobals
var bmod = expr.Builder[Expression, Expression]{}

type Expression struct {
	expr.Chain[Expression, Expression]
}

func (Expression) New(exp any) Expression {
	var b Expression
	b.Base = exp
	return b
}

// Implements fmt.Stringer()
func (x Expression) String() string {
	w := strings.Builder{}
	x.WriteSQL(&w, Dialect, 1) //nolint:errcheck
	return w.String()
}

// BETWEEN SYMMETRIC a AND b
func (x Expression) BetweenSymmetric(a, b any) Expression {
	return bmod.X(expr.Join{Exprs: []any{
		x.Base, "BETWEEN SYMMETRIC", a, "AND", b,
	}})
}

// NOT BETWEEN SYMMETRIC a AND b
func (x Expression) NotBetweenSymmetric(a, b any) Expression {
	return bmod.X(expr.Join{Exprs: []any{
		x.Base, "NOT BETWEEN SYMMETRIC", a, "AND", b,
	}})
}
