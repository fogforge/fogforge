package config

import (
	tflang "github.com/hashicorp/terraform/config/lang"
	tfast "github.com/hashicorp/terraform/config/lang/ast"
)

// InterpolationString is a specialization of string that can be parsed
// into an interpolation AST.
type InterpolationString string

// ParseInterpolations parses the string for Terraform-style interpolations and
// returns the resulting root AST node, or an error if the string is invalid.
func (v InterpolationString) ParseInterpolations() (tfast.Node, error) {
	return tflang.Parse(string(v))
}
