package green

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SearchRequest struct {
	requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

func (req *SearchRequest) Invoke(client *sdk.Client) (resp *SearchResponse, err error) {
	req.InitWithApiInfo("Green", "2017-08-25", "Search", "/green/sface/search", "green", "")
	req.Method = "POST"

	resp = &SearchResponse{}
	err = client.DoAction(req, resp)
	return
}

type SearchResponse struct {
	responses.BaseResponse
}
