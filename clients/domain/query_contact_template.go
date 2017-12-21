package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryContactTemplateRequest struct {
	requests.RpcRequest
	CCompany          string `position:"Query" name:"CCompany"`
	AuditStatus       string `position:"Query" name:"AuditStatus"`
	DefaultTemplate   string `position:"Query" name:"DefaultTemplate"`
	ECompany          string `position:"Query" name:"ECompany"`
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	PageSize          int    `position:"Query" name:"PageSize"`
	Lang              string `position:"Query" name:"Lang"`
	PageNum           int    `position:"Query" name:"PageNum"`
	ContactTemplateId int64  `position:"Query" name:"ContactTemplateId"`
	RegType           string `position:"Query" name:"RegType"`
}

func (req *QueryContactTemplateRequest) Invoke(client *sdk.Client) (resp *QueryContactTemplateResponse, err error) {
	req.InitWithApiInfo("Domain", "2016-05-11", "QueryContactTemplate", "", "")
	resp = &QueryContactTemplateResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryContactTemplateResponse struct {
	responses.BaseResponse
	RequestId        string
	TotalItemNum     int
	CurrentPageNum   int
	TotalPageNum     int
	PageSize         int
	PrePage          bool
	NextPage         bool
	ContactTemplates QueryContactTemplateContactTemplateList
}

type QueryContactTemplateContactTemplate struct {
	ContactTemplateId int64
	CreateTime        string
	UpdateTime        string
	UserId            string
	RegType           string
	DefaultTemplate   bool
	AuditStatus       string
	CName             string
	EName             string
	CCompany          string
	ECompany          string
	CCountry          string
	CProvince         string
	EProvince         string
	CCity             string
	ECity             string
	CVenu             string
	EVenu             string
	Email             string
	TelArea           string
	PostalCode        string
	TelMain           string
	TelExt            string
}

type QueryContactTemplateContactTemplateList []QueryContactTemplateContactTemplate

func (list *QueryContactTemplateContactTemplateList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryContactTemplateContactTemplate)
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
