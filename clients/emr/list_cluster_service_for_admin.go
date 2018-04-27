package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterServiceForAdminRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListClusterServiceForAdminRequest) Invoke(client *sdk.Client) (resp *ListClusterServiceForAdminResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterServiceForAdmin", "", "")
	resp = &ListClusterServiceForAdminResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterServiceForAdminResponse struct {
	responses.BaseResponse
	RequestId          string
	TotalCount         int
	PageNumber         int
	PageSize           int
	ClusterServiceList ListClusterServiceForAdminClusterServiceList
}

type ListClusterServiceForAdminClusterService struct {
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
	ServiceActionList  ListClusterServiceForAdminServiceActionList
}

type ListClusterServiceForAdminServiceAction struct {
	ServiceName   string
	ComponentName string
	ActionName    string
	Command       string
	DisplayName   string
}

type ListClusterServiceForAdminClusterServiceList []ListClusterServiceForAdminClusterService

func (list *ListClusterServiceForAdminClusterServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceForAdminClusterService)
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

type ListClusterServiceForAdminServiceActionList []ListClusterServiceForAdminServiceAction

func (list *ListClusterServiceForAdminServiceActionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceForAdminServiceAction)
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
