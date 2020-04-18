package resolver

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/tsuki42/graphql-meetup/validator"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func validation(ctx context.Context, v validator.Validation) bool {
	isValid, errors := v.Validate()

	if !isValid {
		for k, e := range errors {
			graphql.AddError(ctx, &gqlerror.Error{
				Message: e,
				Extensions: map[string]interface{}{
					"field": k,
				},
			})
		}
	}
	return isValid
}
