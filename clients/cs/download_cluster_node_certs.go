package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DownloadClusterNodeCertsRequest struct {
	Token  string `position:"Path" name:"Token"`
	NodeId string `position:"Path" name:"NodeId"`
}

func (r DownloadClusterNodeCertsRequest) Invoke(client *sdk.Client) (response *DownloadClusterNodeCertsResponse, err error) {
	req := struct {
		*requests.RoaRequest
		DownloadClusterNodeCertsRequest
	}{
		&requests.RoaRequest{},
		r,
	}
	req.InitWithApiInfo("CS", "2015-12-15", "DownloadClusterNodeCerts", "/token/[Token]/nodes/[NodeId]/certs", "", "")
	req.Method = "GET"

	resp := struct {
		*responses.BaseResponse
		DownloadClusterNodeCertsResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DownloadClusterNodeCertsResponse

	err = client.DoAction(&req, &resp)
	return
}

type DownloadClusterNodeCertsResponse struct {
}
