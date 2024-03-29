/*
 * MomenGo API
 *
 * Get weather data of airports
 *
 * API version: 1.0.0
 * Contact: apiteam@momengo.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Captor struct {
	Id int64 `json:"id,omitempty"`

	IATA string `json:"IATA,omitempty"`

	Type_ string `json:"type,omitempty"`

	QOS int64 `json:"QOS,omitempty"`
}
