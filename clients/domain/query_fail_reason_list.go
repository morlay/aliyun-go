package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryFailReasonListRequest struct {
	requests.RpcRequest
	SaleId            string `position:"Query" name:"SaleId"`
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	DomainName        string `position:"Query" name:"DomainName"`
	Lang              string `position:"Query" name:"Lang"`
	ContactTemplateId int64  `position:"Query" name:"ContactTemplateId"`
}

func (req *QueryFailReasonListRequest) Invoke(client *sdk.Client) (resp *QueryFailReasonListResponse, err error) {
	req.InitWithApiInfo("Domain", "2016-05-11", "QueryFailReasonList", "", "")
	resp = &QueryFailReasonListResponse{}
	err = client.DoAction(req, resp)
	return
}

type QueryFailReasonListResponse struct {
	responses.BaseResponse
	RequestId string
	Data      QueryFailReasonListFailRecordList
}

type QueryFailReasonListFailRecord struct {
	Date       string
	FailReason string
}

type QueryFailReasonListFailRecordList []QueryFailReasonListFailRecord

func (list *QueryFailReasonListFailRecordList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]QueryFailReasonListFailRecord)
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
