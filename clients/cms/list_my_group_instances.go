package cms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	req.InitWithApiInfo("Cms", "2017-03-01", "ListMyGroupInstances", "cms", "")
	resp = &ListMyGroupInstancesResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListMyGroupInstancesResponse struct {
	responses.BaseResponse
	RequestId    string
	Success      bool
	ErrorCode    int
	ErrorMessage string
	PageNumber   int
	PageSize     int
	Total        int
	Resources    ListMyGroupInstancesResourceList
}

type ListMyGroupInstancesResource struct {
	Id         int64
	AliUid     int64
	RegionId   string
	InstanceId string
	Category   string
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
