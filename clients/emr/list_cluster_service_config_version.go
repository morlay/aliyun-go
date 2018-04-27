package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListClusterServiceConfigVersionRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
}

func (req *ListClusterServiceConfigVersionRequest) Invoke(client *sdk.Client) (resp *ListClusterServiceConfigVersionResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterServiceConfigVersion", "", "")
	resp = &ListClusterServiceConfigVersionResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterServiceConfigVersionResponse struct {
	responses.BaseResponse
	RequestId         string
	PageNumber        int
	PageSize          int
	TotalCount        int
	ConfigVersionList ListClusterServiceConfigVersionConfigVersionList
}

type ListClusterServiceConfigVersionConfigVersion struct {
	ConfigVersion string
	Applied       bool
	CreateTime    int64
	Author        string
	Comment       string
}

type ListClusterServiceConfigVersionConfigVersionList []ListClusterServiceConfigVersionConfigVersion

func (list *ListClusterServiceConfigVersionConfigVersionList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceConfigVersionConfigVersion)
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
