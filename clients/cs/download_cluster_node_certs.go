package cs

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DownloadClusterNodeCertsRequest struct {
	requests.RoaRequest
	Token  string `position:"Path" name:"Token"`
	NodeId string `position:"Path" name:"NodeId"`
}

func (req *DownloadClusterNodeCertsRequest) Invoke(client *sdk.Client) (resp *DownloadClusterNodeCertsResponse, err error) {
	req.InitWithApiInfo("CS", "2015-12-15", "DownloadClusterNodeCerts", "/token/[Token]/nodes/[NodeId]/certs", "", "")
	req.Method = "GET"

	resp = &DownloadClusterNodeCertsResponse{}
	err = client.DoAction(req, resp)
	return
}

type DownloadClusterNodeCertsResponse struct {
	responses.BaseResponse
}
