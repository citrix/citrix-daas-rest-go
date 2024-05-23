package apis

type Service struct {
	client *APIClient
}

type APIClient struct {
	common Service // Reuse a single struct instead of allocating one for each service on the heap.

	// ComputerName, AdUserName, and AdPassword are the credentials for the StoreFront server.
	computerName string
	adUserName   string
	adPassword   string

	// API Services for StoreFront
	DeploymentSF            *STFDeployment
	StoreSF                 *STFStore
	AuthenticationServiceSF *STFAuthentication
	WebReceiverSF           *STFWebReceiver
	RoamingSF               *STFRoaming
}

func (o *APIClient) GetComputerName() string {
	return o.computerName
}

func (o *APIClient) SetComputerName(v string) {
	o.computerName = v
}

func (o *APIClient) GetAdUserName() string {
	return o.adUserName
}

func (o *APIClient) SetAdUserName(v string) {
	o.adUserName = v
}

func (o *APIClient) GetAdPassword() string {
	return o.adPassword
}

func (o *APIClient) SetAdPassword(v string) {
	o.adPassword = v
}

func NewAPIClient() *APIClient {
	c := &APIClient{}
	c.common.client = c

	c.DeploymentSF = (*STFDeployment)(&c.common)
	c.StoreSF = (*STFStore)(&c.common)
	c.AuthenticationServiceSF = (*STFAuthentication)(&c.common)
	c.WebReceiverSF = (*STFWebReceiver)(&c.common)
	c.RoamingSF = (*STFRoaming)(&c.common)

	return c
}
