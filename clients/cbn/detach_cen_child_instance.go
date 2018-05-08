package cbn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DetachCenChildInstanceRequest struct {
	requests.RpcRequest
	ChildInstanceId       string `position:"Query" name:"ChildInstanceId"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                 string `position:"Query" name:"CenId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	CenOwnerId            int64  `position:"Query" name:"CenOwnerId"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	ChildInstanceType     string `position:"Query" name:"ChildInstanceType"`
	ChildInstanceOwnerId  int64  `position:"Query" name:"ChildInstanceOwnerId"`
	ChildInstanceRegionId string `position:"Query" name:"ChildInstanceRegionId"`
}

func (req *DetachCenChildInstanceRequest) Invoke(client *sdk.Client) (resp *DetachCenChildInstanceResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "DetachCenChildInstance", "cbn", "")
	resp = &DetachCenChildInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type DetachCenChildInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
