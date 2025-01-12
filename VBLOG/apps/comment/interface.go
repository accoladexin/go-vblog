package comment

import "context"

type Service interface {
	CreateComment(ctx context.Context) error
}
