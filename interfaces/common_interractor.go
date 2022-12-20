package interfaces

import (
	"context"
)

type CommonInteractor interface {
	DivideMessage(context.Context)
	// UseCaseLineRepository(ctx context.Context)
}
