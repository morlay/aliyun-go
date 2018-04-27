package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	RequestId         string
	TotalCount        int
	PageNumber        int
	PageSize          int
	ConfigHistoryList ListClusterServiceConfigHistoryConfigHistoryList
}

type ListClusterServiceConfigHistoryConfigHistory struct {
	ServiceName    string
	ConfigVersion  string
	ConfigFileName string
	ConfigItemName string
	NewValue       string
	OldValue       string
	Applied        bool
	CreateTime     int64
	Author         string
	Comment        string
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
