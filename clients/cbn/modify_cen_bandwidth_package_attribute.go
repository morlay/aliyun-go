package cbn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type ModifyCenBandwidthPackageAttributeRequest struct {
	requests.RpcRequest
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	CenBandwidthPackageId string `position:"Query" name:"CenBandwidthPackageId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	Name                  string `position:"Query" name:"Name"`
	Description           string `position:"Query" name:"Description"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
}

func (req *ModifyCenBandwidthPackageAttributeRequest) Invoke(client *sdk.Client) (resp *ModifyCenBandwidthPackageAttributeResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "ModifyCenBandwidthPackageAttribute", "cbn", "")
	resp = &ModifyCenBandwidthPackageAttributeResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyCenBandwidthPackageAttributeResponse struct {
	responses.BaseResponse
	RequestId common.String
}
