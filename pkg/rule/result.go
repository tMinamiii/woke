package rule

import (
	"fmt"
	"go/token"
	"strings"
)

type Result struct {
	Rule     *Rule
	Match    string
	Position *token.Position
}

// Reason outputs the suggested alternatives for this rule
func (r *Result) Reason() string {
	return fmt.Sprintf("Instead of '%s', consider the following alternative(s): '%s'", r.Match, r.Rule.Alternatives)
}

func (r *Result) String() string {
	return fmt.Sprintf("[%s] %s", r.Position.String(), r.Reason())
}

type Results struct {
	Results []Result
}

// Add adds a single Result object onto the Results stack
func (rs *Results) Add(r *Result) {
	rs.Results = append(rs.Results, *r)
}

// Push pushes a list of Result objects onto the Results stack
func (rs *Results) Push(r ...Result) {
	for _, result := range r {
		rs.Add(&result)
	}
}

func (rs *Results) String() string {
	s := []string{}
	for _, r := range rs.Results {
		s = append(s, r.String())
	}
	return strings.Join(s, "\n")
}