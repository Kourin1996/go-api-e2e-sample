package main_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/k1LoW/runn"
)

func CompareModelObject(x, y interface{}) bool {
	return cmp.Diff(x, y, cmp.Options{
		cmpopts.IgnoreMapEntries(func(k string, v interface{}) bool {
			return k == "created_at" || k == "updated_at"
		}),
	}) == ""
}

func TestE2E(t *testing.T) {
	// jwt, err := helper.GetAuth0JWT(
	// )

	// if err != nil {
	// 	t.Errorf("failed to fetch JWT: %v", err)

	// 	return
	// }

	runner, err := runn.Load(
		"e2e/**/*.yaml",
		runn.T(t),
		runn.Runner("req", "http://localhost:8080"),
		runn.Func("customcompare", CompareModelObject),
		// runn.Var("token", jwt),
	)

	if err != nil {
		t.Fatal(err)
	}

	ctx := context.TODO()
	if err := runner.RunN(ctx); err != nil {
		t.Fatal(err)
	}
}
