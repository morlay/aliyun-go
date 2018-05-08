package afs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type CreateConfigurationRequest struct {
	requests.RpcRequest
	ResourceOwnerId     int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp            string `position:"Query" name:"SourceIp"`
	ConfigurationName   string `position:"Query" name:"ConfigurationName"`
	MaxPV               string `position:"Query" name:"MaxPV"`
	ConfigurationMethod string `position:"Query" name:"ConfigurationMethod"`
	ApplyType           string `position:"Query" name:"ApplyType"`
	Scene               string `position:"Query" name:"Scene"`
}

func (req *CreateConfigurationRequest) Invoke(client *sdk.Client) (resp *CreateConfigurationResponse, err error) {
	req.InitWithApiInfo("afs", "2018-01-12", "CreateConfiguration", "", "")
	resp = &CreateConfigurationResponse{}
	err = client.DoAction(req, resp)
	return
}

type CreateConfigurationResponse struct {
	responses.BaseResponse
	RequestId common.String
	BizCode   common.String
}
