package csb

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
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
	Code      common.Integer
	Message   common.String
	RequestId common.String
	Data      FindCredentialsListData
}

type FindCredentialsListData struct {
	CurrentPage    common.Integer
	PageNumber     common.Integer
	CredentialList FindCredentialsListCredentialList
}

type FindCredentialsListCredential struct {
	GmtCreate         common.Long
	Id                common.Long
	Name              common.String
	OwnerAttr         common.String
	UserId            common.String
	CurrentCredential FindCredentialsListCurrentCredential
	NewCredential     FindCredentialsListNewCredential
}

type FindCredentialsListCurrentCredential struct {
	AccessKey common.String
	SecretKey common.String
}

type FindCredentialsListNewCredential struct {
	AccessKey common.String
	SecretKey common.String
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
