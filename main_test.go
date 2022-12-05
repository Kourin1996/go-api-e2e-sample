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
			// created_atとupdated_atのフィールドは比較をスキップする
			return k == "created_at" || k == "updated_at"
		}),
	}) == ""
}

func TestE2E(t *testing.T) {
	// JWTを使用する際は、JWTをテスト前に取得する
	// jwt, err := helper.GetAuth0JWT(
	// )

	// if err != nil {
	// 	t.Errorf("failed to fetch JWT: %v", err)

	// 	return
	// }

	runner, err := runn.Load(
		"e2e/**/*.yaml",
		runn.T(t),
		// ベースパスの指定
		runn.Runner(
			"req",
			"http://localhost:8080",
			runn.OpenApi3("schemas/book.yaml"),
		),
		// テスト中で関数を実行できるよう、Funcを使って指定する
		runn.Func("customcompare", CompareModelObject),
		// テスト中にJWTを指定できるよう、Varを使って指定する
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
