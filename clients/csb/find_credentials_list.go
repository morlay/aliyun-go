package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type FindCredentialsListRequest struct {
	requests.RpcRequest
	CsbId     int64  `position:"Query" name:"CsbId"`
	PageNum   int    `position:"Query" name:"PageNum"`
	GroupName string `position:"Query" name:"GroupName"`
}

func (req *FindCredentialsListRequest) Invoke(client *sdk.Client) (resp *FindCredentialsListResponse, err error) {
	req.InitWithApiInfo("CSB", "2017-11-18", "FindCredentialsList", "CSB", "")
	resp = &FindCredentialsListResponse{}
	err = client.DoAction(req, resp)
	return
}

type FindCredentialsListResponse struct {
	responses.BaseResponse
	Code      int
	Message   string
	RequestId string
	Data      FindCredentialsListData
}

type FindCredentialsListData struct {
	CurrentPage    int
	PageNumber     int
	CredentialList FindCredentialsListCredentialList
}

type FindCredentialsListCredential struct {
	GmtCreate         int64
	Id                int64
	Name              string
	OwnerAttr         string
	UserId            string
	CurrentCredential FindCredentialsListCurrentCredential
	NewCredential     FindCredentialsListNewCredential
}

type FindCredentialsListCurrentCredential struct {
	AccessKey string
	SecretKey string
}

type FindCredentialsListNewCredential struct {
	AccessKey string
	SecretKey string
}

type FindCredentialsListCredentialList []FindCredentialsListCredential

func (list *FindCredentialsListCredentialList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]FindCredentialsListCredential)
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
