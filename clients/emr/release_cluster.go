package emr

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ReleaseClusterRequest struct {
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	Id              string `position:"Query" name:"Id"`
	ForceRelease    string `position:"Query" name:"ForceRelease"`
}

func (r ReleaseClusterRequest) Invoke(client *sdk.Client) (response *ReleaseClusterResponse, err error) {
	req := struct {
		*requests.RpcRequest
		ReleaseClusterRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Emr", "2016-04-08", "ReleaseCluster", "", "")

	resp := struct {
		*responses.BaseResponse
		ReleaseClusterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ReleaseClusterResponse

	err = client.DoAction(&req, &resp)
	return
}

type ReleaseClusterResponse struct {
	RequestId string
}
