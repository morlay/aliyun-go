package ubsms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SetUserBusinessStatusesRequest struct {
	requests.RpcRequest
	Uid           string `position:"Query" name:"Uid"`
	ServiceCode   string `position:"Query" name:"ServiceCode"`
	StatusKey1    string `position:"Query" name:"StatusKey.1"`
	StatusValue1  string `position:"Query" name:"StatusValue.1"`
	StatusKey2    string `position:"Query" name:"StatusKey.2"`
	StatusValue2  string `position:"Query" name:"StatusValue.2"`
	StatusKey3    string `position:"Query" name:"StatusKey.3"`
	StatusValue3  string `position:"Query" name:"StatusValue.3"`
	StatusKey4    string `position:"Query" name:"StatusKey.4"`
	StatusValue4  string `position:"Query" name:"StatusValue.4"`
	StatusKey5    string `position:"Query" name:"StatusKey.5"`
	StatusValue5  string `position:"Query" name:"StatusValue.5"`
	StatusKey6    string `position:"Query" name:"StatusKey.6"`
	StatusValue6  string `position:"Query" name:"StatusValue.6"`
	StatusKey7    string `position:"Query" name:"StatusKey.7"`
	StatusValue7  string `position:"Query" name:"StatusValue.7"`
	StatusKey8    string `position:"Query" name:"StatusKey.8"`
	StatusValue8  string `position:"Query" name:"StatusValue.8"`
	StatusKey9    string `position:"Query" name:"StatusKey.9"`
	StatusValue9  string `position:"Query" name:"StatusValue.9"`
	StatusKey10   string `position:"Query" name:"StatusKey.10"`
	StatusValue10 string `position:"Query" name:"StatusValue.10"`
	Password      string `position:"Query" name:"Password"`
}

func (req *SetUserBusinessStatusesRequest) Invoke(client *sdk.Client) (resp *SetUserBusinessStatusesResponse, err error) {
	req.InitWithApiInfo("Ubsms", "2015-06-23", "SetUserBusinessStatuses", "", "")
	resp = &SetUserBusinessStatusesResponse{}
	err = client.DoAction(req, resp)
	return
}

type SetUserBusinessStatusesResponse struct {
	responses.BaseResponse
	RequestId string
	Success   bool
}
