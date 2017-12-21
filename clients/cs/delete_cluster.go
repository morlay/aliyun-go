package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DeleteClusterRequest struct {
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (r DeleteClusterRequest) Invoke(client *sdk.Client) (response *DeleteClusterResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DeleteClusterRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DeleteCluster", "/clusters/[ClusterId]", "", "")
	req.Method = "DELETE"

	resp := struct {
		*responses.BaseResponse
		DeleteClusterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DeleteClusterResponse

	err = client.DoAction(&req, &resp)
	return
}

type DeleteClusterResponse struct {
}
