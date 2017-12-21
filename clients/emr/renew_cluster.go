package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type RenewClusterRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ECSIds          string `position:"Query" name:"ECSIds"`
	Period          string `position:"Query" name:"Period"`
}

func (req *RenewClusterRequest) Invoke(client *sdk.Client) (resp *RenewClusterResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "RenewCluster", "", "")
	resp = &RenewClusterResponse{}
	err = client.DoAction(req, resp)
	return
}

type RenewClusterResponse struct {
	responses.BaseResponse
	RequestId string
	ClusterId string
}
