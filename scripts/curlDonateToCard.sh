#!/bin/bash

curl -XPOST \
    http://127.0.0.1:1317/cardservice/donate_to_card \
    -d '{
    "base_req":{
        "from":"'$(cscli keys show alice --address)'",
        "chain_id":"testCardchain",
        "gas": "auto",
        "gas_adjustment": "1.5"
      },
      "cardid":"1",
      "amount":"1credits",
      "donator":"'$(cscli keys show alice --address)'"
    }' > unsignedTx.json
