// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/555082f38110f65b60d470107d211fc354a5c55a


package types

// DataframeAnalyticsAuthorization type.
//
// https://github.com/elastic/elasticsearch-specification/blob/555082f38110f65b60d470107d211fc354a5c55a/specification/ml/_types/Authorization.ts#L45-L57
type DataframeAnalyticsAuthorization struct {
	// ApiKey If an API key was used for the most recent update to the job, its name and
	// identifier are listed in the response.
	ApiKey *ApiKeyAuthorization `json:"api_key,omitempty"`
	// Roles If a user ID was used for the most recent update to the job, its roles at the
	// time of the update are listed in the response.
	Roles []string `json:"roles,omitempty"`
	// ServiceAccount If a service account was used for the most recent update to the job, the
	// account name is listed in the response.
	ServiceAccount *string `json:"service_account,omitempty"`
}

// NewDataframeAnalyticsAuthorization returns a DataframeAnalyticsAuthorization.
func NewDataframeAnalyticsAuthorization() *DataframeAnalyticsAuthorization {
	r := &DataframeAnalyticsAuthorization{}

	return r
}
