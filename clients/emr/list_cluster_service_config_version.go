package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	RequestId         common.String
	PageNumber        common.Integer
	PageSize          common.Integer
	TotalCount        common.Integer
	ConfigVersionList ListClusterServiceConfigVersionConfigVersionList
}

type ListClusterServiceConfigVersionConfigVersion struct {
	ConfigVersion common.String
	Applied       bool
	CreateTime    common.Long
	Author        common.String
	Comment       common.String
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
