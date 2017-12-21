package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type QueryFailReasonListRequest struct {
	SaleId            string `position:"Query" name:"SaleId"`
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	DomainName        string `position:"Query" name:"DomainName"`
	Lang              string `position:"Query" name:"Lang"`
	ContactTemplateId int64  `position:"Query" name:"ContactTemplateId"`
}

func (r QueryFailReasonListRequest) Invoke(client *sdk.Client) (response *QueryFailReasonListResponse, err error) {
	req := struct {
		*requests.RpcRequest
		QueryFailReasonListRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Domain", "2016-05-11", "QueryFailReasonList", "", "")

	resp := struct {
		*responses.BaseResponse
		QueryFailReasonListResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.QueryFailReasonListResponse

	err = client.DoAction(&req, &resp)
	return
}

type QueryFailReasonListResponse struct {
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
