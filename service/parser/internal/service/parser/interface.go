package parser

import "context"

type Parser interface {
	Parse(ctx context.Context, sel, url string) (float64, error)
}
