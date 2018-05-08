package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId          common.String
	TotalCount         common.Integer
	PageNumber         common.Integer
	PageSize           common.Integer
	ClusterServiceList ListClusterServiceForAdminClusterServiceList
}

type ListClusterServiceForAdminClusterService struct {
	ServiceName        common.String
	ServiceDisplayName common.String
	ServiceVersion     common.String
	InstallStatus      bool
	ClientType         bool
	ServiceStatus      common.String
	HealthStatus       common.String
	NeedRestartInfo    common.String
	NotStartInfo       common.String
	AbnormalNum        common.Integer
	StoppedNum         common.Integer
	NeedRestartNum     common.Integer
	ServiceActionList  ListClusterServiceForAdminServiceActionList
}

type ListClusterServiceForAdminServiceAction struct {
	ServiceName   common.String
	ComponentName common.String
	ActionName    common.String
	Command       common.String
	DisplayName   common.String
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
