package vpc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddBgpNetworkRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	RouterId             string `position:"Query" name:"RouterId"`
	VpcId                string `position:"Query" name:"VpcId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	DstCidrBlock         string `position:"Query" name:"DstCidrBlock"`
}

func (r AddBgpNetworkRequest) Invoke(client *sdk.Client) (response *AddBgpNetworkResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddBgpNetworkRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Vpc", "2016-04-28", "AddBgpNetwork", "vpc", "")

	resp := struct {
		*responses.BaseResponse
		AddBgpNetworkResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddBgpNetworkResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddBgpNetworkResponse struct {
	RequestId string
}
