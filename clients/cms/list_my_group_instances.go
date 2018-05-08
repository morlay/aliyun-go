package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListMyGroupInstancesRequest struct {
	requests.RpcRequest
	Total      string `position:"Query" name:"Total"`
	GroupId    int64  `position:"Query" name:"GroupId"`
	PageSize   int    `position:"Query" name:"PageSize"`
	Category   string `position:"Query" name:"Category"`
	PageNumber int    `position:"Query" name:"PageNumber"`
}

func (req *ListMyGroupInstancesRequest) Invoke(client *sdk.Client) (resp *ListMyGroupInstancesResponse, err error) {
	req.InitWithApiInfo("Cms", "2018-03-08", "ListMyGroupInstances", "cms", "")
	resp = &ListMyGroupInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListMyGroupInstancesResponse struct {
	responses.BaseResponse
	RequestId    common.String
	Success      bool
	ErrorCode    common.Integer
	ErrorMessage common.String
	PageNumber   common.Integer
	PageSize     common.Integer
	Total        common.Integer
	Resources    ListMyGroupInstancesResourceList
}

type ListMyGroupInstancesResource struct {
	Id         common.Long
	AliUid     common.Long
	RegionId   common.String
	InstanceId common.String
	Category   common.String
}

type ListMyGroupInstancesResourceList []ListMyGroupInstancesResource

func (list *ListMyGroupInstancesResourceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListMyGroupInstancesResource)
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
