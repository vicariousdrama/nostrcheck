package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"github.com/nbd-wtf/go-nostr"
)


var (
	notejson string
	notefile string
)
func init() {
	flag.StringVar(&notejson, "notejson", `{"id":"","pubkey":"","created_at":0,"kind":1,"tags":[],"content":"","sig":""}`, "JSON of Note to check")
	flag.StringVar(&notefile, "notefile", "", "Filepath to JSON of a note to check")
}


func main() {

	fmt.Println()
	flag.Parse()

	var evt nostr.Event
	var bdata []byte

	// Check if notefile provided
	if len(notefile) > 0 {
		// Use notefile
		if fdata, err := os.ReadFile(notefile); err != nil {
			fmt.Println("Failed to load file: " + err.Error())
			os.Exit(1)
		} else {
			bdata = fdata
		}
	} else {
		// Use notejson
		bdata = []byte(notejson)
	}

	// Parse the event
	if err := json.Unmarshal(bdata, &evt); err != nil {
		fmt.Println("Failed to decode event: " + err.Error())
		os.Exit(1)
	}

	// Display the JSON
	if evt.ID == "8d600afa179801376a3c784bc79cec7753c768b8943911d3fb56fe29d63b221c" {
		fmt.Println("Using note posted by Mike Digler about Gossip as a sample")
	}
	fmt.Println("Formatted JSON of note event:")
	if b, err := json.MarshalIndent(evt, "", "  "); err != nil {
		fmt.Println("Failed marshalling json: " + err.Error())
		os.Exit(1)
	} else {
		fmt.Println(string(b))
	}

	// Get and reprot on the hash compared to previously assigned ID
	serialized := evt.Serialize()
	hash := sha256.Sum256(serialized)
	matchresult := "(does not match)"
	if (evt.ID == hex.EncodeToString(hash[:])) {
		matchresult = "(matched)"
	}
	fmt.Println("Expected hash: " + hex.EncodeToString(hash[:]) + " " + matchresult)

	// Check signature verification
	if ok, err := evt.CheckSignature(); err != nil {
		fmt.Println("Failed to verify signature: " + err.Error())
	} else if !ok {
		fmt.Println("Signature is invalid!")
	} else {
		// signature passed. report who its from
		fmt.Println("Signature is valid from pubkey: " + evt.PubKey)
	}
}
