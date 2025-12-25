//go:build generate

package tools

// no imports required for go:generate

// Format Terraform code for use in documentation.
// If you do not have Terraform installed, you can remove the formatting command, but it is suggested
// to ensure the documentation is formatted properly.
//go:generate terraform fmt -recursive ../examples/

// Generate documentation.
// Add go:generate in front of the below to add documentation generation. This is commented out to prevent manual
// changes in the documentation from being overwritten.
// go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-dir .. -provider-name ipam
