package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/free5gc/chf/cdr/cdrFile"
	//"sgx-workspace/ego/CDRFile"

        //"cdrfile/cdrFile"
	//"sgx-workspace/ego/CDRFile/cdrfile"
        //"github.com/stretchr/testify/require"
	// 필요한 추가 import들 (예: "bytes", "encoding/binary" 등)
)

// CDRFile, CdrFileHeader, CdrHeader, CdrHdrTimeStamp 등의 타입 정의
// ...

func main() {
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
			LostCdrIndicator:          4,
			LengthOfCdrRouteingFilter: 4,
			CDRRouteingFilter:                     []byte("abcd"),
			LengthOfPrivateExtension: 5,
			PrivateExtension:                      []byte("fghjk"), // vendor specific
			HighReleaseIdentifierExtension: 2,
			LowReleaseIdentifierExtension:  3,
		},
		CdrList: []CDR{{
			Hdr:CdrHeader {
				CdrLength                  :3,
				ReleaseIdentifier          :Rel6, // octet 3 bit 6..8
				VersionIdentifier          :3,                // otcet 3 bit 1..5
				DataRecordFormat           :UnalignedPackedEncodingRules,  // octet 4 bit 6..8
				TsNumber                   : TS32253,   // octet 4 bit 1..5
				ReleaseIdentifierExtension :4,
			},
			CdrByte:[]byte("abc"),
		},},
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
			LostCdrIndicator:          			   4,
			LengthOfCdrRouteingFilter: 			   5,
			CDRRouteingFilter:                     []byte("gfdss"),
			LengthOfPrivateExtension:    	       7,
			PrivateExtension:                      []byte("abcdefg"), // vendor specific
			HighReleaseIdentifierExtension: 	   1,
			LowReleaseIdentifierExtension:  	   2,
		},
		CdrList: []CDR{
			{
				Hdr:CdrHeader {
					CdrLength                  :3,
					ReleaseIdentifier          :Rel6, 
					VersionIdentifier          :3,            
					DataRecordFormat           :UnalignedPackedEncodingRules,  
					TsNumber                   : TS32253,   
					ReleaseIdentifierExtension :4,
				},
				CdrByte:[]byte("abc"),
			},
			{
				Hdr:CdrHeader {
					CdrLength                  :6,
					ReleaseIdentifier          :Rel5,
					VersionIdentifier          :2,               
					DataRecordFormat           :AlignedPackedEncodingRules1,  
					TsNumber                   : TS32205,   
					ReleaseIdentifierExtension :2,
				},
				CdrByte:[]byte("ghjklm"),
			},
			{
				Hdr:CdrHeader {
					CdrLength                  :2,
					ReleaseIdentifier          :Rel9,
					VersionIdentifier          :3,               
					DataRecordFormat           :AlignedPackedEncodingRules1,  
					TsNumber                   : TS32225,  
					ReleaseIdentifierExtension :1,
				},
				CdrByte:[]byte("cv"),
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
}

