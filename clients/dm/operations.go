package dm

import (
	"github.com/morlay/aliyun-go/core"
)

func (c *DmClient) CreateMailAddress(req *CreateMailAddressArgs) (resp *CreateMailAddressResponse, err error) {
	resp = &CreateMailAddressResponse{}
	err = c.Request("CreateMailAddress", req, resp)
	return
}

type CreateMailAddressArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	AccountName          string
	ReplyAddress         string
	Sendtype             string
}
type CreateMailAddressResponse struct {
	RequestId string
}

func (c *DmClient) SingleSendMail(req *SingleSendMailArgs) (resp *SingleSendMailResponse, err error) {
	resp = &SingleSendMailResponse{}
	err = c.Request("SingleSendMail", req, resp)
	return
}

type SingleSendMailArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	AccountName          string
	AddressType          int
	TagName              string
	ReplyToAddress       core.Bool
	ToAddress            string
	Subject              string
	HtmlBody             string
	TextBody             string
	FromAlias            string
	ReplyAddress         string
	ReplyAddressAlias    string
	ClickTrace           string
}
type SingleSendMailResponse struct {
	RequestId string
	EnvId     string
}

func (c *DmClient) BatchSendMail(req *BatchSendMailArgs) (resp *BatchSendMailResponse, err error) {
	resp = &BatchSendMailResponse{}
	err = c.Request("BatchSendMail", req, resp)
	return
}

type BatchSendMailArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	TemplateName         string
	AccountName          string
	ReceiversName        string
	AddressType          int
	TagName              string
	ReplyAddress         string
	ReplyAddressAlias    string
	ClickTrace           string
}
type BatchSendMailResponse struct {
	RequestId string
	EnvId     string
}
