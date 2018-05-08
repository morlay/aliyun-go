package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListClusterHealthRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                               `position:"Query" name:"ResourceOwnerId"`
	ClusterIdLists  *ListClusterHealthClusterIdListList `position:"Query" type:"Repeated" name:"ClusterIdList"`
}

func (req *ListClusterHealthRequest) Invoke(client *sdk.Client) (resp *ListClusterHealthResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListClusterHealth", "", "")
	resp = &ListClusterHealthResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListClusterHealthResponse struct {
	responses.BaseResponse
	RequestId                 common.String
	ClusterHealthResponseList ListClusterHealthClusterHealthResponseList
}

type ListClusterHealthClusterHealthResponse struct {
	ClusterId             common.Long
	ServiceHealthInfoList ListClusterHealthServiceHealthInfoList
	HealthResult          ListClusterHealthHealthResult
}

type ListClusterHealthServiceHealthInfo struct {
	Key           common.String
	PassNumber    common.Integer
	ErrorNumber   common.Integer
	WarningNumber common.Integer
	UnKnownNumber common.Integer
}

type ListClusterHealthHealthResult struct {
	Key           common.String
	PassNumber    common.Integer
	ErrorNumber   common.Integer
	WarningNumber common.Integer
	UnKnownNumber common.Integer
}

type ListClusterHealthClusterIdListList []string

func (list *ListClusterHealthClusterIdListList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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

type ListClusterHealthClusterHealthResponseList []ListClusterHealthClusterHealthResponse

func (list *ListClusterHealthClusterHealthResponseList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterHealthClusterHealthResponse)
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

type ListClusterHealthServiceHealthInfoList []ListClusterHealthServiceHealthInfo

func (list *ListClusterHealthServiceHealthInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListClusterHealthServiceHealthInfo)
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
