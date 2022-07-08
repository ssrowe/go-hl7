package hl7v23

import "time"

// Extended Composite ID With Check Digit
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CX
type CX struct {
	Id                                         string `hl7:"0.1"`
	CheckDigit                                 string `hl7:"0.2"`
	CodeIdentifyingTheCheckDigitSchemeEmployed HD     `hl7:"0.3"`
	AssigningAuthority                         string `hl7:"0.4"`
	AssigningFacility                          HD     `hl7:"0.5"`
}

// XPN - Extended Person Name
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XPN
type XPN struct {
	FamilyName          string `hl7:"0.1"`
	GivenName           string `hl7:"0.2"`
	MiddleInitialOrName string `hl7:"0.3"`
	Suffix              string `hl7:"0.4"`
	Prefix              string `hl7:"0.5"`
	Degree              string `hl7:"0.6"`
	NameTypeCode        string `hl7:"0.7"`
	NameRepresentation  string `hl7:"0.8"`
}

// XAD - Extended Address
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XAD
type XAD struct {
	StreetAddress              string `hl7:"0.1"`
	OtherDesignation           string `hl7:"0.2"`
	City                       string `hl7:"0.3"`
	StateOrProvince            string `hl7:"0.4"`
	ZipOrPostalCode            string `hl7:"0.5"`
	Country                    string `hl7:"0.6"`
	AddressType                string `hl7:"0.7"`
	OtherGeographicDesignation string `hl7:"0.8"`
	CountyCode                 string `hl7:"0.9"`
	CensusTract                string `hl7:"0.10"`
}

// XTN - Extended Telecommunication Number
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XTN
type XTN struct {
	TelephoneNumber                string `hl7:"0.1"`
	TelecommunicationUseCode       string `hl7:"0.2"`
	TelecommunicationEquipmentType string `hl7:"0.3"`
	EmailAddress                   string `hl7:"0.4"`
	CountryCode                    int    `hl7:"0.5"`
	AreaCode                       int    `hl7:"0.6"`
	PhoneNumber                    int    `hl7:"0.7"`
	Extension                      int    `hl7:"0.8"`
	AnyText                        string `hl7:"0.9"`
}

// CE - Coded Element
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CE
type CE struct {
	Identifier                  string `hl7:"0.1"`
	Text                        string `hl7:"0.2"`
	NameOFCodingSystem          string `hl7:"0.3"`
	AlternateIdentifier         string `hl7:"0.4"`
	AlternateText               string `hl7:"0.5"`
	NameOfAlternateCodingSystem string `hl7:"0.6"`
}

// DLN - Driver's License Number
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/DLN
type DLN struct {
	DriverLicenseNumber         string `hl7:"0.1"`
	IssuingStateProvinceCountry string `hl7:"0.2"`
	ExpirationDate              string `hl7:"0.3"`
}

// XON - Extended Composite Name And ID For Organizations
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XON
type XON struct {
}

// XCN - Extended Composite ID Number And Name
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/XCN
type XCN struct {
}

// HD - Hierarchic Designator
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/HD
type HD struct {
	NamespaceId     string `hl7:"0.1"`
	UniversalId     string `hl7:"0.2"`
	UniversalIdType string `hl7:"0.3"`
}

// CM_AUI - Authorization Information
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_AUI
type CM_AUI struct {
	AuthorizationNumber string    `hl7:"0.1"`
	Date                time.Time `hl7:"0.2,shortdate"`
	Source              string    `hl7:"0.3"`
}

// CM_RMC - Room Coverage
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_RMC
type CM_RMC struct {
	RoomType       string  `hl7:"0.1"`
	AmmountType    string  `hl7:"0.2"`
	CoverageAmount float32 `hl7:"0.3"`
}

// CM_PTA - Policy Type
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_PTA
type CM_PTA struct {
	PolicyType  string  `hl7:"0.1"`
	AmountClass string  `hl7:"0.2"`
	Amount      float32 `hl7:"0.3"`
}

// CM_DDI - Daily Deductible
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_DDI
type CM_DDI struct {
	DelayDays    float32 `hl7:"0.1"`
	Amount       float32 `hl7:"0.2"`
	NumberOfDays float32 `hl7:"0.3"`
}

// JCC - Job Code Class
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/JCC
type JCC struct {
	JobCode  string `hl7:"0.1"`
	JobClass string `hl7:"0.2"`
}

// HL7 v2.3 - CP - Composite Price
// https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CP
type CP struct {
	Price      float32 `hl7:"0.1"`
	PriceType  string  `hl7:"0.2"`
	FromValue  float32 `hl7:"0.3"`
	ToValue    float32 `hl7:"0.4"`
	RangeUnits CE      `hl7:"0.5"`
	RangeType  string  `hl7:"0.6"`
}

type Sex string

const (
	Female  Sex = "F"
	Male    Sex = "M"
	Other   Sex = "O"
	Unknown Sex = "U"
)

//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/PL
type PL struct {
	PointOfCare        string `hl7:"0.1"`
	Room               string `hl7:"0.2"`
	Bed                string `hl7:"0.3"`
	Facility           HD     `hl7:"0.4"`
	LocationStatus     string `hl7:"0.5"`
	PersonLocationType string `hl7:"0.6"`
	Building           string `hl7:"0.7"`
	Floor              string `hl7:"0.8"`
	LocationType       string `hl7:"0.9"`
}

//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_DLD
type CM_DLD struct {
	DischargeLocation string    `hl7:"0.1"`
	EffectiveDate     time.Time `hl7:"0.2"`
}

//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/FC
type FC struct {
	FinancialClass string    `hl7:"0.1"`
	EffectiveDate  time.Time `hl7:"0.2"`
}

//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_EIP
type CM_EIP struct {
	ParentsPlacerOrderNumber string `hl7:"0.1"`
	ParentsFillerOrderNumber string `hl7:"0.2"`
}

//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/TQ
type TQ struct {
	Quantity        CQ        `hl7:"0.1"`
	Interval        RI        `hl7:"0.2"`
	Duration        string    `hl7:"0.3"`
	StartDatetime   time.Time `hl7:"0.4"`
	EndDatetime     time.Time `hl7:"0.5"`
	Priority        string    `hl7:"0.6"`
	Condition       string    `hl7:"0.7"`
	Text            string    `hl7:"0.8"`
	Conjunction     string    `hl7:"0.9"`
	OrderSequencing CM_OSD    `hl7:"0.10"`
}

//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CQ
type CQ struct {
	Quantity int    `hl7:"0.1"`
	Units    string `hl7:"0.2"`
}

//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/RI
type RI struct {
	RepeatPattern        string `hl7:"0.1"`
	ExplicitTimeInterval string `hl7:"0.2"`
}

//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/CM_OSD
type CM_OSD struct {
	SequenceResultsFlag               string `hl7:"0.1"`
	PlacerOrderNumberEntityIdentifier string `hl7:"0.2"`
	PlacerOrderNumberNamespaceID      string `hl7:"0.3"`
	FillerOrderNumberEntityIdentifier string `hl7:"0.4"`
	FillerOrderNumberNamespaceID      string `hl7:"0.5"`
	SequenceConditionValue            string `hl7:"0.6"`
	MaximumNumberOfRepeats            int    `hl7:"0.7"`
	PlacerOrderNumberUniversalID      string `hl7:"0.8"`
	PlacerOrderNumberUniversalIDType  string `hl7:"0.9"`
	FillerOrderNumberUniversalID      string `hl7:"0.10"`
	FillerOrderNumberUniversalIDType  string `hl7:"0.11"`
}

//https://hl7-definition.caristix.com/v2/HL7v2.3/DataTypes/EI
type EI struct {
	EntityIdentifier string `hl7:"0.1"`
	NamespaceId      string `hl7:"0.2"`
	UniversalId      string `hl7:"0.3"`
	UniversalIdType  string `hl7:"0.4"`
}
