package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifySnatEntryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SnatTableId          string `position:"Query" name:"SnatTableId"`
	SnatEntryId          string `position:"Query" name:"SnatEntryId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	SnatIp               string `position:"Query" name:"SnatIp"`
}

func (r ModifySnatEntryRequest) Invoke(client *sdk.Client) (response *ModifySnatEntryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ModifySnatEntryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "ModifySnatEntry", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		ModifySnatEntryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ModifySnatEntryResponse

	err = client.DoAction(&req, &resp)
	return
}

type ModifySnatEntryResponse struct {
	RequestId string
}
