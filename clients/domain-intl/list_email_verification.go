package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type ListEmailVerificationRequest struct {
	requests.RpcRequest
	BeginCreateTime    int64  `position:"Query" name:"BeginCreateTime"`
	EndCreateTime      int64  `position:"Query" name:"EndCreateTime"`
	PageSize           int    `position:"Query" name:"PageSize"`
	Lang               string `position:"Query" name:"Lang"`
	PageNum            int    `position:"Query" name:"PageNum"`
	Email              string `position:"Query" name:"Email"`
	VerificationStatus int    `position:"Query" name:"VerificationStatus"`
}

func (req *ListEmailVerificationRequest) Invoke(client *sdk.Client) (resp *ListEmailVerificationResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "ListEmailVerification", "domain", "")
	resp = &ListEmailVerificationResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListEmailVerificationResponse struct {
	responses.BaseResponse
	RequestId      string
	TotalItemNum   int
	CurrentPageNum int
	TotalPageNum   int
	PageSize       int
	PrePage        bool
	NextPage       bool
	Data           ListEmailVerificationEmailVerificationList
}

type ListEmailVerificationEmailVerification struct {
	GmtCreate           string
	GmtModified         string
	Email               string
	UserId              string
	EmailVerificationNo string
	TokenSendTime       string
	VerificationStatus  int
	VerificationTime    string
	SendIp              string
	ConfirmIp           string
}

type ListEmailVerificationEmailVerificationList []ListEmailVerificationEmailVerification

func (list *ListEmailVerificationEmailVerificationList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]ListEmailVerificationEmailVerification)
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
