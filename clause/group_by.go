package clause

import (
	"io"

	"github.com/stephenafamo/bob/query"
)

type GroupBy struct {
	Groups   []any
	Distinct bool
	With     string // ROLLUP | CUBE
}

func (g *GroupBy) AppendGroup(e any) {
	g.Groups = append(g.Groups, e)
}

func (g *GroupBy) SetGroupWith(with string) {
	g.With = with
}

func (g *GroupBy) SetGroupByDistinct(distinct bool) {
	g.Distinct = distinct
}

func (g GroupBy) WriteSQL(w io.Writer, d query.Dialect, start int) ([]any, error) {
	w.Write([]byte("GROUP BY "))
	if g.Distinct {
		w.Write([]byte("DISTINCT "))
	}

	args, err := query.ExpressSlice(w, d, start, g.Groups, "", ", ", "")
	if err != nil {
		return nil, err
	}

	if g.With != "" {
		w.Write([]byte(" WITH "))
		w.Write([]byte(g.With))
	}

	return args, nil
}

type GroupingSet struct {
	Groups []query.Expression
	Type   string // GROUPING SET | CUBE | ROLLUP
}

func (g GroupingSet) WriteSQL(w io.Writer, d query.Dialect, start int) ([]any, error) {
	w.Write([]byte(g.Type))
	args, err := query.ExpressSlice(w, d, start, g.Groups, " (", ", ", ")")
	if err != nil {
		return nil, err
	}

	return args, nil
}