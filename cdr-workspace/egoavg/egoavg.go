package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// CDRFile, CdrFileHeader, CdrHeader, CdrHdrTimeStamp 등의 타입 정의

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

type CdrHdrTimeStamp struct {
	MonthLocal                            uint8
	DateLocal                             uint8
	HourLocal                             uint8
	MinuteLocal                           uint8
	SignOfTheLocalTimeDifferentialFromUtc uint8 // bit set to "1" expresses "+" or bit set to "0" expresses "-" time deviation)
	HourDeviation                         uint8
	MinuteDeviation                       uint8
}

type FileClosureTriggerReasonType uint8

const (
	NormalClosure                     FileClosureTriggerReasonType = 0
	FileSizeLimitReached              FileClosureTriggerReasonType = 1
	FileOpentimeLimitedReached        FileClosureTriggerReasonType = 2
	MaximumNumberOfCdrsInFileReached  FileClosureTriggerReasonType = 3
	FileClosedByManualIntervention    FileClosureTriggerReasonType = 4
	CdrReleaseVersionOrEncodingChange FileClosureTriggerReasonType = 5
	AbnormalFileClosure               FileClosureTriggerReasonType = 128
	FileSystemError                   FileClosureTriggerReasonType = 129
	FileSystemStorageExhausted        FileClosureTriggerReasonType = 130
	FileIntegrityError                FileClosureTriggerReasonType = 131
)

type ReleaseIdentifierType uint8

const (
	Rel99 ReleaseIdentifierType = iota
	Rel4
	Rel5
	Rel6
	Rel7
	Rel8
	Rel9
	BeyondRel9
)

type DataRecordFormatType uint8

const (
	BasicEncodingRules DataRecordFormatType = iota + 1
	UnalignedPackedEncodingRules
	AlignedPackedEncodingRules1
	XMLEncodingRules
)

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

// CDRFile에 대한 Encoding 함수 추가
func (c *CDRFile) Encoding(fileName string) error {
	fileData, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, fileData, 0644)
}

// CDRFile에 대한 Decoding 함수 추가
func (c *CDRFile) Decoding(fileName string) error {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, c)
}

func performEncodingDecodingTasks() {
	// 테스트 케이스 정의
	cdrFile1 := CDRFile{
		Hdr: CdrFileHeader{
			FileLength:                            71,
			HeaderLength:                          63,
			HighReleaseIdentifier:                 2,
			HighVersionIdentifier:                 3,
			LowReleaseIdentifier:                  4,
			LowVersionIdentifier:                  5,
			FileOpeningTimestamp:                  CdrHdrTimeStamp{4, 28, 17, 18, 1, 8, 0},
			TimestampWhenLastCdrWasAppendedToFIle: CdrHdrTimeStamp{1, 2, 3, 4, 1, 6, 30},
			NumberOfCdrsInFile:                    1,
			FileSequenceNumber:                    11,
			FileClosureTriggerReason:              4,
			IpAddressOfNodeThatGeneratedFile:      [20]byte{0xa, 0xb, 0xa, 0xb, 0xa, 0xb, 0xa, 0xb, 0xa, 0xb, 0xa, 0xb, 0xa, 0xb, 0xa, 0xb, 0xa, 0xb, 0xa, 0xb},
			LostCdrIndicator:                      4,
			LengthOfCdrRouteingFilter:             4,
			CDRRouteingFilter:                     []byte("abcd"),
			LengthOfPrivateExtension:              5,
			PrivateExtension:                      []byte("fghjk"), // vendor specific
			HighReleaseIdentifierExtension:        2,
			LowReleaseIdentifierExtension:         3,
		},
		CdrList: []CDR{{
			Hdr: CdrHeader{
				CdrLength:                  3,
				ReleaseIdentifier:          Rel6,                         // octet 3 bit 6..8
				VersionIdentifier:          3,                            // otcet 3 bit 1..5
				DataRecordFormat:           UnalignedPackedEncodingRules, // octet 4 bit 6..8
				TsNumber:                   TS32253,                      // octet 4 bit 1..5
				ReleaseIdentifierExtension: 4,
			},
			CdrByte: []byte("abc"),
		}},
	}
	cdrFile2 := CDRFile{
		Hdr: CdrFileHeader{
			FileLength:                            92,
			HeaderLength:                          66,
			HighReleaseIdentifier:                 4,
			HighVersionIdentifier:                 5,
			LowReleaseIdentifier:                  5,
			LowVersionIdentifier:                  6,
			FileOpeningTimestamp:                  CdrHdrTimeStamp{1, 2, 11, 56, 1, 7, 30},
			TimestampWhenLastCdrWasAppendedToFIle: CdrHdrTimeStamp{4, 3, 2, 1, 0, 4, 0},
			NumberOfCdrsInFile:                    3,
			FileSequenceNumber:                    65,
			FileClosureTriggerReason:              2,
			IpAddressOfNodeThatGeneratedFile:      [20]byte{0xc, 0xd, 0xc, 0xd, 0xc, 0xd, 0xc, 0xd, 0xc, 0xd, 0xc, 0xd, 0xc, 0xd, 0xc, 0xd, 0xc, 0xd, 0xc, 0xd},
			LostCdrIndicator:                      4,
			LengthOfCdrRouteingFilter:             5,
			CDRRouteingFilter:                     []byte("gfdss"),
			LengthOfPrivateExtension:              7,
			PrivateExtension:                      []byte("abcdefg"), // vendor specific
			HighReleaseIdentifierExtension:        1,
			LowReleaseIdentifierExtension:         2,
		},
		CdrList: []CDR{
			{
				Hdr: CdrHeader{
					CdrLength:                  3,
					ReleaseIdentifier:          Rel6,
					VersionIdentifier:          3,
					DataRecordFormat:           UnalignedPackedEncodingRules,
					TsNumber:                   TS32253,
					ReleaseIdentifierExtension: 4,
				},
				CdrByte: []byte("abc"),
			},
			{
				Hdr: CdrHeader{
					CdrLength:                  6,
					ReleaseIdentifier:          Rel5,
					VersionIdentifier:          2,
					DataRecordFormat:           AlignedPackedEncodingRules1,
					TsNumber:                   TS32205,
					ReleaseIdentifierExtension: 2,
				},
				CdrByte: []byte("ghjklm"),
			},
			{
				Hdr: CdrHeader{
					CdrLength:                  2,
					ReleaseIdentifier:          Rel9,
					VersionIdentifier:          3,
					DataRecordFormat:           AlignedPackedEncodingRules1,
					TsNumber:                   TS32225,
					ReleaseIdentifierExtension: 1,
				},
				CdrByte: []byte("cv"),
			},
		},
	}

	fileName1 := "encoding0.txt"
	cdrFile1.Encoding(fileName1)
	newCdrFile1 := CDRFile{}
	newCdrFile1.Decoding(fileName1)
	e1 := os.Remove(fileName1)
	if e1 != nil {
		fmt.Println(e1)
	}

	fileName2 := "encoding1.txt"
	cdrFile2.Encoding(fileName2)
	newCdrFile2 := CDRFile{}
	newCdrFile2.Decoding(fileName2)
	e2 := os.Remove(fileName2)
	if e2 != nil {
		fmt.Println(e2)
	}

	fileName := "temp_encoding_decoding.txt"
	cdrFile := CDRFile{} // 실제 사용 시, CDRFile 구조체를 적절하게 초기화해야 합니다.

	cdrFile.Encoding(fileName)
	newCdrFile := CDRFile{}
	newCdrFile.Decoding(fileName)
	os.Remove(fileName)
}


func main() {
	const iterations = 5
	var totalExecutionTime time.Duration

	for i := 0; i < iterations; i++ {
		startTime := time.Now()

		performEncodingDecodingTasks()

		executionTime := time.Since(startTime)
		totalExecutionTime += executionTime
		fmt.Printf("Iteration %d Execution Time: %s\n", i+1, executionTime)
	}

	averageExecutionTime := totalExecutionTime / iterations
	fmt.Printf("Average Execution Time: %s\n", averageExecutionTime)
}
