package Ado

type AdoPools struct {
	Count int64 `json:"count"`
	Value []struct {
		AgentCloudID  int64 `json:"agentCloudId"`
		AutoProvision bool  `json:"autoProvision"`
		AutoSize      bool  `json:"autoSize"`
		AutoUpdate    bool  `json:"autoUpdate"`
		CreatedBy     struct {
			Links struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"_links"`
			Descriptor  string `json:"descriptor"`
			DisplayName string `json:"displayName"`
			ID          string `json:"id"`
			ImageURL    string `json:"imageUrl"`
			UniqueName  string `json:"uniqueName"`
			URL         string `json:"url"`
		} `json:"createdBy"`
		CreatedOn string `json:"createdOn"`
		ID        int64  `json:"id"`
		IsHosted  bool   `json:"isHosted"`
		IsLegacy  bool   `json:"isLegacy"`
		Name      string `json:"name"`
		Options   string `json:"options"`
		Owner     struct {
			Links struct {
				Avatar struct {
					Href string `json:"href"`
				} `json:"avatar"`
			} `json:"_links"`
			Descriptor  string `json:"descriptor"`
			DisplayName string `json:"displayName"`
			ID          string `json:"id"`
			ImageURL    string `json:"imageUrl"`
			UniqueName  string `json:"uniqueName"`
			URL         string `json:"url"`
		} `json:"owner"`
		PoolType   string `json:"poolType"`
		Scope      string `json:"scope"`
		Size       int64  `json:"size"`
		TargetSize int64  `json:"targetSize"`
	} `json:"value"`
}
