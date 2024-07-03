package main

import (
	"fmt"
	"strings"
)

func EmailInfo(email string) string {
	// get domain
	domainStart := strings.LastIndex(email, "@") + 1
	domainEnd := strings.Index(email[domainStart:], ".")

	// get tld
	tldStart := domainEnd + domainStart + 1
	tldEnd := len(email)

	// get domain
	domain := email[domainStart : domainEnd+domainStart]

	// get tld
	tld := email[tldStart:tldEnd]

	// return result
	return fmt.Sprintf("Domain: %s dan TLD: %s", domain, tld)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
	fmt.Println(EmailInfo("ptmencaricintasejati@gmail.co.id"))
}
