// stellar_payment.go

package main

import (
	"log"
	"net/http" //go get github.com/stellar/go/keypair

	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
)

func fillAccounts(addresses [2]string) {
	for _, address := range addresses {
		friendBotResp, err := http.Get("https://horizon-testnet.stellar.org/friendbot?addr=" + address)
		if err != nil {
			log.Fatal(err)
		}
		defer friendBotResp.Body.Close()
	}
}

func logBalances(addresses [2]string) {
	for _, address := range addresses {
		account, err := horizon.DefaultTestNetClient.LoadAccount(address)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Balances for address:", address)
		for _, balance := range account.Balances {
			log.Println(balance)
		}
	}
}

func main() {
	sourcePair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}

	destinationPair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(sourcePair.Address())
	log.Println(destinationPair.Address())

	addresses := [2]string{sourcePair.Address(), destinationPair.Address()}

	fillAccounts(addresses)
	logBalances(addresses)
}
