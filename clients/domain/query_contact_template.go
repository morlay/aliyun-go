package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryContactTemplateRequest struct {
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

func (r QueryContactTemplateRequest) Invoke(client *sdk.Client) (response *QueryContactTemplateResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryContactTemplateRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Domain", "2016-05-11", "QueryContactTemplate", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryContactTemplateResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryContactTemplateResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryContactTemplateResponse struct {
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
