package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyClusterNameRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Name            string `position:"Query" name:"Name"`
	Id              string `position:"Query" name:"Id"`
}

func (req *ModifyClusterNameRequest) Invoke(client *sdk.Client) (resp *ModifyClusterNameResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyClusterName", "", "")
	resp = &ModifyClusterNameResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyClusterNameResponse struct {
	responses.BaseResponse
	RequestId string
}
