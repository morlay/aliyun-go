package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ModifyClusterServiceConfigRequest struct {
	requests.RpcRequest
	ResourceOwnerId    int64  `position:"Query" name:"ResourceOwnerId"`
	CustomConfigParams string `position:"Query" name:"CustomConfigParams"`
	GroupId            int64  `position:"Query" name:"GroupId"`
	ServiceName        string `position:"Query" name:"ServiceName"`
	Comment            string `position:"Query" name:"Comment"`
	ClusterId          string `position:"Query" name:"ClusterId"`
	ConfigParams       string `position:"Query" name:"ConfigParams"`
}

func (req *ModifyClusterServiceConfigRequest) Invoke(client *sdk.Client) (resp *ModifyClusterServiceConfigResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "ModifyClusterServiceConfig", "", "")
	resp = &ModifyClusterServiceConfigResponse{}
	err = client.DoAction(req, resp)
	return
}

type ModifyClusterServiceConfigResponse struct {
	responses.BaseResponse
	RequestId string
}
