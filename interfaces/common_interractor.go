package interfaces

import (
	"context"
	"net/http"
)

type CommonInteractor interface {
	DivideMessage(context.Context, *http.Request)
	// UseCaseLineRepository(ctx context.Context)
}
