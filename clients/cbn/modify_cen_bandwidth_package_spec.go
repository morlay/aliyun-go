package cbn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyCenBandwidthPackageSpecRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	Bandwidth             int    `position:"Query" name:"Bandwidth"`
	CenBandwidthPackageId string `position:"Query" name:"CenBandwidthPackageId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyCenBandwidthPackageSpecRequest) Invoke(client *sdk.Client) (resp *ModifyCenBandwidthPackageSpecResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "ModifyCenBandwidthPackageSpec", "cbn", "")
	resp = &ModifyCenBandwidthPackageSpecResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCenBandwidthPackageSpecResponse struct {
	responses.BaseResponse
	RequestId common.String
}
