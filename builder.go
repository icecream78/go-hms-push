package hms

// HuaweiClientBuilder creates a instance of the huawei cloud common client
// It's contained in huawei cloud app and provides service through huawei cloud app
type HuaweiClientBuilder struct {
	appId     string
	appSecret string

	transport Transporter
}

// NewHuaweiClient func for initialisation of client builder
func NewHuaweiClient() *HuaweiClientBuilder {
	return &HuaweiClientBuilder{}
}

func (b *HuaweiClientBuilder) SetAppId(appId string) *HuaweiClientBuilder {
	b.appId = appId
	return b
}

func (b *HuaweiClientBuilder) SetSecret(secret string) *HuaweiClientBuilder {
	b.appSecret = secret
	return b
}

func (b *HuaweiClientBuilder) SetTransport(transport Transporter) *HuaweiClientBuilder {
	b.transport = transport
	return b
}

func (b *HuaweiClientBuilder) Build() (*HuaweiClient, error) {
	if b.appId == "" {
		return nil, ErrorAppIdEmpty
	}

	if b.appSecret == "" {
		return nil, ErrorSecretEmpty
	}

	transport := b.transport
	if transport == nil {
		transport, _ = NewHTTPTransport(DefaultRetryCount, DefaultRetryIntervalMs)
	}

	client := &HuaweiClient{
		appId:     b.appId,
		appSecret: b.appSecret,
		transport: transport,
	}

	return client, nil
}
