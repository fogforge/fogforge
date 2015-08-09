package config

import (
	"testing"

	"github.com/hashicorp/terraform/config/lang/ast"
)

func TestInterpolationParsing(t *testing.T) {
	// Our interpolation parsing is really just a wrapper around Terraform's,
	// so we don't need to do any significant testing of our own.
	// Here we're just testing that our specialized string type works as
	// we expect it to.
	v := InterpolationString("Hello ${name}")
	node, err := v.ParseInterpolations()
	if err != nil {
		t.Fatalf("error in parsing: %s", err)
	}
	if _, ok := node.(*ast.Concat); !ok {
		t.Errorf("node is %#v; want *ast.Concat", node)
	}
}
