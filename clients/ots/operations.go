package ots

func (c *OtsClient) GetInstance(req *GetInstanceArgs) (resp *GetInstanceResponse, err error) {
	resp = &GetInstanceResponse{}
	err = c.Request("GetInstance", req, resp)
	return
}

type GetInstanceArgs struct {
	Accept      string
	VersionName string
}
type GetInstanceResponse struct {
	RequestId string
}
