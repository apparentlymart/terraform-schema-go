// Package tfschema contains types for describing the expected structure
// of a Terraform configuration block whose shape is not known until runtime.
//
// For example, this is used to describe the expected contents of a resource
// configuration block, which is defined by the corresponding provider plugin
// and thus not compiled into Terraform Core.
//
// A config schema initially describes the shape of configuration, but from it
// we can also derive the expected structure of the values that represent
// these objects in other artifacts such as the state and in plans.
//
// This package only contains the models for representing a schema once it's
// been obtained. This package does not itself know how to retrieve a schema
// from a provider plugin, but other packages that do interact with provider
// plugins may use these types to represent the schemas they obtain.
package tfschema
