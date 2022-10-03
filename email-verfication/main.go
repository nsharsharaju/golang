package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	var domain string
	fmt.Scan(&domain)
	checkDomain(domain)
}

func checkDomain(domain string) {
	fmt.Println("domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord")

	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	//MX Record
	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMx = true
	}

	//hasSPF and spfRecord
	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	//fmt.Println(domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord)
	fmt.Println("domain:", domain)
	fmt.Println("hasMx:", hasMx)
	fmt.Println("hasSPF:", hasSPF)
	fmt.Println("spfRecord:", spfRecord)
	fmt.Println("hasDMARC:", hasDMARC)
	fmt.Println("dmarcRecord:", dmarcRecord)
}
