/*
 * HAProxy Data Plane API
 *
 * API for editing and managing haproxy instances. Provides process information, configuration management, haproxy stats and logs.  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 1.2
 * Contact: support@haproxy.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ProcessInfoHaproxy struct for ProcessInfoHaproxy
type ProcessInfoHaproxy struct {
	// Process id of the replying worker process
	Pid *int32 `json:"pid,omitempty"`
	// Number of spawned processes
	Processes *int32 `json:"processes,omitempty"`
	// HAProxy version release date
	ReleaseDate string `json:"release_date,omitempty"`
	// HAProxy uptime in s
	Uptime *int32 `json:"uptime,omitempty"`
	// HAProxy version string
	Version string `json:"version,omitempty"`
}
