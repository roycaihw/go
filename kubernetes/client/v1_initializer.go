/*
 * Kubernetes
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

// Initializer is information about an initializer that has not yet completed.
type V1Initializer struct {

	// name of the process that is responsible for initializing this object.
	Name string `json:"name"`
}
