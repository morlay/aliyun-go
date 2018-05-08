package dcdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeDcdnDomainCnameRequest struct {
	requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDcdnDomainCnameRequest) Invoke(client *sdk.Client) (resp *DescribeDcdnDomainCnameResponse, err error) {
	req.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainCname", "dcdn", "")
	resp = &DescribeDcdnDomainCnameResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDcdnDomainCnameResponse struct {
	responses.BaseResponse
	RequestId  common.String
	CnameDatas DescribeDcdnDomainCnameDataList
}

type DescribeDcdnDomainCnameData struct {
	Domain common.String
	Cname  common.String
	Status common.Integer
}

type DescribeDcdnDomainCnameDataList []DescribeDcdnDomainCnameData

func (list *DescribeDcdnDomainCnameDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDcdnDomainCnameData)
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
