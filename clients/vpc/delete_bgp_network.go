package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteBgpNetworkRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	RouterId             string `position:"Query" name:"RouterId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DstCidrBlock         string `position:"Query" name:"DstCidrBlock"`
}

func (r DeleteBgpNetworkRequest) Invoke(client *sdk.Client) (response *DeleteBgpNetworkResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DeleteBgpNetworkRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "DeleteBgpNetwork", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		DeleteBgpNetworkResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteBgpNetworkResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteBgpNetworkResponse struct {
	RequestId string
}
