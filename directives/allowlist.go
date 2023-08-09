package directives

import (
	"context"
	"fmt"
"strings"
	"github.com/99designs/gqlgen/graphql"
)

func AllowQuery(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	// fieldName := graphql.GetFieldName(ctx)
	fieldName := graphql.GetFieldContext(ctx).Field.Name
	operation := strings.ToLower(graphql.GetOperationContext(ctx).Operation.Name)
	allowedOperations := map[string]bool{
		"allPersons":  true,
		// "allPosts":    true,
		// "personById":  true,
		// "createPerson": true,
		// "updatePerson": true,
		// "deletePerson": true,
		// "createPost":   true,
		// "updatePost":   true,
		// "deletePost":   true,
	}

	allowedFields := map[string]bool{
		"Person.name":  true,
		// "Person.age":   true,
		// "Post.title":   true,
		// "Person.posts": true,
	}

	// Check if the operation is allowed
	if !allowedOperations[operation] {
		return nil, fmt.Errorf("operation not allowed: %s", operation)
	}

	// Check if the field is allowed
	if !allowedFields[fieldName] {
		return nil, fmt.Errorf("field not allowed: %s", fieldName)
	}

	// If both operation and field are allowed, continue with the resolver chain
	return next(ctx)
}
