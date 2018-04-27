package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
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
	req.InitWithApiInfo("Domain", "2018-01-29", "QueryTransferInList", "", "")
	resp = &QueryTransferInListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryTransferInListResponse struct {
	responses.BaseResponse
	RequestId      string
	TotalItemNum   int
	CurrentPageNum int
	TotalPageNum   int
	PageSize       int
	PrePage        bool
	NextPage       bool
	Data           QueryTransferInListTransferInInfoList
}

type QueryTransferInListTransferInInfo struct {
	SubmissionDate                              string
	ModificationDate                            string
	UserId                                      string
	InstanceId                                  string
	DomainName                                  string
	Status                                      int
	SimpleTransferInStatus                      string
	ResultCode                                  string
	ResultDate                                  string
	ResultMsg                                   string
	TransferAuthorizationCodeSubmissionDate     string
	NeedMailCheck                               bool
	Email                                       string
	WhoisMailStatus                             bool
	ExpirationDate                              string
	ProgressBarType                             int
	SubmissionDateLong                          int64
	ModificationDateLong                        int64
	ResultDateLong                              int64
	ExpirationDateLong                          int64
	TransferAuthorizationCodeSubmissionDateLong int64
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
