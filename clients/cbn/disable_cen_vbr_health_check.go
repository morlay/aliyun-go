package cbn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DisableCenVbrHealthCheckRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	VbrInstanceOwnerId   int64  `position:"Query" name:"VbrInstanceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VbrInstanceId        string `position:"Query" name:"VbrInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VbrInstanceRegionId  string `position:"Query" name:"VbrInstanceRegionId"`
}

func (req *DisableCenVbrHealthCheckRequest) Invoke(client *sdk.Client) (resp *DisableCenVbrHealthCheckResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "DisableCenVbrHealthCheck", "cbn", "")
	resp = &DisableCenVbrHealthCheckResponse{}
	err = client.DoAction(req, resp)
	return
}

type DisableCenVbrHealthCheckResponse struct {
	responses.BaseResponse
	RequestId string
}
