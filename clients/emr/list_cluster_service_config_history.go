package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListClusterServiceConfigHistoryRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	PageSize        int    `position:"Query" name:"PageSize"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	PageNumber      int    `position:"Query" name:"PageNumber"`
	ConfigVersion   string `position:"Query" name:"ConfigVersion"`
}

func (req *ListClusterServiceConfigHistoryRequest) Invoke(client *sdk.Client) (resp *ListClusterServiceConfigHistoryResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterServiceConfigHistory", "", "")
	resp = &ListClusterServiceConfigHistoryResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterServiceConfigHistoryResponse struct {
	responses.BaseResponse
	RequestId         common.String
	TotalCount        common.Integer
	PageNumber        common.Integer
	PageSize          common.Integer
	ConfigHistoryList ListClusterServiceConfigHistoryConfigHistoryList
}

type ListClusterServiceConfigHistoryConfigHistory struct {
	ServiceName    common.String
	ConfigVersion  common.String
	ConfigFileName common.String
	ConfigItemName common.String
	NewValue       common.String
	OldValue       common.String
	Applied        bool
	CreateTime     common.Long
	Author         common.String
	Comment        common.String
}

type ListClusterServiceConfigHistoryConfigHistoryList []ListClusterServiceConfigHistoryConfigHistory

func (list *ListClusterServiceConfigHistoryConfigHistoryList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterServiceConfigHistoryConfigHistory)
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
