package ossadmin

func (c *OssadminClient) ReleaseOssInstance(req *ReleaseOssInstanceArgs) (resp *ReleaseOssInstanceResponse, err error) {
	resp = &ReleaseOssInstanceResponse{}
	err = c.Request("ReleaseOssInstance", req, resp)
	return
}

type ReleaseOssInstanceArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	Region               string
	OwnerAccount         string
}
type ReleaseOssInstanceResponse struct {
	RequestId string
}

func (c *OssadminClient) StopOssInstance(req *StopOssInstanceArgs) (resp *StopOssInstanceResponse, err error) {
	resp = &StopOssInstanceResponse{}
	err = c.Request("StopOssInstance", req, resp)
	return
}

type StopOssInstanceArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	Region               string
	OwnerAccount         string
}
type StopOssInstanceResponse struct {
	RequestId string
}

func (c *OssadminClient) CreateOssInstance(req *CreateOssInstanceArgs) (resp *CreateOssInstanceResponse, err error) {
	resp = &CreateOssInstanceResponse{}
	err = c.Request("CreateOssInstance", req, resp)
	return
}

type CreateOssInstanceArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	Region               string
	OwnerAccount         string
}
type CreateOssInstanceResponse struct {
	RequestId      string
	AliUid         int64
	InstanceId     string
	InstacneStatus string
	InstanceName   string
	StartTime      string
	EndTime        string
}

func (c *OssadminClient) RestartOssInstance(req *RestartOssInstanceArgs) (resp *RestartOssInstanceResponse, err error) {
	resp = &RestartOssInstanceResponse{}
	err = c.Request("RestartOssInstance", req, resp)
	return
}

type RestartOssInstanceArgs struct {
	OwnerId              int64
	ResourceOwnerAccount string
	ResourceOwnerId      int64
	Region               string
	OwnerAccount         string
}
type RestartOssInstanceResponse struct {
	RequestId string
}
