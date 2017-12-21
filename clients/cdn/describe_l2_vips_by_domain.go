package cdn

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeL2VipsByDomainRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
}

func (r DescribeL2VipsByDomainRequest) Invoke(client *sdk.Client) (response *DescribeL2VipsByDomainResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeL2VipsByDomainRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("Cdn", "2014-11-11", "DescribeL2VipsByDomain", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeL2VipsByDomainResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeL2VipsByDomainResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeL2VipsByDomainResponse struct {
	RequestId  string
	DomainName string
	Vips       DescribeL2VipsByDomainVipList
}

type DescribeL2VipsByDomainVipList []string

func (list *DescribeL2VipsByDomainVipList) UnmarshalJSON(data []byte) error {
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
