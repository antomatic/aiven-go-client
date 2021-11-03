// Copyright (c) 2017 jelmersnoeck
// Copyright (c) 2018 Aiven, Helsinki, Finland. https://aiven.io/

package aiven

type (
	// GetServiceTypesResponse Aiven API request
	// GET https://api.aiven.io/v1/project/<project>/service_types
	ServiceTypesResponse struct {
		APIResponse
		ServiceTypes map[string]struct {
			Description            string                 `json:"description"`
			DefaultVersion         string                 `json:"default_version"`
			LatestAvailableVersion string                 `json:"latest_available_version"`
			UserConfigSchema       map[string]interface{} `json:"user_config_schema"`
			ServicePlans           []struct {
				ServicePlan      string      `json:"service_plan"`
				ServiceType      string      `json:"service_type"`
				NodeCount        int         `json:"node_count"`
				MaxMemoryPercent int         `json:"max_memory_percent"`
				BackupConfig     interface{} `json:"backup_config"`
				Regions          map[string]struct {
					DiskSpaceCapMB      int `json:"disk_space_cap_mb"`
					DiskSpaceMB         int `json:"disk_space_mb"`
					DiskSpaceStepMB     int `json:"disk_space_step_mb"`
					DiskSpaceGBPriceUSD int `json:"disk_space_gb_price_usd"`
					NodeCPUCount        int `json:"node_cpu_count"`
					NodeMemoryMB        int `json:"node_memory_mb"`
					PriceUSD            int `json:"price_usd"`
				} `json:"regions"`
			} `json:"service_plans"`
		} `json:"service_types"`
	}

	// ServicesHandler is the client that interacts with the Service Types API
	// endpoints on Aiven.
	ServiceTypesHandler struct {
		client *Client
	}
)

// Create creates the given Service on Aiven.
func (h *ServiceTypesHandler) Get(project string) (*ServiceTypesResponse, error) {
	path := buildPath("project", project, "service_types")
	bts, err := h.client.doGetRequest(path, nil)
	if err != nil {
		return nil, err
	}

	var r ServiceTypesResponse
	errR := checkAPIResponse(bts, &r)

	return &r, errR
}
