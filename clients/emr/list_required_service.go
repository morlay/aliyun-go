package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ListRequiredServiceRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	EmrVersion      string `position:"Query" name:"EmrVersion"`
	ServiceNameList string `position:"Query" name:"ServiceNameList"`
}

func (req *ListRequiredServiceRequest) Invoke(client *sdk.Client) (resp *ListRequiredServiceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ListRequiredService", "", "")
	resp = &ListRequiredServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListRequiredServiceResponse struct {
	responses.BaseResponse
	RequestId        common.String
	NeedOtherService bool
	ServiceList      ListRequiredServiceServiceList
}

type ListRequiredServiceService struct {
	ServiceName        common.String
	ServiceDisplayName common.String
}

type ListRequiredServiceServiceList []ListRequiredServiceService

func (list *ListRequiredServiceServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListRequiredServiceService)
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
