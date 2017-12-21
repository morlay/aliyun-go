package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ResizeClusterRequest struct {
	ClusterId          string `position:"Query" name:"ClusterId"`
	NewMasterInstances int    `position:"Query" name:"NewMasterInstances"`
	NewCoreInstances   int    `position:"Query" name:"NewCoreInstances"`
	NewTaskInstances   int    `position:"Query" name:"NewTaskInstances"`
	ChargeType         string `position:"Query" name:"ChargeType"`
	Period             int    `position:"Query" name:"Period"`
}

func (r ResizeClusterRequest) Invoke(client *sdk.Client) (response *ResizeClusterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ResizeClusterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "ResizeCluster", "", "")

	resp := struct {
		*responses.BaseResponse
		ResizeClusterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ResizeClusterResponse

	err = client.DoAction(&req, &resp)
	return
}

type ResizeClusterResponse struct {
	RequestId   string
	ClusterId   string
	EmrOrderId  string
	CoreOrderId string
}
