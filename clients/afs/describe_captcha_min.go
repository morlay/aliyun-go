package afs

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCaptchaMinRequest struct {
	requests.RpcRequest
	ResourceOwnerId int64  `position:"Query" name:"ResourceOwnerId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	ConfigName      string `position:"Query" name:"ConfigName"`
	Time            string `position:"Query" name:"Time"`
	Type            string `position:"Query" name:"Type"`
}

func (req *DescribeCaptchaMinRequest) Invoke(client *sdk.Client) (resp *DescribeCaptchaMinResponse, err error) {
	req.InitWithApiInfo("afs", "2018-01-12", "DescribeCaptchaMin", "", "")
	resp = &DescribeCaptchaMinResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCaptchaMinResponse struct {
	responses.BaseResponse
	RequestId   string
	BizCode     string
	HasData     bool
	CaptchaMins DescribeCaptchaMinCaptchaMinList
}

type DescribeCaptchaMinCaptchaMin struct {
	Time         string
	Pass         string
	Interception string
}

type DescribeCaptchaMinCaptchaMinList []DescribeCaptchaMinCaptchaMin

func (list *DescribeCaptchaMinCaptchaMinList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCaptchaMinCaptchaMin)
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	for _, v := range m {
		*list = v
		break
	}
	return nil
}
