package cdrFile

import (
	//"bytes"
	//"encoding/binary"
	//"fmt"
	//"io/ioutil"
	//"os"
	"time"
)

type CdrFileHeader struct {
	FileLength                            uint16
	HeaderLength                          uint16
	HighReleaseIdentifier                 uint8
	HighVersionIdentifier                 uint8
	LowReleaseIdentifier                  uint8
	LowVersionIdentifier                  uint8
	FileOpeningTimestamp                  *time.Time
	TimestampWhenLastCdrWasAppendedToFIle *time.Time
	NumberOfCdrsInFile                    uint32
	FileSequenceNumber                    uint16
	FileClosureTriggerReason              uint8
	// IpAddressOfNodeThatGeneratedFile [20]byte
	LostCdrIndicator          uint8
	LengthOfCdrRouteingFilter uint8
	CDRRouteingFilter         []byte
	LengthOfPrivateExtension  uint8
	PrivateExtension          []byte
	HighReleaseIdentifierExtension uint8
	LowReleaseIdentifierExtension  uint8
}

type CdrHeader struct {
	CdrLength                  uint8
	ReleaseIdentifier          uint8
	VersionIdentifier          uint8
	DataRecordFormat           uint8
	TsNumber                   uint8
	ReleaseIdentifierExtension uint8
}

type CDR struct {
	hdr     CdrHeader
	cdrByte []byte
}

type CDRFile struct {
	hdr     CdrFileHeader
	cdrList []CDR
}

func main() {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	timeNow := time.Now().In(loc)
	cdrf := CdrFileHeader{
		FileLength:                            5,
		HeaderLength:                          6,
		HighReleaseIdentifier:                 2,
		HighVersionIdentifier:                 3,
		LowReleaseIdentifier:                  4,
		LowVersionIdentifier:                  5,
		FileOpeningTimestamp:                  &timeNow,
		TimestampWhenLastCdrWasAppendedToFIle: &timeNow,
		NumberOfCdrsInFile:                    1,
		FileSequenceNumber:                    11,
		FileClosureTriggerReason:              4,
		LostCdrIndicator:                     4,
		LengthOfCdrRouteingFilter:            4,
		CDRRouteingFilter:                    []byte("abcd"),
		LengthOfPrivateExtension:              5,
		PrivateExtension:                      []byte("fghjk"),
		HighReleaseIdentifierExtension:        2,
		LowReleaseIdentifierExtension:         3,
	}

	cdrHeader := CdrHeader{
		CdrLength:                  3,
		ReleaseIdentifier:          6,
		VersionIdentifier:          3,
		DataRecordFormat:           6,
		TsNumber:                   53,
		ReleaseIdentifierExtension: 4,
	}

	cdrFile := CDRFile{
		hdr: cdrf,
		cdrList: []CDR{
			{hdr: cdrHeader, cdrByte: []byte("abc")},
		},
	}

	cdrFile.Encoding()
	cdrFile = CDRFile{}
	cdrFile.Decoding("encoding.txt")
}
