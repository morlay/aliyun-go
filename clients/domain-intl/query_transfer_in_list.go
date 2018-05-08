package domain_intl

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type QueryTransferInListRequest struct {
	requests.RpcRequest
	SubmissionStartDate    int64  `position:"Query" name:"SubmissionStartDate"`
	UserClientIp           string `position:"Query" name:"UserClientIp"`
	SubmissionEndDate      int64  `position:"Query" name:"SubmissionEndDate"`
	DomainName             string `position:"Query" name:"DomainName"`
	SimpleTransferInStatus string `position:"Query" name:"SimpleTransferInStatus"`
	PageSize               int    `position:"Query" name:"PageSize"`
	Lang                   string `position:"Query" name:"Lang"`
	PageNum                int    `position:"Query" name:"PageNum"`
}

func (req *QueryTransferInListRequest) Invoke(client *sdk.Client) (resp *QueryTransferInListResponse, err error) {
	req.InitWithApiInfo("Domain-intl", "2017-12-18", "QueryTransferInList", "domain", "")
	resp = &QueryTransferInListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryTransferInListResponse struct {
	responses.BaseResponse
	RequestId      common.String
	TotalItemNum   common.Integer
	CurrentPageNum common.Integer
	TotalPageNum   common.Integer
	PageSize       common.Integer
	PrePage        bool
	NextPage       bool
	Data           QueryTransferInListTransferInInfoList
}

type QueryTransferInListTransferInInfo struct {
	SubmissionDate                              common.String
	ModificationDate                            common.String
	UserId                                      common.String
	InstanceId                                  common.String
	DomainName                                  common.String
	Status                                      common.Integer
	SimpleTransferInStatus                      common.String
	ResultCode                                  common.String
	ResultDate                                  common.String
	ResultMsg                                   common.String
	TransferAuthorizationCodeSubmissionDate     common.String
	NeedMailCheck                               bool
	Email                                       common.String
	WhoisMailStatus                             bool
	ExpirationDate                              common.String
	ProgressBarType                             common.Integer
	SubmissionDateLong                          common.Long
	ModificationDateLong                        common.Long
	ResultDateLong                              common.Long
	ExpirationDateLong                          common.Long
	TransferAuthorizationCodeSubmissionDateLong common.Long
}

type QueryTransferInListTransferInInfoList []QueryTransferInListTransferInInfo

func (list *QueryTransferInListTransferInInfoList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryTransferInListTransferInInfo)
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
