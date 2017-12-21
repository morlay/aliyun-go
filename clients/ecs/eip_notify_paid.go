package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EipNotifyPaidRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	Data                 string `position:"Query" name:"Data"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	UserCidr             string `position:"Query" name:"UserCidr"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r EipNotifyPaidRequest) Invoke(client *sdk.Client) (response *EipNotifyPaidResponse, err error) {
	req := struct {
		*requests.RpcRequest
		EipNotifyPaidRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "EipNotifyPaid", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		EipNotifyPaidResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.EipNotifyPaidResponse

	err = client.DoAction(&req, &resp)
	return
}

type EipNotifyPaidResponse struct {
	RequestId string
	Data      string
	Code      string
	Message   string
	Success   bool
}
