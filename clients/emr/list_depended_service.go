package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListDependedServiceRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ServiceName     string `position:"Query" name:"ServiceName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *ListDependedServiceRequest) Invoke(client *sdk.Client) (resp *ListDependedServiceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListDependedService", "", "")
	resp = &ListDependedServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListDependedServiceResponse struct {
	responses.BaseResponse
	RequestId     common.String
	ExistServices bool
	ServiceList   ListDependedServiceServiceList
}

type ListDependedServiceService struct {
	ServiceName        common.String
	ServiceDisplayName common.String
}

type ListDependedServiceServiceList []ListDependedServiceService

func (list *ListDependedServiceServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListDependedServiceService)
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
