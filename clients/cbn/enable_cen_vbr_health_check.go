package cbn

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type EnableCenVbrHealthCheckRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	CenId                string `position:"Query" name:"CenId"`
	HealthCheckSourceIp  string `position:"Query" name:"HealthCheckSourceIp"`
	VbrInstanceOwnerId   int64  `position:"Query" name:"VbrInstanceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	VbrInstanceId        string `position:"Query" name:"VbrInstanceId"`
	HealthCheckTargetIp  string `position:"Query" name:"HealthCheckTargetIp"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
	VbrInstanceRegionId  string `position:"Query" name:"VbrInstanceRegionId"`
}

func (req *EnableCenVbrHealthCheckRequest) Invoke(client *sdk.Client) (resp *EnableCenVbrHealthCheckResponse, err error) {
	req.InitWithApiInfo("Cbn", "2017-09-12", "EnableCenVbrHealthCheck", "cbn", "")
	resp = &EnableCenVbrHealthCheckResponse{}
	err = client.DoAction(req, resp)
	return
}

type EnableCenVbrHealthCheckResponse struct {
	responses.BaseResponse
	RequestId string
}
