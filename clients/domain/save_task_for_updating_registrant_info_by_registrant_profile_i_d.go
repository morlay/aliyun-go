package domain

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest struct {
	requests.RpcRequest
	UserClientIp          string                                                                `position:"Query" name:"UserClientIp"`
	RegistrantProfileId   int64                                                                 `position:"Query" name:"RegistrantProfileId"`
	DomainNames           *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDDomainNameList `position:"Query" type:"Repeated" name:"DomainName"`
	TransferOutProhibited string                                                                `position:"Query" name:"TransferOutProhibited"`
	Lang                  string                                                                `position:"Query" name:"Lang"`
}

func (req *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDRequest) Invoke(client *sdk.Client) (resp *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse, err error) {
	req.InitWithApiInfo("Domain", "2018-01-29", "SaveTaskForUpdatingRegistrantInfoByRegistrantProfileID", "", "")
	resp = &SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse{}
	err = client.DoAction(req, resp)
	return
}

type SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponse struct {
	responses.BaseResponse
	RequestId string
	TaskNo    string
}

type SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDDomainNameList []string

func (list *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDDomainNameList) UnmarshalJSON(data []byte) error {
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
