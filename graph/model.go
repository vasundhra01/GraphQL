package graph

import (
	"encoding/binary"
	"math"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type FlexibleFloat64 float64

func (ff *FlexibleFloat64) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	switch t {
	case bsontype.Double:
		if len(data) < 8 {
			*ff = 0
			return nil
		}
		bits := binary.LittleEndian.Uint64(data)
		*ff = FlexibleFloat64(math.Float64frombits(bits))
	case bsontype.Int32:
		if len(data) < 4 {
			*ff = 0
			return nil
		}
		val := int32(binary.LittleEndian.Uint32(data))
		*ff = FlexibleFloat64(val)
	case bsontype.Int64:
		if len(data) < 8 {
			*ff = 0
			return nil
		}
		val := int64(binary.LittleEndian.Uint64(data))
		*ff = FlexibleFloat64(val)
	case bsontype.String:
		if len(data) < 4 {
			*ff = 0
			return nil
		}
		length := int(binary.LittleEndian.Uint32(data))
		if len(data) < 4+length {
			*ff = 0
			return nil
		}
		strVal := string(data[4 : 4+length-1])
		parsed, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			*ff = 0
		} else {
			*ff = FlexibleFloat64(parsed)
		}
	case bsontype.Boolean:
		if len(data) > 0 && data[0] == 1 {
			*ff = 1
		} else {
			*ff = 0
		}
	default:
		*ff = 0
	}
	return nil
}

type FlexibleInt int

func (fi *FlexibleInt) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	switch t {
	case bsontype.Int32:
		if len(data) < 4 {
			*fi = 0
			return nil
		}
		*fi = FlexibleInt(int32(binary.LittleEndian.Uint32(data)))
	case bsontype.Int64:
		if len(data) < 8 {
			*fi = 0
			return nil
		}
		*fi = FlexibleInt(int64(binary.LittleEndian.Uint64(data)))
	case bsontype.Double:
		if len(data) < 8 {
			*fi = 0
			return nil
		}
		bits := binary.LittleEndian.Uint64(data)
		*fi = FlexibleInt(math.Float64frombits(bits))
	case bsontype.String:
		if len(data) < 4 {
			*fi = 0
			return nil
		}
		length := int(binary.LittleEndian.Uint32(data))
		if len(data) < 4+length {
			*fi = 0
			return nil
		}
		strVal := string(data[4 : 4+length-1])
		parsed, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			*fi = 0
		} else {
			*fi = FlexibleInt(parsed)
		}
	case bsontype.Boolean:
		if len(data) > 0 && data[0] == 1 {
			*fi = 1
		} else {
			*fi = 0
		}
	default:
		*fi = 0
	}
	return nil
}

type FlexibleBool bool

func (fb *FlexibleBool) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	switch t {
	case bsontype.Boolean:
		if len(data) > 0 && data[0] == 1 {
			*fb = true
		} else {
			*fb = false
		}
	case bsontype.String:
		if len(data) < 4 {
			*fb = false
			return nil
		}
		length := int(binary.LittleEndian.Uint32(data))
		if len(data) < 4+length {
			*fb = false
			return nil
		}
		strVal := string(data[4 : 4+length-1])
		*fb = FlexibleBool(strVal == "true" || strVal == "1" || strVal == "yes")
	case bsontype.Int32:
		if len(data) < 4 {
			*fb = false
			return nil
		}
		val := int32(binary.LittleEndian.Uint32(data))
		*fb = FlexibleBool(val != 0)
	case bsontype.Int64:
		if len(data) < 8 {
			*fb = false
			return nil
		}
		val := int64(binary.LittleEndian.Uint64(data))
		*fb = FlexibleBool(val != 0)
	case bsontype.Double:
		if len(data) < 8 {
			*fb = false
			return nil
		}
		bits := binary.LittleEndian.Uint64(data)
		val := math.Float64frombits(bits)
		*fb = FlexibleBool(val != 0)
	default:
		*fb = false
	}
	return nil
}

type FlexibleString string

func (fs *FlexibleString) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	switch t {
	case bsontype.String:
		if len(data) < 4 {
			*fs = ""
			return nil
		}
		length := int(binary.LittleEndian.Uint32(data))
		if len(data) < 4+length {
			*fs = ""
			return nil
		}
		*fs = FlexibleString(data[4 : 4+length-1])
	case bsontype.Int32:
		if len(data) < 4 {
			*fs = ""
			return nil
		}
		val := int32(binary.LittleEndian.Uint32(data))
		*fs = FlexibleString(strconv.FormatInt(int64(val), 10))
	case bsontype.Int64:
		if len(data) < 8 {
			*fs = ""
			return nil
		}
		val := int64(binary.LittleEndian.Uint64(data))
		*fs = FlexibleString(strconv.FormatInt(val, 10))
	case bsontype.Double:
		if len(data) < 8 {
			*fs = ""
			return nil
		}
		bits := binary.LittleEndian.Uint64(data)
		val := math.Float64frombits(bits)
		*fs = FlexibleString(strconv.FormatFloat(val, 'f', -1, 64))
	case bsontype.Boolean:
		if len(data) > 0 && data[0] == 1 {
			*fs = "true"
		} else {
			*fs = "false"
		}
	default:
		*fs = ""
	}
	return nil
}

type FlexibleStringSlice []string

func (fss *FlexibleStringSlice) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	switch t {
	case bsontype.Array:
		var arr []string
		rawArray := bson.RawValue{Type: bsontype.Array, Value: data}
		if err := rawArray.Unmarshal(&arr); err != nil {
			*fss = nil
			return nil
		}
		*fss = FlexibleStringSlice(arr)
	case bsontype.String:
		if len(data) < 4 {
			*fss = nil
			return nil
		}
		length := int(binary.LittleEndian.Uint32(data))
		if len(data) < 4+length {
			*fss = nil
			return nil
		}
		strVal := string(data[4 : 4+length-1])
		if strVal == "" {
			*fss = nil
		} else {
			*fss = FlexibleStringSlice{strVal}
		}
	default:
		*fss = nil
	}
	return nil
}

type User struct {
	ID    string `bson:"_id,omitempty"`
	Name  string `bson:"fullName" json:"name"`
	Email string `bson:"eMail" json:"email"`
}

type DeviceInstanceTag struct {
	SensorID        string `bson:"sensor_id"`
	Source          string `bson:"source"`
	TagID           string `bson:"tag_id"`
	VirtualDeviceID string `bson:"virtual_device_id"`
}

type DeviceInstanceBelongsRef struct {
	ParentID string `bson:"parent_id"`
	NodeID   string `bson:"node_id"`
}

type DeviceInstanceAssignRef struct {
	ParentID string `bson:"parent_id"`
	NodeID   string `bson:"node_id"`
}

type DeviceInstanceInfo struct {
	PowerOutage      FlexibleBool               `bson:"power_outage"`
	BelongsRef       []DeviceInstanceBelongsRef `bson:"belongs_ref"`
	AssignRef        []DeviceInstanceAssignRef  `bson:"assign_ref"`
	Timeout          FlexibleString             `bson:"timeout"`
	DeviceComID      FlexibleString             `bson:"device_com_id"`
	Make             string                     `bson:"make"`
	ModelNumber      string                     `bson:"modelnumber"`
	ProtocolCategory FlexibleStringSlice        `bson:"protocolcategory"`
	DeviceModelName  string                     `bson:"device_model_name"`
	DeviceModelRefID string                     `bson:"device_model_ref_id"`
	DeviceName       string                     `bson:"device_name"`
	DeviceSelection  FlexibleStringSlice        `bson:"deviceselection"`
	IsDeleted        FlexibleBool               `bson:"isdeleted"`
	IsDisabled       FlexibleBool               `bson:"isdisabled"`
	MacID            string                     `bson:"mac_id"`
}

type DeviceInstance struct {
	ID                string             `bson:"_id"`
	DeviceInstanceID  string             `bson:"device_instance_id"`
	SiteID            string             `bson:"site_id"`
	GatewayID         string             `bson:"gateway_id"`
	ClientID          string             `bson:"client_id"`
	Default           FlexibleBool       `bson:"default"`
	GatewayInstanceID string             `bson:"gateway_instance_id"`
	EventUserID       string             `bson:"event_user_id"`
	GeneralInfo       DeviceInstanceInfo `bson:"general_info"`
	TagsData          []DeviceInstanceTag `bson:"tagsData"`
}

type DeviceModelInfo struct {
	DeviceModelName string `bson:"device_model_name"`
	Make            string `bson:"make"`
	ModelNumber     string `bson:"modelnumber"`
}

type DeviceModel struct {
	ID            string          `bson:"_id"`
	DeviceModelID string          `bson:"device_model_id"`
	GeneralInfo   DeviceModelInfo `bson:"general_info"`
}

type SensorInfo struct {
	DeviceModelName   string `bson:"device_model_name"`
	DeviceName        string `bson:"device_name"`
	Make              string `bson:"make"`
	ModelNumber       string `bson:"modelnumber"`
	GatewayInstanceID string `bson:"gateway_instance_id"`
}

type Sensor struct {
	ID               string     `bson:"_id"`
	DeviceInstanceID string     `bson:"device_instance_id"`
	SiteID           string     `bson:"site_id"`
	GeneralInfo      SensorInfo `bson:"general_info"`
}

type Batch struct {
	ID               string `bson:"_id"`
	BatchID          string `bson:"batch_id"`
	BatchName        string `bson:"batchName"`
	BatchDescription string `bson:"batchDescription"`
	SiteID           string `bson:"site_id"`
	ClientID         string `bson:"client_id"`
}

type Client struct {
	ID                string `bson:"_id"`
	ClientID          string `bson:"client_id"`
	ClientName        string `bson:"client_name"`
	ClientDescription string `bson:"client_description"`
	ClientAddress     string `bson:"client_address"`
	ClientLogo        string `bson:"client_logo"`
	UserID            string `bson:"user_id"`
}

type Cost struct {
	ID                  string          `bson:"_id"`
	CostManagementID    string          `bson:"cost_management_id"`
	CostManagementKey   string          `bson:"cost_management_key"`
	CostManagementKeyID string          `bson:"cost_management_key_id"`
	CostType            FlexibleInt     `bson:"cost_type"`
	DgBaseCost          FlexibleFloat64 `bson:"dg_base_cost"`
	DgCost              FlexibleFloat64 `bson:"dg_cost"`
	EbBaseCost          FlexibleFloat64 `bson:"eb_base_cost"`
	EbCost              FlexibleFloat64 `bson:"eb_cost"`
	SiteID              string          `bson:"site_id"`
	ClientID            string          `bson:"client_id"`
	FromDate            time.Time       `bson:"from_date"`
	FromDateStr         string          `bson:"from_date_str"`
	ToDate              time.Time       `bson:"to_date"`
	ToDateStr           string          `bson:"to_date_str"`
	RechargeTax         FlexibleFloat64 `bson:"recharge_tax"`
	ServiceProvider     string          `bson:"service_provider"`
	ServiceProviderID   string          `bson:"service_provider_id"`
	TariffCateg         string          `bson:"tariff_categ"`
	TariffCategoryID    string          `bson:"tariff_category_id"`
	CgstTax             FlexibleFloat64 `bson:"cgst_tax"`
	SgstTax             FlexibleFloat64 `bson:"sgst_tax"`
	CamCost             FlexibleFloat64 `bson:"cam_cost"`
	Ed1phase            FlexibleFloat64 `bson:"ed_1phase"`
	Ed3phase            FlexibleFloat64 `bson:"ed_3phase"`
}

type Site struct {
	ID           string `bson:"_id,omitempty"`
	SiteID       string `bson:"site_id"`
	SiteName     string `bson:"siteName"`
	ClientID     string `bson:"client_id"`
	IndustryID   string `bson:"industry_id"`
	IndustryName string `bson:"industry_name"`
}

type Gateway struct {
	ID                string          `bson:"_id,omitempty"`
	GatewayID         string          `bson:"gateway_id"`
	GatewayName       string          `bson:"gatewayname"`
	SiteID            string          `bson:"site_id"`
	DataReadFrequency FlexibleFloat64 `bson:"data_read_frequency"`
	GatewayModelRefID string          `bson:"gateway_model_ref_id"`
	GatewayModelName  string          `bson:"gatewaymodelname"`
	IsDeleted         string          `bson:"isdeleted"`
	Type              string          `bson:"type"`
	MacID             string          `bson:"mac_id"`
	Make              string          `bson:"make"`
	ModelNumber       string          `bson:"modelnumber"`
	SerialNo          string          `bson:"serialno"`
	UniqueID          string          `bson:"uniqueid"`
	UniqueIP          string          `bson:"uniqueip"`
	ClientID          string          `bson:"client_id"`
	Default           FlexibleBool    `bson:"default"`
	EventUserID       string          `bson:"event_user_id"`
	Timeout           FlexibleString  `bson:"timeout"`
	StatusTimeout     FlexibleFloat64 `bson:"status_timeout"`
}

type GatewayInstanceAssignRef struct {
	NodeID   string `bson:"node_id"`
	ParentID string `bson:"parent_id"`
}

type GatewayInstance struct {
	ID                   string                     `bson:"_id"`
	GatewayInstanceID    string                     `bson:"gateway_instance_id"`
	GatewayName          string                     `bson:"gatewayname"`
	MacID                string                     `bson:"mac_id"`
	SiteID               string                     `bson:"site_id"`
	ClientID             string                     `bson:"client_id"`
	AssignedIndustry     string                     `bson:"assigned_industry"`
	AwsLicense           FlexibleBool               `bson:"aws_license"`
	BaudRate             FlexibleString             `bson:"baudRate"`
	ComPort              FlexibleString             `bson:"com_port"`
	DataBit              FlexibleString             `bson:"dataBit"`
	Parity               FlexibleString             `bson:"parity"`
	StopBit              FlexibleString             `bson:"stopBit"`
	Protocol             string                     `bson:"protocol"`
	StatusTimeout        FlexibleFloat64            `bson:"status_timeout"`
	Timeout              FlexibleFloat64            `bson:"timeout"`
	Type                 string                     `bson:"type"`
	Default              FlexibleBool               `bson:"default"`
	IsConfigured         FlexibleBool               `bson:"isConfigured"`
	IsDeleted            FlexibleBool               `bson:"isdeleted"`
	IsDisabled           FlexibleBool               `bson:"isdisabled"`
	DataReadFrequency    FlexibleFloat64            `bson:"data_read_frequency"`
	DataLoggingFrequency FlexibleFloat64            `bson:"data_logging_frequency"`
	AssignRef            []GatewayInstanceAssignRef `bson:"assign_ref"`
	SerialNo             string                     `bson:"serialno"`
}

type AlarmNotificationProfileUser struct {
	Label string `bson:"label"`
	Type  string `bson:"type"`
	Value string `bson:"value"`
}

type AlarmNotificationProfile struct {
	UsersOrUserGroup       []AlarmNotificationProfileUser `bson:"usersOrUserGroup"`
	IsNotificationToneShow FlexibleBool                   `bson:"isNotificationToneShow"`
	NotificationProfile    FlexibleStringSlice            `bson:"notificationProfile"`
	NotificationTone       string                         `bson:"notificationTone"`
}

type AlarmLevel struct {
	Commands             FlexibleStringSlice        `bson:"commands"`
	NotificationProfiles []AlarmNotificationProfile `bson:"notificationProfiles"`
	Suppress             string                     `bson:"suppress"`
}

type RuleSetAndOrOperationData struct {
	IsAnd FlexibleBool `bson:"isAnd"`
	IsOr  FlexibleBool `bson:"isOr"`
}

type RuleLeftHandSide struct {
	Tag string `bson:"tag"`
}

type RuleRightHandSide struct {
	CompareOption string            `bson:"compareOption"`
	CustomValue   *FlexibleFloat64  `bson:"customValue"`
	Tag           *string           `bson:"tag"`
	Threshold     *FlexibleFloat64  `bson:"threshold"`
}

type RuleResetValue struct {
	CompareOption        string            `bson:"compareOption"`
	CustomValue          *FlexibleFloat64  `bson:"customValue"`
	IsResetValueRequired FlexibleBool      `bson:"isResetValueRequired"`
	PTolerenceValue      *FlexibleFloat64  `bson:"pTolerenceValue"`
	Tag                  *string           `bson:"tag"`
	Threshold            *FlexibleFloat64  `bson:"threshold"`
	TolerenceValue       *FlexibleFloat64  `bson:"tolerenceValue"`
}

type Rule struct {
	Condition     string            `bson:"condition"`
	LeftHandSide  RuleLeftHandSide  `bson:"leftHandSide"`
	ResetValue    RuleResetValue    `bson:"resetValue"`
	RightHandSide RuleRightHandSide `bson:"rightHandSide"`
}

type RuleSet struct {
	RuleAndOrOperationData RuleSetAndOrOperationData `bson:"ruleAndOrOperationData"`
	Rules                  []Rule                    `bson:"rules"`
}

type AlarmConfiguration struct {
	ID                        string                    `bson:"_id"`
	AlarmID                   string                    `bson:"id"`
	AlarmName                 string                    `bson:"alarmName"`
	AlarmDescription          string                    `bson:"alarmDescription"`
	AlarmType                 string                    `bson:"alarmType"`
	AlarmCategory             string                    `bson:"alarmCategory"`
	AlarmDuration             FlexibleString            `bson:"alarmDuration"`
	AlarmTemplate             string                    `bson:"alarmTemplate"`
	AlarmSMSTemplateID        string                    `bson:"alarm_sms_template_id"`
	ClientID                  string                    `bson:"client_id"`
	SiteID                    string                    `bson:"site_id"`
	EventUserID               string                    `bson:"event_user_id"`
	Priority                  FlexibleString            `bson:"priority"`
	PriorityLevel             FlexibleInt               `bson:"priority_level"`
	Acknowledgement           FlexibleBool              `bson:"acknowledgement"`
	Edited                    FlexibleBool              `bson:"edited"`
	Enabled                   FlexibleBool              `bson:"enabled"`
	IsDeleted                 FlexibleBool              `bson:"isdeleted"`
	IsTypeIsAlarm             FlexibleBool              `bson:"isTypeIsAlarm"`
	Devices                   FlexibleStringSlice       `bson:"devices"`
	Levels                    []AlarmLevel              `bson:"levels"`
	RuleSetAndOrOperationData RuleSetAndOrOperationData `bson:"ruleSetAndOrOperationData"`
	RuleSets                  []RuleSet                 `bson:"ruleSets"`
}

type BillMasterData struct {
	Customers           FlexibleFloat64   `bson:"customers"`
	DisconnectionDays   FlexibleFloat64   `bson:"disconnection_days"`
	ElectricalDuties    FlexibleStringSlice `bson:"electrical_duties"`
	ExcelFile           string            `bson:"excel_file"`
	FixedCostDeduction  []FlexibleFloat64 `bson:"fixed_cost_deduction"`
	FromDate            string            `bson:"from_date"`
	ToDate              string            `bson:"to_date"`
	SelectTods          FlexibleStringSlice `bson:"select_tods"`
	AddDeduction        []FlexibleFloat64 `bson:"add_deduction"`
	BillCycle           FlexibleFloat64   `bson:"bill_cycle"`
	BillDueDays         FlexibleFloat64   `bson:"bill_due_days"`
	IsTax               FlexibleFloat64   `bson:"is_tax"`
	StartTime           string            `bson:"start_time"`
}

type BillMaster struct {
	ID                string            `bson:"_id"`
	ClientID          string            `bson:"client_id"`
	CustomerName      string            `bson:"customer_name"`
	CustomerAddress   string            `bson:"customer_address"`
	CustomerID        FlexibleFloat64   `bson:"customer_id"`
	CustomerPh        FlexibleString    `bson:"customer_ph"`
	CustomerType      FlexibleFloat64   `bson:"customer_type"`
	DeviceInstanceID  string            `bson:"device_instance_id"`
	GatewayInstanceID string            `bson:"gateway_instance_id"`
	SiteID            string            `bson:"site_id"`
	EbTag             string            `bson:"eb_tag"`
	DgTag             string            `bson:"dg_tag"`
	FromDate          string            `bson:"from_date"`
	StartTime         FlexibleFloat64   `bson:"start_time"`
	FixedChargeDg     FlexibleFloat64   `bson:"fixed_charge_dg"`
	FixedChargeEb     *FlexibleFloat64  `bson:"fixed_charge_eb"` // can be null
	OtherCharges      FlexibleFloat64   `bson:"other_charges"`
	IsDeleted         string            `bson:"isdeleted"`
	IsDisabled        string            `bson:"isdisabled"`
	FlatNo            FlexibleString    `bson:"flat_no"`
	SqurArea          FlexibleString    `bson:"squr_area"`
	Data              BillMasterData    `bson:"data"`
}

type EmailGateway struct {
	ID                  string          `bson:"_id"`
	EmailGatewayID      string          `bson:"email_gateway_id"`
	EmailType           string          `bson:"email_type"`
	Encryption          string          `bson:"encryption"`
	MailServer          string          `bson:"mail_server"`
	Username            string          `bson:"username"`
	Password            string          `bson:"password"`
	SenderEmail         string          `bson:"sender_email"`
	SenderName          string          `bson:"sender_name"`
	SMTPPort            FlexibleFloat64 `bson:"smtp_port"`
	SMTPAuthorization   FlexibleBool    `bson:"smtp_authorization"`
	ProfileName         string          `bson:"profile_name"`
	ConfigurationStatus string          `bson:"configuration_status"`
	SiteID              string          `bson:"site_id"`
	ClientID            string          `bson:"client_id"`
	EventUserID         string          `bson:"event_user_id"`
	Default             FlexibleBool    `bson:"default"`
	SystemDefaults      FlexibleBool    `bson:"system_defaults"`
	IsDeleted           FlexibleBool    `bson:"isdeleted"`
}

type WorkGroup struct {
	ID                   string       `bson:"_id"`
	WorkGroupID          string       `bson:"work_group_id"`
	WorkGroupName        string       `bson:"workGroupName"`
	WorkGroupDescription string       `bson:"workGroupDescription"`
	SiteID               string       `bson:"site_id"`
	ClientID             string       `bson:"client_id"`
	EventUserID          string       `bson:"event_user_id"`
	Default              FlexibleBool `bson:"default"`
	IsDeleted            FlexibleBool `bson:"isdeleted"`
}

type SmartMeterSensorID struct {
	SensorID   string `bson:"sensor_id"`
	SensorName string `bson:"sensor_name"`
}

type SmartMeterGeneralInfoBody struct {
	Area                       FlexibleFloat64      `bson:"area"`
	AvenueNo                   FlexibleString       `bson:"avenue_no"`
	BalanceType                FlexibleFloat64      `bson:"balance_type"`
	BillType                   string               `bson:"bill_type"`
	BlockNumber                FlexibleString       `bson:"block_number"`
	BlockType                  string               `bson:"block_type"`
	CostManagementID           string               `bson:"cost_management_id"`
	CustomerAddress            string               `bson:"customer_address"`
	CustomerID                 FlexibleFloat64      `bson:"customer_id"`
	CustomerName               string               `bson:"customer_name"`
	DgTag                      string               `bson:"dg_tag"`
	DgThresholdKva             FlexibleFloat64      `bson:"dg_threshold_kva"`
	EbTag                      string               `bson:"eb_tag"`
	EmailID                    string               `bson:"email_id"`
	FixedChargeDg              FlexibleFloat64      `bson:"fixed_charge_dg"`
	FixedChargeEb              FlexibleFloat64      `bson:"fixed_charge_eb"`
	Flat                       string               `bson:"flat"`
	FlatNo                     FlexibleString       `bson:"flat_no"`
	GSTApplicable              FlexibleStringSlice  `bson:"gst_applicable"`
	GSTNumber                  string               `bson:"gst_number"`
	IsDeleted                  FlexibleStringSlice  `bson:"isdeleted"`
	IsDisabled                 FlexibleStringSlice  `bson:"isdisabled"`
	MeterModel                 string               `bson:"meter_model"`
	MeterType                  string               `bson:"meter_type"`
	MobileNo                   FlexibleFloat64      `bson:"mobile_no"`
	OccupancyDate              FlexibleString       `bson:"occupancy_date"`
	OtherCharges               FlexibleFloat64      `bson:"other_charges"`
	PlanType                   string               `bson:"plan_type"`
	RelayStatus                string               `bson:"relay_status"`
	SensorID                   []SmartMeterSensorID `bson:"sensor_id"`
	ServiceProvider            string               `bson:"service_provider"`
	ServiceProviderID          string               `bson:"service_provider_id"`
	SmartMeterConfigurationID  string               `bson:"smart_meter_configuration_id"`
	SubscriberType             FlexibleFloat64      `bson:"subscriber_type"`
	Tag1624                    FlexibleFloat64      `bson:"tag_1624"`
	Tag1625                    FlexibleFloat64      `bson:"tag_1625"`
	Tag1626                    FlexibleFloat64      `bson:"tag_1626"`
	Tag1627                    FlexibleFloat64      `bson:"tag_1627"`
	Tag1628                    FlexibleFloat64      `bson:"tag_1628"`
	Tag1629                    FlexibleFloat64      `bson:"tag_1629"`
	Tag1630                    FlexibleFloat64      `bson:"tag_1630"`
	Tag1631                    FlexibleFloat64      `bson:"tag_1631"`
	Tag2952                    FlexibleFloat64      `bson:"tag_2952"`
	Tag3046                    FlexibleFloat64      `bson:"tag_3046"`
	Tag3047                    FlexibleFloat64      `bson:"tag_3047"`
	TariffCateg                string               `bson:"tariff_categ"`
	TariffCategoryID           FlexibleStringSlice  `bson:"tariff_category_id"`
	Tower                      string               `bson:"tower"`
	VacancyDate                FlexibleString       `bson:"vacancy_date"`
}

type SmartMeterGeneralInfo struct {
	BodyContent SmartMeterGeneralInfoBody `bson:"bodyContent"`
}

type SmartMeterSlabSettingBody struct {
	BlockNumber    string               `bson:"block_number"`
	BlockType      string               `bson:"block_type"`
	SensorID       []SmartMeterSensorID `bson:"sensor_id"`
	Tag1655        FlexibleFloat64      `bson:"tag_1655"`
	Tag1657        FlexibleFloat64      `bson:"tag_1657"`
	Tag1673        FlexibleFloat64      `bson:"tag_1673"`
	Tag1674        FlexibleFloat64      `bson:"tag_1674"`
	Tag1675        FlexibleFloat64      `bson:"tag_1675"`
	Tag1676        FlexibleFloat64      `bson:"tag_1676"`
	Tag1677        FlexibleFloat64      `bson:"tag_1677"`
	Tag1678        FlexibleFloat64      `bson:"tag_1678"`
	Tag1681        FlexibleFloat64      `bson:"tag_1681"`
	Tag1682        FlexibleFloat64      `bson:"tag_1682"`
	Tag1683        FlexibleFloat64      `bson:"tag_1683"`
	Tag2900        FlexibleFloat64      `bson:"tag_2900"`
	Tag2901        FlexibleFloat64      `bson:"tag_2901"`
	Tag2902        string               `bson:"tag_2902"`
	Tag2903        string               `bson:"tag_2903"`
	Tag2904        FlexibleFloat64      `bson:"tag_2904"`
	Tag2905        FlexibleFloat64      `bson:"tag_2905"`
	Tag2906        FlexibleFloat64      `bson:"tag_2906"`
	Tag2907        FlexibleFloat64      `bson:"tag_2907"`
	Tag2908        FlexibleFloat64      `bson:"tag_2908"`
	Tag2909        FlexibleFloat64      `bson:"tag_2909"`
	Tag2910        FlexibleFloat64      `bson:"tag_2910"`
	Tag2911        FlexibleFloat64      `bson:"tag_2911"`
	Tag2912        FlexibleFloat64      `bson:"tag_2912"`
	Tag2913        FlexibleFloat64      `bson:"tag_2913"`
	Tag2914        FlexibleFloat64      `bson:"tag_2914"`
	Tag2915        FlexibleFloat64      `bson:"tag_2915"`
	Tag2916        FlexibleFloat64      `bson:"tag_2916"`
	Tag2917        FlexibleFloat64      `bson:"tag_2917"`
	Tag2918        FlexibleFloat64      `bson:"tag_2918"`
	Tag2919        FlexibleFloat64      `bson:"tag_2919"`
	Tag2920        FlexibleFloat64      `bson:"tag_2920"`
	Tag2921        FlexibleFloat64      `bson:"tag_2921"`
	Tag2922        FlexibleFloat64      `bson:"tag_2922"`
	Tag2923        FlexibleFloat64      `bson:"tag_2923"`
	Tag2924        FlexibleFloat64      `bson:"tag_2924"`
	Tag2925        FlexibleFloat64      `bson:"tag_2925"`
	Tag2926        FlexibleFloat64      `bson:"tag_2926"`
	Tag2927        FlexibleFloat64      `bson:"tag_2927"`
	Tag2928        FlexibleFloat64      `bson:"tag_2928"`
	Tag2929        FlexibleFloat64      `bson:"tag_2929"`
	Tag2930        FlexibleFloat64      `bson:"tag_2930"`
	Tag2931        FlexibleFloat64      `bson:"tag_2931"`
	Tag2932        FlexibleFloat64      `bson:"tag_2932"`
	Tag2933        FlexibleFloat64      `bson:"tag_2933"`
	Tag2934        FlexibleFloat64      `bson:"tag_2934"`
	Tag2935        FlexibleFloat64      `bson:"tag_2935"`
	Tag2936        FlexibleFloat64      `bson:"tag_2936"`
	Tag2952        FlexibleFloat64      `bson:"tag_2952"`
	Tag2973        FlexibleFloat64      `bson:"tag_2973"`
	Tag3054        FlexibleFloat64      `bson:"tag_3054"`
	Tag3521        FlexibleFloat64      `bson:"tag_3521"`
}

type SmartMeterSlabSetting struct {
	BodyContent SmartMeterSlabSettingBody `bson:"bodyContent"`
}

type SmartMeterConfiguration struct {
	ID               string                 `bson:"_id"`
	ClientID         string                 `bson:"client_id"`
	DeviceInstanceID string                 `bson:"device_instance_id"`
	SensorID         string                 `bson:"sensor_id"`
	SiteID           string                 `bson:"site_id"`
	IsDeleted        string                 `bson:"isdeleted"`
	IsDisabled       string                 `bson:"isdisabled"`
	GeneralInfo      SmartMeterGeneralInfo  `bson:"general_info"`
	SlabSetting      SmartMeterSlabSetting  `bson:"slab_setting"`
}
