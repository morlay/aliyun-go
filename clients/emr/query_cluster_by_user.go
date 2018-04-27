package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryClusterByUserRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
}

func (req *QueryClusterByUserRequest) Invoke(client *sdk.Client) (resp *QueryClusterByUserResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "QueryClusterByUser", "", "")
	resp = &QueryClusterByUserResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryClusterByUserResponse struct {
	responses.BaseResponse
	RequestId string
	Exist     bool
}
