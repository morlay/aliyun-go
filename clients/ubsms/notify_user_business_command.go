package ubsms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type NotifyUserBusinessCommandRequest struct {
	Uid         string `position:"Query" name:"Uid"`
	ServiceCode string `position:"Query" name:"ServiceCode"`
	Cmd         string `position:"Query" name:"Cmd"`
	Region      string `position:"Query" name:"Region"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	ClientToken string `position:"Query" name:"ClientToken"`
	Password    string `position:"Query" name:"Password"`
}

func (r NotifyUserBusinessCommandRequest) Invoke(client *sdk.Client) (response *NotifyUserBusinessCommandResponse, err error) {
	req := struct {
		*requests.RpcRequest
		NotifyUserBusinessCommandRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Ubsms", "2015-06-23", "NotifyUserBusinessCommand", "", "")

	resp := struct {
		*responses.BaseResponse
		NotifyUserBusinessCommandResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.NotifyUserBusinessCommandResponse

	err = client.DoAction(&req, &resp)
	return
}

type NotifyUserBusinessCommandResponse struct {
	RequestId string
	Success   bool
}
