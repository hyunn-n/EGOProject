package main

import (
	"time"
	"chr"
)

func main() {
	chr.DoSomething()
	loc, _ := time.LoadLocation("Asia/Kolkata")
	timeNow := time.Now().In(loc)
	// timeNow := time.Now()
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
		//IpAddressOfNodeThatGeneratedFile      [20]byte(),
		LostCdrIndicator:          4,
		LengthOfCdrRouteingFilter: 4,
		CDRRouteingFilter:                     []byte("abcd"),
		LengthOfPrivateExtension: 5,
		PrivateExtension:                      []byte("fghjk"), // vendor specific
		HighReleaseIdentifierExtension: 2,
		LowReleaseIdentifierExtension:  3,
	}

	cdrHeader := CdrHeader {
		CdrLength                  :3,
		ReleaseIdentifier          :Rel6, // octet 3 bit 6..8
		VersionIdentifier          :3,                // otcet 3 bit 1..5
		DataRecordFormat           :UnalignedPackedEncodingRules,  // octet 4 bit 6..8
		TsNumber                   : TS32253,   // octet 4 bit 1..5
		ReleaseIdentifierExtension :4,
	}

	cdrFile := CDRFile{
		Hdr: cdrf,
		CdrList: []CDR{{Hdr:cdrHeader, CdrByte:[]byte("abc")},},
	}

	cdrFile.Encoding("encoding.txt")
	cdrFile = CDRFile{}
	cdrFile.Decoding("encoding.txt")
}

