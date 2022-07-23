// Code generated by download. DO NOT EDIT.

package iso20022_reda_005_001_03

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AdditionalReference10 struct {
	Ref     Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Ref"`
	RefIssr PartyIdentification139 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 RefIssr,omitempty"`
	MsgNm   Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 MsgNm,omitempty"`
}

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type AnyBICDec2014Identifier string

// Must match the pattern [A-Z]{2,2}
type CountryCode string

// May be one of DIST, ACCU
type DistributionPolicy1Code string

type Document struct {
	InvstmtFndRptReq InvestmentFundReportRequestV03 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 InvstmtFndRptReq"`
}

// May be no more than 4 items long
type ExternalFinancialInstrumentIdentificationType1Code string

type FinancialInstrument71 struct {
	Id          SecurityIdentification19 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Id"`
	ShrtNm      Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 ShrtNm,omitempty"`
	Nm          Max350Text               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Nm,omitempty"`
	SplmtryId   Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 SplmtryId,omitempty"`
	ClssTp      Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 ClssTp,omitempty"`
	SctiesForm  FormOfSecurity1Code      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 SctiesForm,omitempty"`
	DstrbtnPlcy DistributionPolicy1Code  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 DstrbtnPlcy,omitempty"`
	PdctGrp     Max140Text               `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 PdctGrp,omitempty"`
}

// May be one of BEAR, REGD
type FormOfSecurity1Code string

type FundParameters4Choice struct {
	NoCrit NoCriteria1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 NoCrit,omitempty"`
	Params FundParameters5 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Params,omitempty"`
}

type FundParameters5 struct {
	FinInstrmDtls   []FinancialInstrument71  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 FinInstrmDtls,omitempty"`
	FndMgmtCpny     []PartyIdentification139 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 FndMgmtCpny,omitempty"`
	DtFr            ISODate                  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 DtFr,omitempty"`
	CtryOfDmcl      CountryCode              `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 CtryOfDmcl,omitempty"`
	RegdDstrbtnCtry []CountryCode            `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 RegdDstrbtnCtry,omitempty"`
}

type GenericIdentification1 struct {
	Id      Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Id"`
	SchmeNm Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 SchmeNm,omitempty"`
	Issr    Max35Text `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Issr,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[A-Z0-9]{9,9}[0-9]{1,1}
type ISINOct2015Identifier string

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

type IdentificationSource3Choice struct {
	Cd    ExternalFinancialInstrumentIdentificationType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Cd,omitempty"`
	Prtry Max35Text                                          `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Prtry,omitempty"`
}

type InvestmentFundReportRequestV03 struct {
	MsgId   MessageIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 MsgId"`
	PrvsRef AdditionalReference10   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 PrvsRef,omitempty"`
	RltdRef AdditionalReference10   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 RltdRef,omitempty"`
	RptReq  []FundParameters4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 RptReq"`
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

// May be no more than 140 items long
type Max140Text string

// May be no more than 16 items long
type Max16Text string

// May be no more than 350 items long
type Max350Text string

// May be no more than 35 items long
type Max35Text string

// May be no more than 70 items long
type Max70Text string

type MessageIdentification1 struct {
	Id      Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Id"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 CreDtTm"`
}

type NameAndAddress5 struct {
	Nm  Max350Text     `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Nm"`
	Adr PostalAddress1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Adr,omitempty"`
}

// May be one of NOCR
type NoCriteria1Code string

type OtherIdentification1 struct {
	Id  Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Id"`
	Sfx Max16Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Sfx,omitempty"`
	Tp  IdentificationSource3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Tp"`
}

type PartyIdentification125Choice struct {
	AnyBIC   AnyBICDec2014Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 AnyBIC,omitempty"`
	PrtryId  GenericIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 PrtryId,omitempty"`
	NmAndAdr NameAndAddress5         `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 NmAndAdr,omitempty"`
}

type PartyIdentification139 struct {
	Pty PartyIdentification125Choice `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Pty"`
	LEI LEIIdentifier                `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 LEI,omitempty"`
}

type PostalAddress1 struct {
	AdrTp       AddressType2Code `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 AdrTp,omitempty"`
	AdrLine     []Max70Text      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 AdrLine,omitempty"`
	StrtNm      Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 StrtNm,omitempty"`
	BldgNb      Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 BldgNb,omitempty"`
	PstCd       Max16Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 PstCd,omitempty"`
	TwnNm       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 TwnNm,omitempty"`
	CtrySubDvsn Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 CtrySubDvsn,omitempty"`
	Ctry        CountryCode      `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Ctry"`
}

type SecurityIdentification19 struct {
	ISIN   ISINOct2015Identifier  `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 ISIN,omitempty"`
	OthrId []OtherIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 OthrId,omitempty"`
	Desc   Max140Text             `xml:"urn:iso:std:iso:20022:tech:xsd:reda.005.001.03 Desc,omitempty"`
}

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}