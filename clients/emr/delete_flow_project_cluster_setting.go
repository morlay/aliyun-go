package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DeleteFlowProjectClusterSettingRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string `position:"Query" name:"ClusterId"`
	ProjectId       string `position:"Query" name:"ProjectId"`
}

func (req *DeleteFlowProjectClusterSettingRequest) Invoke(client *sdk.Client) (resp *DeleteFlowProjectClusterSettingResponse, err error) {
	req.InitWithApiInfo("Emr", "2016-04-08", "DeleteFlowProjectClusterSetting", "", "")
	resp = &DeleteFlowProjectClusterSettingResponse{}
	err = client.DoAction(req, resp)
	return
}

type DeleteFlowProjectClusterSettingResponse struct {
	responses.BaseResponse
	RequestId common.String
	Data      bool
}
