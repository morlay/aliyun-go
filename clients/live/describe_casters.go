package live

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type DescribeCastersRequest struct {
	SecurityToken string `position:"Query" name:"SecurityToken"`
	CasterName    string `position:"Query" name:"CasterName"`
	CasterId      string `position:"Query" name:"CasterId"`
	PageSize      int    `position:"Query" name:"PageSize"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       int64  `position:"Query" name:"OwnerId"`
	PageNum       int    `position:"Query" name:"PageNum"`
	Version       string `position:"Query" name:"Version"`
	Status        int    `position:"Query" name:"Status"`
}

func (r DescribeCastersRequest) Invoke(client *sdk.Client) (response *DescribeCastersResponse, err error) {
	req := struct {
		*requests.RpcRequest
		DescribeCastersRequest
	}{
		&requests.RpcRequest{},
		r,
	}
	req.InitWithApiInfo("live", "2016-11-01", "DescribeCasters", "", "")

	resp := struct {
		*responses.BaseResponse
		DescribeCastersResponse
	}{
		BaseResponse: &responses.BaseResponse{},
	}
	response = &resp.DescribeCastersResponse

	err = client.DoAction(&req, &resp)
	return
}

type DescribeCastersResponse struct {
	RequestId  string
	Total      int
	CasterList DescribeCastersCasterList
}

type DescribeCastersCaster struct {
	Status         int
	NormType       int
	CasterId       string
	CasterName     string
	CreateTime     string
	Period         int
	ChargeType     string
	CasterTemplate string
}

type DescribeCastersCasterList []DescribeCastersCaster

func (list *DescribeCastersCasterList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]DescribeCastersCaster)
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
