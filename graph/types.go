package graph

import "time"

// UserGQL is the GraphQL-facing shape (ID as string hex).
type UserGQL struct {
	ID    string
	Name  string
	Email string
}

func toGQL(u *User) *UserGQL {
	return &UserGQL{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

type DeviceInstanceTagGQL struct {
	SensorID        string `json:"sensorId"`
	Source          string `json:"source"`
	TagID           string `json:"tagId"`
	VirtualDeviceID string `json:"virtualDeviceId"`
}

type DeviceInstanceBelongsRefGQL struct {
	ParentID string `json:"parentId"`
	NodeID   string `json:"nodeId"`
}

type DeviceInstanceAssignRefGQL struct {
	ParentID string `json:"parentId"`
	NodeID   string `json:"nodeId"`
}

type DeviceInstanceGeneralInfoGQL struct {
	PowerOutage      *bool                          `json:"powerOutage"`
	BelongsRef       []*DeviceInstanceBelongsRefGQL `json:"belongsRef"`
	AssignRef        []*DeviceInstanceAssignRefGQL  `json:"assignRef"`
	Timeout          *string                        `json:"timeout"`
	DeviceComID      *string                        `json:"deviceComId"`
	Make             *string                        `json:"make"`
	ModelNumber      *string                        `json:"modelNumber"`
	ProtocolCategory []string                       `json:"protocolCategory"`
	DeviceModelName  *string                        `json:"deviceModelName"`
	DeviceModelRefID *string                        `json:"deviceModelRefId"`
	DeviceName       *string                        `json:"deviceName"`
	DeviceSelection  []string                       `json:"deviceSelection"`
	IsDeleted        *bool                          `json:"isDeleted"`
	IsDisabled       *bool                          `json:"isDisabled"`
	MacID            *string                        `json:"macId"`
}

type DeviceInstanceGQL struct {
	ID                string
	DeviceInstanceID  string
	SiteID            string
	GatewayID         string
	DeviceName        string
	ClientID          *string
	Default           *bool
	GatewayInstanceID *string
	EventUserID       *string
	GeneralInfo       *DeviceInstanceGeneralInfoGQL
	TagsData          []*DeviceInstanceTagGQL
}

func toDeviceInstanceGQL(d *DeviceInstance) *DeviceInstanceGQL {
	var tags []*DeviceInstanceTagGQL
	for _, t := range d.TagsData {
		tags = append(tags, &DeviceInstanceTagGQL{
			SensorID:        t.SensorID,
			Source:          t.Source,
			TagID:           t.TagID,
			VirtualDeviceID: t.VirtualDeviceID,
		})
	}
	var belongs []*DeviceInstanceBelongsRefGQL
	for _, b := range d.GeneralInfo.BelongsRef {
		belongs = append(belongs, &DeviceInstanceBelongsRefGQL{
			ParentID: b.ParentID,
			NodeID:   b.NodeID,
		})
	}
	var assigns []*DeviceInstanceAssignRefGQL
	for _, a := range d.GeneralInfo.AssignRef {
		assigns = append(assigns, &DeviceInstanceAssignRefGQL{
			ParentID: a.ParentID,
			NodeID:   a.NodeID,
		})
	}

	isDefault := bool(d.Default)
	powerOutage := bool(d.GeneralInfo.PowerOutage)
	isDeleted := bool(d.GeneralInfo.IsDeleted)
	isDisabled := bool(d.GeneralInfo.IsDisabled)
	timeout := string(d.GeneralInfo.Timeout)
	deviceComID := string(d.GeneralInfo.DeviceComID)

	return &DeviceInstanceGQL{
		ID:                d.ID,
		DeviceInstanceID:  d.DeviceInstanceID,
		SiteID:            d.SiteID,
		GatewayID:         d.GatewayID,
		DeviceName:        d.GeneralInfo.DeviceName,
		ClientID:          &d.ClientID,
		Default:           &isDefault,
		GatewayInstanceID: &d.GatewayInstanceID,
		EventUserID:       &d.EventUserID,
		GeneralInfo: &DeviceInstanceGeneralInfoGQL{
			PowerOutage:      &powerOutage,
			BelongsRef:       belongs,
			AssignRef:        assigns,
			Timeout:          &timeout,
			DeviceComID:      &deviceComID,
			Make:             &d.GeneralInfo.Make,
			ModelNumber:      &d.GeneralInfo.ModelNumber,
			ProtocolCategory: []string(d.GeneralInfo.ProtocolCategory),
			DeviceModelName:  &d.GeneralInfo.DeviceModelName,
			DeviceModelRefID: &d.GeneralInfo.DeviceModelRefID,
			DeviceName:       &d.GeneralInfo.DeviceName,
			DeviceSelection:  []string(d.GeneralInfo.DeviceSelection),
			IsDeleted:        &isDeleted,
			IsDisabled:       &isDisabled,
			MacID:            &d.GeneralInfo.MacID,
		},
		TagsData: tags,
	}
}

type DeviceModelGQL struct {
	ID              string
	DeviceModelID   string
	DeviceModelName string
	Make            string
	ModelNumber     string
}

func toDeviceModelGQL(m *DeviceModel) *DeviceModelGQL {
	return &DeviceModelGQL{
		ID:              m.ID,
		DeviceModelID:   m.DeviceModelID,
		DeviceModelName: m.GeneralInfo.DeviceModelName,
		Make:            m.GeneralInfo.Make,
		ModelNumber:     m.GeneralInfo.ModelNumber,
	}
}

type SensorGQL struct {
	ID               string
	DeviceInstanceID string
	SensorID         string
	SiteID           string
	DeviceName       string
	Make             string
}

func toSensorGQL(s *Sensor) *SensorGQL {
	return &SensorGQL{
		ID:               s.ID,
		DeviceInstanceID: s.DeviceInstanceID,
		SensorID:         s.ID,
		SiteID:           s.SiteID,
		DeviceName:       s.GeneralInfo.DeviceName,
		Make:             s.GeneralInfo.Make,
	}
}

type BatchGQL struct {
	ID               string
	BatchID          string
	BatchName        string
	BatchDescription string
	SiteID           string
	ClientID         string
}

func toBatchGQL(b *Batch) *BatchGQL {
	return &BatchGQL{
		ID:               b.ID,
		BatchID:          b.BatchID,
		BatchName:        b.BatchName,
		BatchDescription: b.BatchDescription,
		SiteID:           b.SiteID,
		ClientID:         b.ClientID,
	}
}

type ClientGQL struct {
	ID                string
	ClientID          string
	ClientName        string
	ClientDescription string
	ClientAddress     string
	ClientLogo        string
	UserID            string
}

func toClientGQL(c *Client) *ClientGQL {
	return &ClientGQL{
		ID:                c.ID,
		ClientID:          c.ClientID,
		ClientName:        c.ClientName,
		ClientDescription: c.ClientDescription,
		ClientAddress:     c.ClientAddress,
		ClientLogo:        c.ClientLogo,
		UserID:            c.UserID,
	}
}

type CostGQL struct {
	ID                  string
	CostManagementID    string
	CostType            int
	DgBaseCost          float64
	DgCost              float64
	EbBaseCost          float64
	EbCost              float64
	SiteID              string
	ClientID            string
	CostManagementKey   *string
	CostManagementKeyID *string
	FromDate            *string
	FromDateStr         *string
	ToDate              *string
	ToDateStr           *string
	RechargeTax         *float64
	ServiceProvider     *string
	ServiceProviderID   *string
	TariffCateg         *string
	TariffCategoryID    *string
	CgstTax             *float64
	SgstTax             *float64
	CamCost             *float64
	Ed1phase            *float64
	Ed3phase            *float64
}

func toCostGQL(c *Cost) *CostGQL {
	var fromDateStr *string
	if !c.FromDate.IsZero() {
		s := c.FromDate.Format(time.RFC3339)
		fromDateStr = &s
	}
	var toDateStr *string
	if !c.ToDate.IsZero() {
		s := c.ToDate.Format(time.RFC3339)
		toDateStr = &s
	}

	dgBaseCost := float64(c.DgBaseCost)
	dgCost := float64(c.DgCost)
	ebBaseCost := float64(c.EbBaseCost)
	ebCost := float64(c.EbCost)
	rechargeTax := float64(c.RechargeTax)
	cgstTax := float64(c.CgstTax)
	sgstTax := float64(c.SgstTax)
	camCost := float64(c.CamCost)
	ed1phase := float64(c.Ed1phase)
	ed3phase := float64(c.Ed3phase)

	return &CostGQL{
		ID:                  c.ID,
		CostManagementID:    c.CostManagementID,
		CostType:            int(c.CostType),
		DgBaseCost:          dgBaseCost,
		DgCost:              dgCost,
		EbBaseCost:          ebBaseCost,
		EbCost:              ebCost,
		SiteID:              c.SiteID,
		ClientID:            c.ClientID,
		CostManagementKey:   &c.CostManagementKey,
		CostManagementKeyID: &c.CostManagementKeyID,
		FromDate:            fromDateStr,
		FromDateStr:         &c.FromDateStr,
		ToDate:              toDateStr,
		ToDateStr:           &c.ToDateStr,
		RechargeTax:         &rechargeTax,
		ServiceProvider:     &c.ServiceProvider,
		ServiceProviderID:   &c.ServiceProviderID,
		TariffCateg:         &c.TariffCateg,
		TariffCategoryID:    &c.TariffCategoryID,
		CgstTax:             &cgstTax,
		SgstTax:             &sgstTax,
		CamCost:             &camCost,
		Ed1phase:            &ed1phase,
		Ed3phase:            &ed3phase,
	}
}

type SiteGQL struct {
	ID       string
	SiteID   string
	SiteName string
	ClientID string
}

func toSiteGQL(s *Site) *SiteGQL {
	siteID := s.SiteID
	if siteID == "" {
		if s.IndustryID != "" {
			siteID = s.IndustryID
		} else {
			siteID = s.ID
		}
	}
	siteName := s.SiteName
	if siteName == "" {
		siteName = s.IndustryName
	}
	return &SiteGQL{
		ID:       s.ID,
		SiteID:   siteID,
		SiteName: siteName,
		ClientID: s.ClientID,
	}
}

type GatewayGQL struct {
	ID                string
	GatewayID         string
	GatewayName       string
	SiteID            string
	DataReadFrequency *float64
	GatewayModelRefID *string
	GatewayModelName  *string
	Type              *string
	Default           *bool
	IsDeleted         *string
}

func toGatewayGQL(g *Gateway) *GatewayGQL {
	dataReadFrequency := float64(g.DataReadFrequency)
	isDefault := bool(g.Default)
	return &GatewayGQL{
		ID:                g.ID,
		GatewayID:         g.GatewayID,
		GatewayName:       g.GatewayName,
		SiteID:            g.SiteID,
		DataReadFrequency: &dataReadFrequency,
		GatewayModelRefID: &g.GatewayModelRefID,
		GatewayModelName:  &g.GatewayModelName,
		Type:              &g.Type,
		Default:           &isDefault,
		IsDeleted:         &g.IsDeleted,
	}
}

type GatewayInstanceAssignRefGQL struct {
	NodeID   string `json:"nodeId"`
	ParentID string `json:"parentId"`
}

type GatewayInstanceGQL struct {
	ID                   string
	GatewayInstanceID    string
	GatewayName          string
	MacID                string
	SiteID               string
	ClientID             string
	AssignedIndustry     *string
	AwsLicense           *bool
	BaudRate             *string
	ComPort              *string
	DataBit              *string
	Parity               *string
	StopBit              *string
	Protocol             *string
	StatusTimeout        *float64
	Timeout              *float64
	Type                 *string
	Default              *bool
	IsConfigured         *bool
	IsDeleted            *bool
	IsDisabled           *bool
	DataReadFrequency    *float64
	DataLoggingFrequency *float64
	AssignRef            []*GatewayInstanceAssignRefGQL
}

func toGatewayInstanceGQL(gi *GatewayInstance) *GatewayInstanceGQL {
	var assignRefs []*GatewayInstanceAssignRefGQL
	for _, ar := range gi.AssignRef {
		assignRefs = append(assignRefs, &GatewayInstanceAssignRefGQL{
			NodeID:   ar.NodeID,
			ParentID: ar.ParentID,
		})
	}
	statusTimeout := float64(gi.StatusTimeout)
	timeout := float64(gi.Timeout)
	dataReadFrequency := float64(gi.DataReadFrequency)
	dataLoggingFrequency := float64(gi.DataLoggingFrequency)

	awsLicense := bool(gi.AwsLicense)
	isDefault := bool(gi.Default)
	isConfigured := bool(gi.IsConfigured)
	isDeleted := bool(gi.IsDeleted)
	isDisabled := bool(gi.IsDisabled)
	baudRate := string(gi.BaudRate)
	comPort := string(gi.ComPort)
	dataBit := string(gi.DataBit)
	parity := string(gi.Parity)
	stopBit := string(gi.StopBit)

	return &GatewayInstanceGQL{
		ID:                   gi.ID,
		GatewayInstanceID:    gi.GatewayInstanceID,
		GatewayName:          gi.GatewayName,
		MacID:                gi.MacID,
		SiteID:               gi.SiteID,
		ClientID:             gi.ClientID,
		AssignedIndustry:     &gi.AssignedIndustry,
		AwsLicense:           &awsLicense,
		BaudRate:             &baudRate,
		ComPort:              &comPort,
		DataBit:              &dataBit,
		Parity:               &parity,
		StopBit:              &stopBit,
		Protocol:             &gi.Protocol,
		StatusTimeout:        &statusTimeout,
		Timeout:              &timeout,
		Type:                 &gi.Type,
		Default:              &isDefault,
		IsConfigured:         &isConfigured,
		IsDeleted:            &isDeleted,
		IsDisabled:           &isDisabled,
		DataReadFrequency:    &dataReadFrequency,
		DataLoggingFrequency: &dataLoggingFrequency,
		AssignRef:            assignRefs,
	}
}

type AlarmNotificationProfileUserGQL struct {
	Label string
	Type  string
	Value string
}

type AlarmNotificationProfileGQL struct {
	UsersOrUserGroup       []*AlarmNotificationProfileUserGQL
	IsNotificationToneShow *bool
	NotificationProfile    []string
	NotificationTone       *string
}

type AlarmLevelGQL struct {
	Commands             []string
	NotificationProfiles []*AlarmNotificationProfileGQL
	Suppress             *string
}

type RuleSetAndOrOperationDataGQL struct {
	IsAnd bool
	IsOr  bool
}

type RuleLeftHandSideGQL struct {
	Tag string
}

type RuleRightHandSideGQL struct {
	CompareOption string
	CustomValue   *float64
	Tag           *string
	Threshold     *float64
}

type RuleResetValueGQL struct {
	CompareOption        string
	CustomValue          *float64
	IsResetValueRequired bool
	PTolerenceValue      *float64
	Tag                  *string
	Threshold            *float64
	TolerenceValue       *float64
}

type RuleGQL struct {
	Condition     string
	LeftHandSide  *RuleLeftHandSideGQL
	ResetValue    *RuleResetValueGQL
	RightHandSide *RuleRightHandSideGQL
}

type RuleSetGQL struct {
	RuleAndOrOperationData *RuleSetAndOrOperationDataGQL
	Rules                  []*RuleGQL
}

type AlarmConfigurationGQL struct {
	ID                        string
	AlarmName                 string
	AlarmDescription          string
	AlarmType                 string
	AlarmCategory             string
	AlarmDuration             string
	AlarmTemplate             string
	AlarmSmsTemplateID        *string
	ClientID                  string
	SiteID                    string
	EventUserID               *string
	Priority                  *string
	PriorityLevel             *int
	Acknowledgement           *bool
	Edited                    *bool
	Enabled                   *bool
	IsDeleted                 *bool
	IsTypeIsAlarm             *bool
	Devices                   []string
	Levels                    []*AlarmLevelGQL
	RuleSetAndOrOperationData *RuleSetAndOrOperationDataGQL
	RuleSets                  []*RuleSetGQL
}

func toAlarmConfigurationGQL(ac *AlarmConfiguration) *AlarmConfigurationGQL {
	var levels []*AlarmLevelGQL
	for _, l := range ac.Levels {
		var nps []*AlarmNotificationProfileGQL
		for _, np := range l.NotificationProfiles {
			var users []*AlarmNotificationProfileUserGQL
			for _, u := range np.UsersOrUserGroup {
				users = append(users, &AlarmNotificationProfileUserGQL{
					Label: u.Label,
					Type:  u.Type,
					Value: u.Value,
				})
			}
			isNotificationToneShow := bool(np.IsNotificationToneShow)
			nps = append(nps, &AlarmNotificationProfileGQL{
				UsersOrUserGroup:       users,
				IsNotificationToneShow: &isNotificationToneShow,
				NotificationProfile:    []string(np.NotificationProfile),
				NotificationTone:       &np.NotificationTone,
			})
		}
		levels = append(levels, &AlarmLevelGQL{
			Commands:             []string(l.Commands),
			NotificationProfiles: nps,
			Suppress:             &l.Suppress,
		})
	}

	var ruleSets []*RuleSetGQL
	for _, rs := range ac.RuleSets {
		var rules []*RuleGQL
		for _, r := range rs.Rules {
			var customValue *float64
			if r.RightHandSide.CustomValue != nil {
				cv := float64(*r.RightHandSide.CustomValue)
				customValue = &cv
			}
			var thresholdVal *float64
			if r.RightHandSide.Threshold != nil {
				th := float64(*r.RightHandSide.Threshold)
				thresholdVal = &th
			}

			var resetCustomValue *float64
			if r.ResetValue.CustomValue != nil {
				rcv := float64(*r.ResetValue.CustomValue)
				resetCustomValue = &rcv
			}
			var resetPTolerenceValue *float64
			if r.ResetValue.PTolerenceValue != nil {
				rpt := float64(*r.ResetValue.PTolerenceValue)
				resetPTolerenceValue = &rpt
			}
			var resetThresholdVal *float64
			if r.ResetValue.Threshold != nil {
				rth := float64(*r.ResetValue.Threshold)
				resetThresholdVal = &rth
			}
			var resetTolerenceValue *float64
			if r.ResetValue.TolerenceValue != nil {
				rtv := float64(*r.ResetValue.TolerenceValue)
				resetTolerenceValue = &rtv
			}

			rules = append(rules, &RuleGQL{
				Condition: r.Condition,
				LeftHandSide: &RuleLeftHandSideGQL{
					Tag: r.LeftHandSide.Tag,
				},
				ResetValue: &RuleResetValueGQL{
					CompareOption:        r.ResetValue.CompareOption,
					CustomValue:          resetCustomValue,
					IsResetValueRequired: bool(r.ResetValue.IsResetValueRequired),
					PTolerenceValue:      resetPTolerenceValue,
					Tag:                  r.ResetValue.Tag,
					Threshold:            resetThresholdVal,
					TolerenceValue:       resetTolerenceValue,
				},
				RightHandSide: &RuleRightHandSideGQL{
					CompareOption: r.RightHandSide.CompareOption,
					CustomValue:   customValue,
					Tag:           r.RightHandSide.Tag,
					Threshold:     thresholdVal,
				},
			})
		}
		ruleSets = append(ruleSets, &RuleSetGQL{
			RuleAndOrOperationData: &RuleSetAndOrOperationDataGQL{
				IsAnd: bool(rs.RuleAndOrOperationData.IsAnd),
				IsOr:  bool(rs.RuleAndOrOperationData.IsOr),
			},
			Rules: rules,
		})
	}

	priorityLevel := int(ac.PriorityLevel)
	acknowledgement := bool(ac.Acknowledgement)
	edited := bool(ac.Edited)
	enabled := bool(ac.Enabled)
	isDeleted := bool(ac.IsDeleted)
	isTypeIsAlarm := bool(ac.IsTypeIsAlarm)
	priority := string(ac.Priority)

	return &AlarmConfigurationGQL{
		ID:                 ac.ID,
		AlarmName:          ac.AlarmName,
		AlarmDescription:   ac.AlarmDescription,
		AlarmType:          ac.AlarmType,
		AlarmCategory:      ac.AlarmCategory,
		AlarmDuration:      string(ac.AlarmDuration),
		AlarmTemplate:      ac.AlarmTemplate,
		AlarmSmsTemplateID: &ac.AlarmSMSTemplateID,
		ClientID:           ac.ClientID,
		SiteID:             ac.SiteID,
		EventUserID:        &ac.EventUserID,
		Priority:           &priority,
		PriorityLevel:      &priorityLevel,
		Acknowledgement:    &acknowledgement,
		Edited:             &edited,
		Enabled:            &enabled,
		IsDeleted:          &isDeleted,
		IsTypeIsAlarm:      &isTypeIsAlarm,
		Devices:            []string(ac.Devices),
		Levels:             levels,
		RuleSetAndOrOperationData: &RuleSetAndOrOperationDataGQL{
			IsAnd: bool(ac.RuleSetAndOrOperationData.IsAnd),
			IsOr:  bool(ac.RuleSetAndOrOperationData.IsOr),
		},
		RuleSets: ruleSets,
	}
}

type BillMasterDataGQL struct {
	Customers          float64
	DisconnectionDays  float64
	ElectricalDuties   []string
	ExcelFile          *string
	FixedCostDeduction []float64
	FromDate           *string
	ToDate             *string
	SelectTods         []string
	AddDeduction       []float64
	BillCycle          float64
	BillDueDays        float64
	IsTax              *float64
	StartTime          *string
}

type BillMasterGQL struct {
	ID                string
	ClientID          string
	CustomerName      string
	CustomerAddress   *string
	CustomerID        float64
	CustomerPh        *string
	CustomerType      *float64
	DeviceInstanceID  string
	GatewayInstanceID string
	SiteID            string
	EbTag             *string
	DgTag             *string
	FromDate          *string
	StartTime         *float64
	FixedChargeDg     *float64
	FixedChargeEb     *float64
	OtherCharges      *float64
	IsDeleted         *string
	IsDisabled        *string
	FlatNo            *string
	SqurArea          *string
	Data              *BillMasterDataGQL
}

func toBillMasterGQL(bm *BillMaster) *BillMasterGQL {
	var fixedCosts []float64
	for _, fc := range bm.Data.FixedCostDeduction {
		fixedCosts = append(fixedCosts, float64(fc))
	}
	var addDeductions []float64
	for _, ad := range bm.Data.AddDeduction {
		addDeductions = append(addDeductions, float64(ad))
	}

	var fixedChargeEbVal *float64
	if bm.FixedChargeEb != nil {
		v := float64(*bm.FixedChargeEb)
		fixedChargeEbVal = &v
	}

	isTax := float64(bm.Data.IsTax)
	fixedChargeDg := float64(bm.FixedChargeDg)
	otherCharges := float64(bm.OtherCharges)
	startTime := float64(bm.StartTime)
	custID := float64(bm.CustomerID)
	custType := float64(bm.CustomerType)
	customerPh := string(bm.CustomerPh)
	flatNo := string(bm.FlatNo)
	squrArea := string(bm.SqurArea)

	return &BillMasterGQL{
		ID:                bm.ID,
		ClientID:          bm.ClientID,
		CustomerName:      bm.CustomerName,
		CustomerAddress:   &bm.CustomerAddress,
		CustomerID:        custID,
		CustomerPh:        &customerPh,
		CustomerType:      &custType,
		DeviceInstanceID:  bm.DeviceInstanceID,
		GatewayInstanceID: bm.GatewayInstanceID,
		SiteID:            bm.SiteID,
		EbTag:             &bm.EbTag,
		DgTag:             &bm.DgTag,
		FromDate:          &bm.FromDate,
		StartTime:         &startTime,
		FixedChargeDg:     &fixedChargeDg,
		FixedChargeEb:     fixedChargeEbVal,
		OtherCharges:      &otherCharges,
		IsDeleted:         &bm.IsDeleted,
		IsDisabled:        &bm.IsDisabled,
		FlatNo:            &flatNo,
		SqurArea:          &squrArea,
		Data: &BillMasterDataGQL{
			Customers:          float64(bm.Data.Customers),
			DisconnectionDays:  float64(bm.Data.DisconnectionDays),
			ElectricalDuties:   []string(bm.Data.ElectricalDuties),
			ExcelFile:          &bm.Data.ExcelFile,
			FixedCostDeduction: fixedCosts,
			FromDate:           &bm.Data.FromDate,
			ToDate:             &bm.Data.ToDate,
			SelectTods:         []string(bm.Data.SelectTods),
			AddDeduction:       addDeductions,
			BillCycle:          float64(bm.Data.BillCycle),
			BillDueDays:        float64(bm.Data.BillDueDays),
			IsTax:              &isTax,
			StartTime:          &bm.Data.StartTime,
		},
	}
}

type EmailGatewayGQL struct {
	ID                  string
	EmailGatewayID      string
	EmailType           string
	Encryption          *string
	MailServer          string
	Username            string
	Password            string
	SenderEmail         string
	SenderName          *string
	SMTPPort            float64
	SMTPAuthorization   *bool
	ProfileName         *string
	ConfigurationStatus *string
	SiteID              *string
	ClientID            *string
	EventUserID         *string
	Default             *bool
	SystemDefaults      *bool
	IsDeleted           *bool
}

func toEmailGatewayGQL(eg *EmailGateway) *EmailGatewayGQL {
	smtpPort := float64(eg.SMTPPort)
	smtpAuthorization := bool(eg.SMTPAuthorization)
	isDefault := bool(eg.Default)
	systemDefaults := bool(eg.SystemDefaults)
	isDeleted := bool(eg.IsDeleted)

	return &EmailGatewayGQL{
		ID:                  eg.ID,
		EmailGatewayID:      eg.EmailGatewayID,
		EmailType:           eg.EmailType,
		Encryption:          &eg.Encryption,
		MailServer:          eg.MailServer,
		Username:            eg.Username,
		Password:            eg.Password,
		SenderEmail:         eg.SenderEmail,
		SenderName:          &eg.SenderName,
		SMTPPort:            smtpPort,
		SMTPAuthorization:   &smtpAuthorization,
		ProfileName:         &eg.ProfileName,
		ConfigurationStatus: &eg.ConfigurationStatus,
		SiteID:              &eg.SiteID,
		ClientID:            &eg.ClientID,
		EventUserID:         &eg.EventUserID,
		Default:             &isDefault,
		SystemDefaults:      &systemDefaults,
		IsDeleted:           &isDeleted,
	}
}

type WorkGroupGQL struct {
	ID                   string
	WorkGroupID          string
	WorkGroupName        string
	WorkGroupDescription *string
	SiteID               string
	ClientID             string
	EventUserID          *string
	Default              *bool
	IsDeleted            *bool
}

func toWorkGroupGQL(wg *WorkGroup) *WorkGroupGQL {
	isDefault := bool(wg.Default)
	isDeleted := bool(wg.IsDeleted)
	return &WorkGroupGQL{
		ID:                   wg.ID,
		WorkGroupID:          wg.WorkGroupID,
		WorkGroupName:        wg.WorkGroupName,
		WorkGroupDescription: &wg.WorkGroupDescription,
		SiteID:               wg.SiteID,
		ClientID:             wg.ClientID,
		EventUserID:          &wg.EventUserID,
		Default:              &isDefault,
		IsDeleted:            &isDeleted,
	}
}

type SmartMeterSensorIDGQL struct {
	SensorID   string `json:"sensorId"`
	SensorName string `json:"sensorName"`
}

type SmartMeterGeneralInfoBodyGQL struct {
	Area                      *float64                 `json:"area"`
	AvenueNo                  *string                  `json:"avenueNo"`
	BalanceType               *float64                 `json:"balanceType"`
	BillType                  *string                  `json:"billType"`
	BlockNumber               *string                  `json:"blockNumber"`
	BlockType                 *string                  `json:"blockType"`
	CostManagementID          *string                  `json:"costManagementId"`
	CustomerAddress           *string                  `json:"customerAddress"`
	CustomerID                *float64                 `json:"customerId"`
	CustomerName              *string                  `json:"customerName"`
	DgTag                     *string                  `json:"dgTag"`
	DgThresholdKva            *float64                 `json:"dgThresholdKva"`
	EbTag                     *string                  `json:"ebTag"`
	EmailID                   *string                  `json:"emailId"`
	FixedChargeDg             *float64                 `json:"fixedChargeDg"`
	FixedChargeEb             *float64                 `json:"fixedChargeEb"`
	Flat                      *string                  `json:"flat"`
	FlatNo                    *string                  `json:"flatNo"`
	GSTApplicable             []string                 `json:"gstApplicable"`
	GSTNumber                 *string                  `json:"gstNumber"`
	IsDeleted                 []string                 `json:"isDeleted"`
	IsDisabled                []string                 `json:"isDisabled"`
	MeterModel                *string                  `json:"meterModel"`
	MeterType                 *string                  `json:"meterType"`
	MobileNo                  *float64                 `json:"mobileNo"`
	OccupancyDate             *string                  `json:"occupancyDate"`
	OtherCharges              *float64                 `json:"otherCharges"`
	PlanType                  *string                  `json:"planType"`
	RelayStatus               *string                  `json:"relayStatus"`
	SensorID                  []*SmartMeterSensorIDGQL `json:"sensorId"`
	ServiceProvider           *string                  `json:"serviceProvider"`
	ServiceProviderID         *string                  `json:"serviceProviderId"`
	SmartMeterConfigurationID *string                  `json:"smartMeterConfigurationId"`
	SubscriberType            *float64                 `json:"subscriberType"`
	Tag1624                   *float64                 `json:"tag1624"`
	Tag1625                   *float64                 `json:"tag1625"`
	Tag1626                   *float64                 `json:"tag1626"`
	Tag1627                   *float64                 `json:"tag1627"`
	Tag1628                   *float64                 `json:"tag1628"`
	Tag1629                   *float64                 `json:"tag1629"`
	Tag1630                   *float64                 `json:"tag1630"`
	Tag1631                   *float64                 `json:"tag1631"`
	Tag2952                   *float64                 `json:"tag2952"`
	Tag3046                   *float64                 `json:"tag3046"`
	Tag3047                   *float64                 `json:"tag3047"`
	TariffCateg               *string                  `json:"tariffCateg"`
	TariffCategoryID          []string                 `json:"tariffCategoryId"`
	Tower                     *string                  `json:"tower"`
	VacancyDate               *string                  `json:"vacancyDate"`
}

type SmartMeterGeneralInfoGQL struct {
	BodyContent *SmartMeterGeneralInfoBodyGQL
}

type SmartMeterSlabSettingBodyGQL struct {
	BlockNumber string                   `json:"blockNumber"`
	BlockType   string                   `json:"blockType"`
	SensorID    []*SmartMeterSensorIDGQL `json:"sensorId"`
	Tag1655     *float64                 `json:"tag1655"`
	Tag1657     *float64                 `json:"tag1657"`
	Tag1673     *float64                 `json:"tag1673"`
	Tag1674     *float64                 `json:"tag1674"`
	Tag1675     *float64                 `json:"tag1675"`
	Tag1676     *float64                 `json:"tag1676"`
	Tag1677     *float64                 `json:"tag1677"`
	Tag1678     *float64                 `json:"tag1678"`
	Tag1681     *float64                 `json:"tag1681"`
	Tag1682     *float64                 `json:"tag1682"`
	Tag1683     *float64                 `json:"tag1683"`
	Tag2900     *float64                 `json:"tag2900"`
	Tag2901     *float64                 `json:"tag2901"`
	Tag2902     *string                  `json:"tag2902"`
	Tag2903     *string                  `json:"tag2903"`
	Tag2904     *float64                 `json:"tag2904"`
	Tag2905     *float64                 `json:"tag2905"`
	Tag2906     *float64                 `json:"tag2906"`
	Tag2907     *float64                 `json:"tag2907"`
	Tag2908     *float64                 `json:"tag2908"`
	Tag2909     *float64                 `json:"tag2909"`
	Tag2910     *float64                 `json:"tag2910"`
	Tag2911     *float64                 `json:"tag2911"`
	Tag2912     *float64                 `json:"tag2912"`
	Tag2913     *float64                 `json:"tag2913"`
	Tag2914     *float64                 `json:"tag2914"`
	Tag2915     *float64                 `json:"tag2915"`
	Tag2916     *float64                 `json:"tag2916"`
	Tag2917     *float64                 `json:"tag2917"`
	Tag2918     *float64                 `json:"tag2918"`
	Tag2919     *float64                 `json:"tag2919"`
	Tag2920     *float64                 `json:"tag2920"`
	Tag2921     *float64                 `json:"tag2921"`
	Tag2922     *float64                 `json:"tag2922"`
	Tag2923     *float64                 `json:"tag2923"`
	Tag2924     *float64                 `json:"tag2924"`
	Tag2925     *float64                 `json:"tag2925"`
	Tag2926     *float64                 `json:"tag2926"`
	Tag2927     *float64                 `json:"tag2927"`
	Tag2928     *float64                 `json:"tag2928"`
	Tag2929     *float64                 `json:"tag2929"`
	Tag2930     *float64                 `json:"tag2930"`
	Tag2931     *float64                 `json:"tag2931"`
	Tag2932     *float64                 `json:"tag2932"`
	Tag2933     *float64                 `json:"tag2933"`
	Tag2934     *float64                 `json:"tag2934"`
	Tag2935     *float64                 `json:"tag2935"`
	Tag2936     *float64                 `json:"tag2936"`
	Tag2952     *float64                 `json:"tag2952"`
	Tag2973     *float64                 `json:"tag2973"`
	Tag3054     *float64                 `json:"tag3054"`
	Tag3521     *float64                 `json:"tag3521"`
}

type SmartMeterSlabSettingGQL struct {
	BodyContent *SmartMeterSlabSettingBodyGQL
}

type SmartMeterConfigurationGQL struct {
	ID               string
	ClientID         string
	DeviceInstanceID string
	SensorID         string
	SiteID           string
	Isdeleted        *string
	Isdisabled       *string
	GeneralInfo      *SmartMeterGeneralInfoGQL
	SlabSetting      *SmartMeterSlabSettingGQL
}

func toSmartMeterConfigurationGQL(sm *SmartMeterConfiguration) *SmartMeterConfigurationGQL {
	var genSensors []*SmartMeterSensorIDGQL
	for _, s := range sm.GeneralInfo.BodyContent.SensorID {
		genSensors = append(genSensors, &SmartMeterSensorIDGQL{
			SensorID:   s.SensorID,
			SensorName: s.SensorName,
		})
	}

	var slabSensors []*SmartMeterSensorIDGQL
	for _, s := range sm.SlabSetting.BodyContent.SensorID {
		slabSensors = append(slabSensors, &SmartMeterSensorIDGQL{
			SensorID:   s.SensorID,
			SensorName: s.SensorName,
		})
	}

	area := float64(sm.GeneralInfo.BodyContent.Area)
	balanceType := float64(sm.GeneralInfo.BodyContent.BalanceType)
	customerID := float64(sm.GeneralInfo.BodyContent.CustomerID)
	dgThresholdKva := float64(sm.GeneralInfo.BodyContent.DgThresholdKva)
	fixedChargeDg := float64(sm.GeneralInfo.BodyContent.FixedChargeDg)
	fixedChargeEb := float64(sm.GeneralInfo.BodyContent.FixedChargeEb)
	mobileNo := float64(sm.GeneralInfo.BodyContent.MobileNo)
	otherCharges := float64(sm.GeneralInfo.BodyContent.OtherCharges)
	subscriberType := float64(sm.GeneralInfo.BodyContent.SubscriberType)

	tag1624 := float64(sm.GeneralInfo.BodyContent.Tag1624)
	tag1625 := float64(sm.GeneralInfo.BodyContent.Tag1625)
	tag1626 := float64(sm.GeneralInfo.BodyContent.Tag1626)
	tag1627 := float64(sm.GeneralInfo.BodyContent.Tag1627)
	tag1628 := float64(sm.GeneralInfo.BodyContent.Tag1628)
	tag1629 := float64(sm.GeneralInfo.BodyContent.Tag1629)
	tag1630 := float64(sm.GeneralInfo.BodyContent.Tag1630)
	tag1631 := float64(sm.GeneralInfo.BodyContent.Tag1631)
	tag2952_gen := float64(sm.GeneralInfo.BodyContent.Tag2952)
	tag3046 := float64(sm.GeneralInfo.BodyContent.Tag3046)
	tag3047 := float64(sm.GeneralInfo.BodyContent.Tag3047)

	tag1655 := float64(sm.SlabSetting.BodyContent.Tag1655)
	tag1657 := float64(sm.SlabSetting.BodyContent.Tag1657)
	tag1673 := float64(sm.SlabSetting.BodyContent.Tag1673)
	tag1674 := float64(sm.SlabSetting.BodyContent.Tag1674)
	tag1675 := float64(sm.SlabSetting.BodyContent.Tag1675)
	tag1676 := float64(sm.SlabSetting.BodyContent.Tag1676)
	tag1677 := float64(sm.SlabSetting.BodyContent.Tag1677)
	tag1678 := float64(sm.SlabSetting.BodyContent.Tag1678)
	tag1681 := float64(sm.SlabSetting.BodyContent.Tag1681)
	tag1682 := float64(sm.SlabSetting.BodyContent.Tag1682)
	tag1683 := float64(sm.SlabSetting.BodyContent.Tag1683)
	tag2900 := float64(sm.SlabSetting.BodyContent.Tag2900)
	tag2901 := float64(sm.SlabSetting.BodyContent.Tag2901)
	tag2904 := float64(sm.SlabSetting.BodyContent.Tag2904)
	tag2905 := float64(sm.SlabSetting.BodyContent.Tag2905)
	tag2906 := float64(sm.SlabSetting.BodyContent.Tag2906)
	tag2907 := float64(sm.SlabSetting.BodyContent.Tag2907)
	tag2908 := float64(sm.SlabSetting.BodyContent.Tag2908)
	tag2909 := float64(sm.SlabSetting.BodyContent.Tag2909)
	tag2910 := float64(sm.SlabSetting.BodyContent.Tag2910)
	tag2911 := float64(sm.SlabSetting.BodyContent.Tag2911)
	tag2912 := float64(sm.SlabSetting.BodyContent.Tag2912)
	tag2913 := float64(sm.SlabSetting.BodyContent.Tag2913)
	tag2914 := float64(sm.SlabSetting.BodyContent.Tag2914)
	tag2915 := float64(sm.SlabSetting.BodyContent.Tag2915)
	tag2916 := float64(sm.SlabSetting.BodyContent.Tag2916)
	tag2917 := float64(sm.SlabSetting.BodyContent.Tag2917)
	tag2918 := float64(sm.SlabSetting.BodyContent.Tag2918)
	tag2919 := float64(sm.SlabSetting.BodyContent.Tag2919)
	tag2920 := float64(sm.SlabSetting.BodyContent.Tag2920)
	tag2921 := float64(sm.SlabSetting.BodyContent.Tag2921)
	tag2922 := float64(sm.SlabSetting.BodyContent.Tag2922)
	tag2923 := float64(sm.SlabSetting.BodyContent.Tag2923)
	tag2924 := float64(sm.SlabSetting.BodyContent.Tag2924)
	tag2925 := float64(sm.SlabSetting.BodyContent.Tag2925)
	tag2926 := float64(sm.SlabSetting.BodyContent.Tag2926)
	tag2927 := float64(sm.SlabSetting.BodyContent.Tag2927)
	tag2928 := float64(sm.SlabSetting.BodyContent.Tag2928)
	tag2929 := float64(sm.SlabSetting.BodyContent.Tag2929)
	tag2930 := float64(sm.SlabSetting.BodyContent.Tag2930)
	tag2931 := float64(sm.SlabSetting.BodyContent.Tag2931)
	tag2932 := float64(sm.SlabSetting.BodyContent.Tag2932)
	tag2933 := float64(sm.SlabSetting.BodyContent.Tag2933)
	tag2934 := float64(sm.SlabSetting.BodyContent.Tag2934)
	tag2935 := float64(sm.SlabSetting.BodyContent.Tag2935)
	tag2936 := float64(sm.SlabSetting.BodyContent.Tag2936)
	tag2952_slab := float64(sm.SlabSetting.BodyContent.Tag2952)
	tag2973 := float64(sm.SlabSetting.BodyContent.Tag2973)
	tag3054 := float64(sm.SlabSetting.BodyContent.Tag3054)
	tag3521 := float64(sm.SlabSetting.BodyContent.Tag3521)

	avenueNo := string(sm.GeneralInfo.BodyContent.AvenueNo)
	blockNumber := string(sm.GeneralInfo.BodyContent.BlockNumber)
	flatNo := string(sm.GeneralInfo.BodyContent.FlatNo)
	occupancyDate := string(sm.GeneralInfo.BodyContent.OccupancyDate)
	vacancyDate := string(sm.GeneralInfo.BodyContent.VacancyDate)

	return &SmartMeterConfigurationGQL{
		ID:               sm.ID,
		ClientID:         sm.ClientID,
		DeviceInstanceID: sm.DeviceInstanceID,
		SensorID:         sm.SensorID,
		SiteID:           sm.SiteID,
		Isdeleted:        &sm.IsDeleted,
		Isdisabled:       &sm.IsDisabled,
		GeneralInfo: &SmartMeterGeneralInfoGQL{
			BodyContent: &SmartMeterGeneralInfoBodyGQL{
				Area:                      &area,
				AvenueNo:                  &avenueNo,
				BalanceType:               &balanceType,
				BillType:                  &sm.GeneralInfo.BodyContent.BillType,
				BlockNumber:               &blockNumber,
				BlockType:                 &sm.GeneralInfo.BodyContent.BlockType,
				CostManagementID:          &sm.GeneralInfo.BodyContent.CostManagementID,
				CustomerAddress:           &sm.GeneralInfo.BodyContent.CustomerAddress,
				CustomerID:                &customerID,
				CustomerName:              &sm.GeneralInfo.BodyContent.CustomerName,
				DgTag:                     &sm.GeneralInfo.BodyContent.DgTag,
				DgThresholdKva:            &dgThresholdKva,
				EbTag:                     &sm.GeneralInfo.BodyContent.EbTag,
				EmailID:                   &sm.GeneralInfo.BodyContent.EmailID,
				FixedChargeDg:             &fixedChargeDg,
				FixedChargeEb:             &fixedChargeEb,
				Flat:                      &sm.GeneralInfo.BodyContent.Flat,
				FlatNo:                    &flatNo,
				GSTApplicable:             []string(sm.GeneralInfo.BodyContent.GSTApplicable),
				GSTNumber:                 &sm.GeneralInfo.BodyContent.GSTNumber,
				IsDeleted:                 []string(sm.GeneralInfo.BodyContent.IsDeleted),
				IsDisabled:                []string(sm.GeneralInfo.BodyContent.IsDisabled),
				MeterModel:                &sm.GeneralInfo.BodyContent.MeterModel,
				MeterType:                 &sm.GeneralInfo.BodyContent.MeterType,
				MobileNo:                  &mobileNo,
				OccupancyDate:             &occupancyDate,
				OtherCharges:              &otherCharges,
				PlanType:                  &sm.GeneralInfo.BodyContent.PlanType,
				RelayStatus:               &sm.GeneralInfo.BodyContent.RelayStatus,
				SensorID:                  genSensors,
				ServiceProvider:           &sm.GeneralInfo.BodyContent.ServiceProvider,
				ServiceProviderID:         &sm.GeneralInfo.BodyContent.ServiceProviderID,
				SmartMeterConfigurationID: &sm.GeneralInfo.BodyContent.SmartMeterConfigurationID,
				SubscriberType:            &subscriberType,
				Tag1624:                   &tag1624,
				Tag1625:                   &tag1625,
				Tag1626:                   &tag1626,
				Tag1627:                   &tag1627,
				Tag1628:                   &tag1628,
				Tag1629:                   &tag1629,
				Tag1630:                   &tag1630,
				Tag1631:                   &tag1631,
				Tag2952:                   &tag2952_gen,
				Tag3046:                   &tag3046,
				Tag3047:                   &tag3047,
				TariffCateg:               &sm.GeneralInfo.BodyContent.TariffCateg,
				TariffCategoryID:          []string(sm.GeneralInfo.BodyContent.TariffCategoryID),
				Tower:                     &sm.GeneralInfo.BodyContent.Tower,
				VacancyDate:               &vacancyDate,
			},
		},
		SlabSetting: &SmartMeterSlabSettingGQL{
			BodyContent: &SmartMeterSlabSettingBodyGQL{
				BlockNumber: sm.SlabSetting.BodyContent.BlockNumber,
				BlockType:   sm.SlabSetting.BodyContent.BlockType,
				SensorID:    slabSensors,
				Tag1655:     &tag1655,
				Tag1657:     &tag1657,
				Tag1673:     &tag1673,
				Tag1674:     &tag1674,
				Tag1675:     &tag1675,
				Tag1676:     &tag1676,
				Tag1677:     &tag1677,
				Tag1678:     &tag1678,
				Tag1681:     &tag1681,
				Tag1682:     &tag1682,
				Tag1683:     &tag1683,
				Tag2900:     &tag2900,
				Tag2901:     &tag2901,
				Tag2902:     &sm.SlabSetting.BodyContent.Tag2902,
				Tag2903:     &sm.SlabSetting.BodyContent.Tag2903,
				Tag2904:     &tag2904,
				Tag2905:     &tag2905,
				Tag2906:     &tag2906,
				Tag2907:     &tag2907,
				Tag2908:     &tag2908,
				Tag2909:     &tag2909,
				Tag2910:     &tag2910,
				Tag2911:     &tag2911,
				Tag2912:     &tag2912,
				Tag2913:     &tag2913,
				Tag2914:     &tag2914,
				Tag2915:     &tag2915,
				Tag2916:     &tag2916,
				Tag2917:     &tag2917,
				Tag2918:     &tag2918,
				Tag2919:     &tag2919,
				Tag2920:     &tag2920,
				Tag2921:     &tag2921,
				Tag2922:     &tag2922,
				Tag2923:     &tag2923,
				Tag2924:     &tag2924,
				Tag2925:     &tag2925,
				Tag2926:     &tag2926,
				Tag2927:     &tag2927,
				Tag2928:     &tag2928,
				Tag2929:     &tag2929,
				Tag2930:     &tag2930,
				Tag2931:     &tag2931,
				Tag2932:     &tag2932,
				Tag2933:     &tag2933,
				Tag2934:     &tag2934,
				Tag2935:     &tag2935,
				Tag2936:     &tag2936,
				Tag2952:     &tag2952_slab,
				Tag2973:     &tag2973,
				Tag3054:     &tag3054,
				Tag3521:     &tag3521,
			},
		},
	}
}
