package cbn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type AttachCenChildInstanceRequest struct {
	requests.RpcRequest
	ChildInstanceId       string `position:"Query" name:"ChildInstanceId"`
	ResourceOwnerId       int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount  string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                 string `position:"Query" name:"CenId"`
	OwnerAccount          string `position:"Query" name:"OwnerAccount"`
	OwnerId               int64  `position:"Query" name:"OwnerId"`
	ChildInstanceType     string `position:"Query" name:"ChildInstanceType"`
	ChildInstanceOwnerId  int64  `position:"Query" name:"ChildInstanceOwnerId"`
	ChildInstanceRegionId string `position:"Query" name:"ChildInstanceRegionId"`
}

func (req *AttachCenChildInstanceRequest) Invoke(client *sdk.Client) (resp *AttachCenChildInstanceResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "AttachCenChildInstance", "cbn", "")
	resp = &AttachCenChildInstanceResponse{}
	err = client.DoAction(req, resp)
	return
}

type AttachCenChildInstanceResponse struct {
	responses.BaseResponse
	RequestId common.String
}
