package util

type Config struct {
	Properties struct {
		TimeInWeek []struct {
			Day       string `json:"day"`
			HourSlots []int  `json:"hourSlots"`
		} `json:"timeInWeek"`
		NotAllowedTime []struct {
			Start string `json:"start"`
			End   string `json:"end"`
		} `json:"notAllowedTime"`
	} `json:"properties"`
}

type EksAPIParameter struct {
	SubscriptionId    string
	ResourceGroupName string
	ResourceName      string
	ApiVersion        string
	Location          string
	ConfigName        string
	ConfigFile        Config
}

type Addon struct {
	AddonName     string
	ClusterName   string
	Message_      string
	NodegroupName string
}

type CloudError struct {
	// Error - Details about the error.
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// type CloudErrorBody struct {
// 	// Code - An identifier for the error. Codes are invariant and are intended to be consumed programmatically.
// 	Code *string `json:"code,omitempty" protobuf:"bytes,1,opt,name=code"`
// 	// Message - A message describing the error, intended to be suitable for display in a user interface.
// 	Message *string `json:"message,omitempty" protobuf:"bytes,2,opt,name=message"`
// 	// // Target - The target of the particular error. For example, the name of the property in error.
// 	// Target string `json:"target,omitempty" protobuf:"bytes,3,opt,name=target"`
// 	// // Details - A list of additional details about the error.
// 	// Details []CloudErrorBody `json:"details,omitempty" protobuf:"bytes,4,opt,name=details"`
// }
