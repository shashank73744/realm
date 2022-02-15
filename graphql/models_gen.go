// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphql

import (
	"github.com/kcarretto/realm/ent/credential"
)

type CreateCredentialInput struct {
	Principal string          `json:"Principal"`
	Secret    string          `json:"Secret"`
	Kind      credential.Kind `json:"Kind"`
}

type CreateTargetInput struct {
	Name             string `json:"Name"`
	ForwardConnectIP string `json:"ForwardConnectIP"`
}
