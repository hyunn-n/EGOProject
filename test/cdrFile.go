package cdrFile

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

type CDRFile struct {
	Hdr     CdrFileHeader
	CdrList []CDR
}

type CDR struct {
	Hdr     CdrHeader
	CdrByte []byte
}

type CdrFileHeader struct {
	FileLength                            uint32
	HeaderLength                          uint32
	HighReleaseIdentifier                 uint8 // octet 9 bit 6..8
	HighVersionIdentifier                 uint8 // octet 9 bit 1..5
	LowReleaseIdentifier                  uint8 // octet 10 bit 6..8
	LowVersionIdentifier                  uint8 // octet 10 bit 1..5
	FileOpeningTimestamp                  CdrHdrTimeStamp
	TimestampWhenLastCdrWasAppendedToFIle CdrHdrTimeStamp
	NumberOfCdrsInFile                    uint32
	FileSequenceNumber                    uint32
	FileClosureTriggerReason              FileClosureTriggerReasonType
	IpAddressOfNodeThatGeneratedFile      [20]byte // ip address in ipv6 format
	LostCdrIndicator                      uint8
	LengthOfCdrRouteingFilter             uint16
	CDRRouteingFilter                     []byte // vendor specific
	LengthOfPrivateExtension              uint16
	PrivateExtension                      []byte // vendor specific
	HighReleaseIdentifierExtension        uint8
	LowReleaseIdentifierExtension         uint8
}

type CdrHeader struct {
	CdrLength                  uint16
	ReleaseIdentifier          ReleaseIdentifierType // octet 3 bit 6..8
	VersionIdentifier          uint8                 // otcet 3 bit 1..5
	DataRecordFormat           DataRecordFormatType  // octet 4 bit 6..8
	TsNumber                   TsNumberIdentifier    // octet 4 bit 1..5
	ReleaseIdentifierExtension uint8
}

type DataRecordFormatType uint8


type TsNumberIdentifier uint8

const (
	TS32005 TsNumberIdentifier = 0
	TS32015 TsNumberIdentifier = 1
	TS32205 TsNumberIdentifier = 2
	TS32215 TsNumberIdentifier = 3
	TS32225 TsNumberIdentifier = 4
	TS32235 TsNumberIdentifier = 5
	TS32250 TsNumberIdentifier = 6
	TS32251 TsNumberIdentifier = 7
	TS32260 TsNumberIdentifier = 9
	TS32270 TsNumberIdentifier = 10
	TS32271 TsNumberIdentifier = 11
	TS32272 TsNumberIdentifier = 12
	TS32273 TsNumberIdentifier = 13
	TS32275 TsNumberIdentifier = 14
	TS32274 TsNumberIdentifier = 15
	TS32277 TsNumberIdentifier = 16
	TS32296 TsNumberIdentifier = 17
	TS32278 TsNumberIdentifier = 18
	TS32253 TsNumberIdentifier = 19
	TS32255 TsNumberIdentifier = 20
	TS32254 TsNumberIdentifier = 21
	TS32256 TsNumberIdentifier = 22
	TS28201 TsNumberIdentifier = 23
	TS28202 TsNumberIdentifier = 24
)


// 나머지 코드 유지

func (cdrf CdrFileHeader) Encoding() []byte {
	buf := new(bytes.Buffer)

	// File length
	if err := binary.Write(buf, binary.BigEndian, cdrf.FileLength); err != nil {
		fmt.Println("CdrFileHeader File length failed:", err)
	}

	// Header length
	if err := binary.Write(buf, binary.BigEndian, cdrf.HeaderLength); err != nil {
		fmt.Println("CdrFileHeader Header length failed:", err)
	}

	// High release / version identifier
	var highIdentifier uint8 = (cdrf.HighReleaseIdentifier << 5) | cdrf.HighVersionIdentifier
	if err := binary.Write(buf, binary.BigEndian, highIdentifier); err != nil {
		fmt.Println("CdrFileHeader highIdentifier failed:", err)
	}

	// Low release / version identifier
	var lowIdentifier uint8 = (cdrf.LowReleaseIdentifier << 5) | cdrf.LowVersionIdentifier
	if err := binary.Write(buf, binary.BigEndian, lowIdentifier); err != nil {
		fmt.Println("CdrFileHeader lowIdentifier failed:", err)
	}

	// File opening timestamp
	var ts uint32 = uint32(cdrf.FileOpeningTimestamp.MonthLocal)<<28 |
		uint32(cdrf.FileOpeningTimestamp.DateLocal)<<23 |
		uint32(cdrf.FileOpeningTimestamp.HourLocal)<<18 |
		uint32(cdrf.FileOpeningTimestamp.MinuteLocal)<<12 |
		uint32(cdrf.FileOpeningTimestamp.SignOfTheLocalTimeDifferentialFromUtc)<<11 |
		uint32(cdrf.FileOpeningTimestamp.HourDeviation)<<6 |
		uint32(cdrf.FileOpeningTimestamp.MinuteDeviation)
	if err := binary.Write(buf, binary.BigEndian, ts); err != nil {
		fmt.Println("CdrFileHeader File opening timestamp failed:", err)
	}

	// Timestamp when last CDR was appended to file
	ts = uint32(cdrf.TimestampWhenLastCdrWasAppendedToFIle.MonthLocal)<<28 |
		uint32(cdrf.TimestampWhenLastCdrWasAppendedToFIle.DateLocal)<<23 |
		uint32(cdrf.TimestampWhenLastCdrWasAppendedToFIle.HourLocal)<<18 |
		uint32(cdrf.TimestampWhenLastCdrWasAppendedToFIle.MinuteLocal)<<12 |
		uint32(cdrf.TimestampWhenLastCdrWasAppendedToFIle.SignOfTheLocalTimeDifferentialFromUtc)<<11 |
		uint32(cdrf.TimestampWhenLastCdrWasAppendedToFIle.HourDeviation)<<6 |
		uint32(cdrf.TimestampWhenLastCdrWasAppendedToFIle.MinuteDeviation)
	if err := binary.Write(buf, binary.BigEndian, ts); err != nil {
		fmt.Println("CdrFileHeader TimestampWhenLastCdrWasAppendedToFIle failed:", err)
	}

	// 나머지 필드들의 인코딩 코드 추가

	return buf.Bytes()
}

// 나머지 코드 유지

func (cdfFile CDRFile) Encoding(fileName string) {
	buf := new(bytes.Buffer)

	bufCdrFileHeader := cdfFile.Hdr.Encoding()
	if err := binary.Write(buf, binary.BigEndian, bufCdrFileHeader); err != nil {
		fmt.Println("CDRFile failed:", err)
	}

	for i, cdr := range cdfFile.CdrList {
		bufCdrHeader := cdr.Hdr.Encoding()
		if err := binary.Write(buf, binary.BigEndian, bufCdrHeader); err != nil {
			fmt.Println("CDRFile failed:", err)
		}

		if err := binary.Write(buf, binary.BigEndian, cdr.CdrByte); err != nil {
			fmt.Println("CDRFile failed:", err)
		}

		if len(cdr.CdrByte) != int(cdr.Hdr.CdrLength) {
			fmt.Println("[Encoding Warning]CdrLength field of cdr", i, "header not equals to the length of encoding cdr", i)
			fmt.Println("\tExpected", len(cdr.CdrByte), "Get", int(cdr.Hdr.CdrLength))
		}
	}

	if cdfFile.Hdr.FileLength != uint32(len(buf.Bytes())) && cdfFile.Hdr.FileLength != 0xffffffff {
		fmt.Println("[Encoding Warning]FileLength field of CdfFile Header not equals to the length of encoding file.")
		fmt.Println("\tExpected", uint32(len(buf.Bytes())), "Get", cdfFile.Hdr.FileLength)
	}

	err := ioutil.WriteFile(fileName, buf.Bytes(), 0666)
	if err != nil {
		panic(err)
	}
}

func (cdfFile *CDRFile) Decoding(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// 나머지 디코딩 코드 추가

	tail := n + 2

	for i := 1; i <= int(numberOfCdrsInFile); i++ {
		cdrLength := binary.BigEndian.Uint16(data[tail : tail+2])
		if len(data) < int(tail)+5+int(cdrLength) {
			fmt.Println("[Decoding Error]Length of cdrfile is wrong. cdr:", i)
		}

		cdrHeader := CdrHeader{
			CdrLength:                  cdrLength,
			ReleaseIdentifier:          ReleaseIdentifierType(data[tail+2] >> 5),
			VersionIdentifier:          data[tail+2] & 0b11111,
			DataRecordFormat:           DataRecordFormatType(data[tail+3] >> 5),
			TsNumber:                   TsNumberIdentifier(data[tail+3] & 0b11111),
			ReleaseIdentifierExtension: data[tail+4],
		}

		cdr := CDR{
			Hdr:     cdrHeader,
			CdrByte: data[tail+5 : tail+5+cdrLength],
		}
		cdfFile.CdrList = append(cdfFile.CdrList, cdr)
		tail += 5 + cdrLength
	}
}
