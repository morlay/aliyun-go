package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteSnatEntryRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	SnatTableId          string `position:"Query" name:"SnatTableId"`
	SnatEntryId          string `position:"Query" name:"SnatEntryId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (r DeleteSnatEntryRequest) Invoke(client *sdk.Client) (response *DeleteSnatEntryResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteSnatEntryRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteSnatEntry", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DeleteSnatEntryResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteSnatEntryResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteSnatEntryResponse struct {
	RequestId string
}
