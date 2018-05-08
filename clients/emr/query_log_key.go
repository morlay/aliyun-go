package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryLogKeyRequest struct {
	requests.RpcRequest
	JobId           string `position:"Query" name:"JobId"`
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	InstanceId      string `position:"Query" name:"InstanceId"`
	LogName         string `position:"Query" name:"LogName"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	KeyBase         string `position:"Query" name:"KeyBase"`
	ContainerId     string `position:"Query" name:"ContainerId"`
}

func (req *QueryLogKeyRequest) Invoke(client *sdk.Client) (resp *QueryLogKeyResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "QueryLogKey", "", "")
	resp = &QueryLogKeyResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryLogKeyResponse struct {
	responses.BaseResponse
	RequestId common.String
	LogKey    common.String
}
