package ecs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AddBandwidthPackageIpsRequest struct {
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	BandwidthPackageId   string `position:"Query" name:"BandwidthPackageId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string `position:"Query" name:"ClientToken"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	IpCount              string `position:"Query" name:"IpCount"`
}

func (r AddBandwidthPackageIpsRequest) Invoke(client *sdk.Client) (response *AddBandwidthPackageIpsResponse, err error) {
	req := struct {
		*requests.RpcRequest
		AddBandwidthPackageIpsRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ecs", "2014-05-26", "AddBandwidthPackageIps", "ecs", "")

	resp := struct {
		*responses.BaseResponse
		AddBandwidthPackageIpsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.AddBandwidthPackageIpsResponse

	err = client.DoAction(&req, &resp)
	return
}

type AddBandwidthPackageIpsResponse struct {
	RequestId string
}
