package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest struct {
	requests.RpcRequest
	ContactType           string                                                                  `position:"Query" name:"ContactType"`
	UserClientIp          string                                                                  `position:"Query" name:"UserClientIp"`
	RegistrantProfileId   int64                                                                   `position:"Query" name:"RegistrantProfileId"`
	DomainNames           *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	TransferOutProhibited string                                                                  `position:"Query" name:"TransferOutProhibited"`
	Lang                  string                                                                  `position:"Query" name:"Lang"`
}

func (req *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdRequest) Invoke(client *sdk.Client) (resp *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveBatchTaskForUpdatingContactInfoByRegistrantProfileId", "", "")
	resp = &SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponse struct {
	responses.BaseResponse
	RequestId common.String
	TaskNo    common.String
}

type SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdDomainNameList []string

func (list *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdDomainNameList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]string)
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
