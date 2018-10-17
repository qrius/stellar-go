// stellar_payment.go

package main

import (
	"log"
	"net/http" //go get github.com/stellar/go/keypair

	"github.com/stellar/go/build"
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
func sendLumens(amount string, sourceSeed string, destinationAddress string) {
	tx, err := build.Transaction(
		build.SourceAccount{sourceSeed},
		build.TestNetwork,
		build.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		build.Payment(
			build.Destination{AddressOrSeed: destinationAddress},
			build.NativeAmount{Amount: amount},
		),
	)

	if err != nil {
		panic(err)
	}

	txe, err := tx.Sign(sourceSeed)
	if err != nil {
		panic(err)
	}

	txeB64, err := txe.Base64()
	if err != nil {
		panic(err)
	}

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(txeB64)
	if err != nil {
		panic(err)
	}

	log.Println("Successfully sent", amount, "lumens to", destinationAddress, ". Hash:", resp.Hash)
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

	sendLumens("100", sourcePair.Seed(), destinationPair.Address())

	logBalances(addresses)
}
