package graph

import (
	"context"
	"example/graph/model"
)


type Resolver struct{}


// similiar functions can be created for use 
func (r *Resolver) calculateSum(ctx context.Context, number_1 int, number_2 int) (*model.SummedAns, error) {
	var result = model.SummedAns{}

	result.FirstParam = number_1
	result.SecondParam = number_2
	result.TotalSum = number_1+number_2;

	return &result, nil
}