/*
 * OpenAPI definition
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type HealthCheckStatus struct {
	HealthResults []Health `json:"healthResults,omitempty"`
	SuppressedHealthResults []Health `json:"suppressedHealthResults,omitempty"`
	Healthy bool `json:"healthy,omitempty"`
}
