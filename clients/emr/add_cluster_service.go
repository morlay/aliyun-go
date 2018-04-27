package emr

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddClusterServiceRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64                         `position:"Query" name:"ResourceOwnerId"`
	Services        *AddClusterServiceServiceList `position:"Query" type:"Repeated" name:"Service"`
	ClusterId       string                        `position:"Query" name:"ClusterId"`
}

func (req *AddClusterServiceRequest) Invoke(client *sdk.Client) (resp *AddClusterServiceResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "AddClusterService", "", "")
	resp = &AddClusterServiceResponse{}
	err = client.DoAction(req, resp)
	return
}

type AddClusterServiceService struct {
	ServiceName string `name:"ServiceName"`
}

type AddClusterServiceResponse struct {
	responses.BaseResponse
	RequestId string
}

type AddClusterServiceServiceList []AddClusterServiceService

func (list *AddClusterServiceServiceList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]AddClusterServiceService)
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
