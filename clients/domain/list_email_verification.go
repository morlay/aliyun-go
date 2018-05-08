package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	req.InitWithApiInfo("Domain", "2018-01-29", "ListEmailVerification", "", "")
	resp = &ListEmailVerificationResponse{}
	err = client.DoAction(req, resp)
	return
}

type ListEmailVerificationResponse struct {
	responses.BaseResponse
	RequestId      common.String
	TotalItemNum   common.Integer
	CurrentPageNum common.Integer
	TotalPageNum   common.Integer
	PageSize       common.Integer
	PrePage        bool
	NextPage       bool
	Data           ListEmailVerificationEmailVerificationList
}

type ListEmailVerificationEmailVerification struct {
	GmtCreate           common.String
	GmtModified         common.String
	Email               common.String
	UserId              common.String
	EmailVerificationNo common.String
	TokenSendTime       common.String
	VerificationStatus  common.Integer
	VerificationTime    common.String
	SendIp              common.String
	ConfirmIp           common.String
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
