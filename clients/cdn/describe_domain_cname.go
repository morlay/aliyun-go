package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeDomainCnameRequest struct {
	requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	OwnerId    int64  `position:"Query" name:"OwnerId"`
}

func (req *DescribeDomainCnameRequest) Invoke(client *sdk.Client) (resp *DescribeDomainCnameResponse, err error) {
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainCname", "", "")
	resp = &DescribeDomainCnameResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeDomainCnameResponse struct {
	responses.BaseResponse
	RequestId  string
	CnameDatas DescribeDomainCnameDataList
}

type DescribeDomainCnameData struct {
	Domain string
	Cname  string
	Status int
}

type DescribeDomainCnameDataList []DescribeDomainCnameData

func (list *DescribeDomainCnameDataList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeDomainCnameData)
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
