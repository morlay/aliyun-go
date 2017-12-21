package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ScaleClusterRequest struct {
	ClusterId string `position:"Path" name:"ClusterId"`
}

func (r ScaleClusterRequest) Invoke(client *sdk.Client) (response *ScaleClusterResponse, err error) {
	req := struct {
		*requests.RoaRequest
		ScaleClusterRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "ScaleCluster", "/clusters/[ClusterId]", "", "")
	req.Method = "PUT"

	resp := struct {
		*responses.BaseResponse
		ScaleClusterResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.ScaleClusterResponse

	err = client.DoAction(&req, &resp)
	return
}

type ScaleClusterResponse struct {
}
