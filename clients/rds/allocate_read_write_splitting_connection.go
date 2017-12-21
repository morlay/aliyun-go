package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AllocateReadWriteSplittingConnectionRequest struct {
	requests.RpcRequest
	ResourceOwnerId        int64  `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount   string `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount           string `position:"Query" name:"OwnerAccount"`
	Weight                 string `position:"Query" name:"Weight"`
	OwnerId                int64  `position:"Query" name:"OwnerId"`
	IPType                 string `position:"Query" name:"IPType"`
	Port                   string `position:"Query" name:"Port"`
	DistributionType       string `position:"Query" name:"DistributionType"`
	DBInstanceId           string `position:"Query" name:"DBInstanceId"`
	MaxDelayTime           string `position:"Query" name:"MaxDelayTime"`
}

func (req *AllocateReadWriteSplittingConnectionRequest) Invoke(client *sdk.Client) (resp *AllocateReadWriteSplittingConnectionResponse, err error) {
	req.InitWithApiInfo("Rds", "2014-08-15", "AllocateReadWriteSplittingConnection", "rds", "")
	resp = &AllocateReadWriteSplittingConnectionResponse{}
	err = client.DoAction(req, resp)
	return
}

type AllocateReadWriteSplittingConnectionResponse struct {
	responses.BaseResponse
	RequestId string
}
