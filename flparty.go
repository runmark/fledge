package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	dapr "github.com/dapr/go-sdk/client"
)

const BOARD_URL = "http://%s.board.kubefate.net:81/"
const DEFAULT_NAMESPACE = "default"
const DEFAULT_PARTYID = "0000"
const K8S_NAMESPACE = "K8S_NAMESPACE"
const K8S_PARTYID = "K8S_PARTYID"

const TOPIC_NAME = "fl-party"

type FlParty struct {
	PartyId     string `json:"partyId"`
	PartyName   string `json:"partyName"`
	UpdateAt    time.Time `json:"updateTime"`
	HeartBeatTs int64 `json:"heartbeatTs"`
	BoardUrl    string `json:"boardUrl"`
}

func NewFlParty() *FlParty {

	heartBeat := time.Now()
	heartBeatTs := heartBeat.Unix()

	partyName, ok := os.LookupEnv(K8S_NAMESPACE)
	if !ok {
		partyName = DEFAULT_NAMESPACE
	}

	partyId, ok := os.LookupEnv(K8S_PARTYID)
	if !ok {
		partyId = DEFAULT_PARTYID
	}

	return &FlParty{
		PartyId:     partyId,
		PartyName:   partyName,
		UpdateAt:    heartBeat,
		HeartBeatTs: heartBeatTs,
		BoardUrl:    fmt.Sprintf(BOARD_URL, partyId),
	}

}

func (party *FlParty) pub(ctx context.Context, client dapr.Client) {

	if party == nil {
		party = NewFlParty()
	}

	data, err := json.Marshal(party)
	if err != nil {
		log.Printf("error marshal %v\n", party)
		return
	}

	log.Printf("%v marshed to %s\n", party, data)

	err = client.PublishEvent(ctx, COMPONENT_PUBSUB_NAME, TOPIC_NAME, data)
	if err != nil {
		log.Println(err)
	}

}
