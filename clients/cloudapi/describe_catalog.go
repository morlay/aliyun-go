package cloudapi

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/morlay/aliyun-go/common"
)

type DescribeCatalogRequest struct {
	requests.RpcRequest
	CatalogId string `position:"Query" name:"CatalogId"`
}

func (req *DescribeCatalogRequest) Invoke(client *sdk.Client) (resp *DescribeCatalogResponse, err error) {
	req.InitWithApiInfo("CloudAPI", "2016-07-14", "DescribeCatalog", "apigateway", "")
	resp = &DescribeCatalogResponse{}
	err = client.DoAction(req, resp)
	return
}

type DescribeCatalogResponse struct {
	responses.BaseResponse
	RequestId    common.String
	CatalogId    common.String
	CatalogName  common.String
	Description  common.String
	ParentId     common.String
	CreatedTime  common.String
	ModifiedTime common.String
	RegionId     common.String
	ApiIds       DescribeCatalogApiIdList
}

type DescribeCatalogApiIdList []common.String

func (list *DescribeCatalogApiIdList) UnmarshalJSON(data []byte) error {
	m := make(map[string][]common.String)
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
