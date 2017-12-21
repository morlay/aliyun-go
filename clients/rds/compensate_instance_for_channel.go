package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type CompensateInstanceForChannelRequest struct {
	requests.RpcRequest
	ResourceOwnerId      int64  `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	ZoneId               string `position:"Query" name:"ZoneId"`
	SubDomain            string `position:"Query" name:"SubDomain"`
	DBInstanceId         string `position:"Query" name:"DBInstanceId"`
	OwnerId              int64  `position:"Query" name:"OwnerId"`
}

func (req *CompensateInstanceForChannelRequest) Invoke(client *sdk.Client) (resp *CompensateInstanceForChannelResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "CompensateInstanceForChannel", "rds", "")
	resp = &CompensateInstanceForChannelResponse{}
	err = client.DoAction(req, resp)
	return
}

type CompensateInstanceForChannelResponse struct {
	responses.BaseResponse
	RequestId string
}
