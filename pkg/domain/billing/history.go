package billing

import "time"

type History struct {
	id           string
	description  string
	user         string // *userDomain
	subscription string // enumType SUBSCRIPTION OR POSTPAID
	itemType     string
	status       string
	createdAt    *time.Time
	transaction  string // *transactionDomain
	item         string // *billingItemDomain
}

type Card struct {
	id           string
	batchID      string
	serialNumber string
	kind         string
	reason       string
	disabledAt   *time.Time
	linkedAt     *time.Time
	createdAt    *time.Time
	active       bool
	holder       string // *userDomain
	distributor  string // *userDomain
	customer     string // * userDomain
	zone         string // *zoneDomain
}

type Item struct {
	id                 string
	itemType           string // enumType SUBSCRIPTION OR POSTPAID
	subscription       string // *SubscriptionDomain
	plan               string // *PlanDOmain
	amount             float64
	billingPeriodStart *time.Time
	billingPeriodEnd   *time.Time
	liters             float64
	startConsumption   string // *DeviceMeterUpLinkDomain
	endConsumption     string // *DeviceMeterUpLinkDomain
	device             string // *DeviceDomain
	currency           string
}

type SubscriptionD struct {
	id                  string
	plan                string // *PlanDomain
	subscriber          string // *UserDomain
	name                string
	active              bool
	agent               string
	canceledImmediately string
	startsAt            *time.Time
	endsAt              *time.Time
	trialEndsAt         *time.Time
	canceledAt          *time.Time
	createdAt           *time.Time
	updatedAt           *time.Time
}

type Information struct {
	id           string
	userID       string
	firstName    string
	lastName     string
	addressLine1 string
	addressLine2 string
	city         string
	state        string
	postCode     string
	createdAt    string
	updatedAt    string
}

type PlanFeature struct {
	id     string
	planID string
	code   string
	value  string
	// sortOrder int
}

type Plan struct {
	id              string
	name            string
	description     string
	price           float64
	currency        string
	intervalCount   int64
	interval        string
	groupName       string
	trialPeriodDays string
	recommended     bool
	isActivated     bool
	activatedAt     int64
	features        []*PlanFeature
	// sortOrder int
}

type DeviceMeterUpLinkDomain struct {
	id                    string
	battery1Alarm         bool
	battery2Alarm         bool
	classEnabled          string
	eeAlarm               bool
	emptyTubeAlarm        bool
	flowRate              float64
	gatewayID             string
	overangeAlarm         bool
	pipeLeakageAlarm      bool
	pipeBurstAlarm        bool
	receivedAt            *time.Time
	rechargeBalance       float64
	rechargeTimes         int64
	reverseFlowAlarm      bool
	rssi                  string
	deviceEUI             string
	address               string
	totalConsumption      float64
	valveStatus           float64
	waterTemperature      float64
	waterTemperatureAlarm bool
}

type User struct {
	id           string
	firstName    string
	lastName     string
	userName     string
	phoneNumber  string
	accountType  string
	emailAddress string
	country      string
	medium       string
	orgID        string
}

type DeviceCoordinate struct {
	latitude  float64
	longitude float64
}

type Wallet struct {
	id          string
	name        string
	meta        string
	holder      string // *User
	balance     float64
	description string
	createdAt   *time.Time
}

type PlanSubscription struct {
	id                  string
	name                string
	agent               string
	canceledImmediately bool
	plan                string // *Plan
	trialEndsAt         *time.Time
	startsAt            *time.Time
	endsAt              *time.Time
	canceledAt          *time.Time
	createdAt           *time.Time
	updatedAt           *time.Time
}
