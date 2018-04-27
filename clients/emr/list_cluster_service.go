package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterServiceRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListClusterServiceRequest) Invoke(client *sdk.Client) (resp *ListClusterServiceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterService", "", "")
	resp = &ListClusterServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterServiceResponse struct {
	responses.BaseResponse
	RequestId          string
	TotalCount         int
	PageNumber         int
	PageSize           int
	ClusterServiceList ListClusterServiceClusterServiceList
}

type ListClusterServiceClusterService struct {
	ServiceName        string
	ServiceDisplayName string
	ServiceVersion     string
	InstallStatus      bool
	ClientType         bool
	ServiceStatus      string
	HealthStatus       string
	NeedRestartInfo    string
	NotStartInfo       string
	AbnormalNum        int
	StoppedNum         int
	NeedRestartNum     int
	ServiceActionList  ListClusterServiceServiceActionList
}

type ListClusterServiceServiceAction struct {
	ServiceName   string
	ComponentName string
	ActionName    string
	Command       string
	DisplayName   string
}

type ListClusterServiceClusterServiceList []ListClusterServiceClusterService

func (list *ListClusterServiceClusterServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceClusterService)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}

type ListClusterServiceServiceActionList []ListClusterServiceServiceAction

func (list *ListClusterServiceServiceActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceServiceAction)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
