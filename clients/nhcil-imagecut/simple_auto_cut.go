package nhcil_imagecut

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SimpleAutoCutRequest struct {
	requests.RoaRequest
	ImageUrl string `position:"Query" name:"ImageUrl"`
}

func (req *SimpleAutoCutRequest) Invoke(client *sdk.Client) (resp *SimpleAutoCutResponse, err error) {
	req.InitWithApiInfo("Nhcil-imagecut", "2018-05-20", "SimpleAutoCut", "/simpleautocut", "", "")
	req.Method = "GET"

	resp = &SimpleAutoCutResponse{}
	err = client.DoAction(req, resp)
	return
}

type SimpleAutoCutResponse struct {
	responses.BaseResponse
	Code      common.Integer
	Data      common.String
	RequestId common.String
	Success   bool
	Errmsg    common.String
}
