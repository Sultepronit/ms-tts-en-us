#!/bin/bash

source ../.env

curl -o voices.json --location --request GET "https://$SPEECH_REGION.tts.speech.microsoft.com/cognitiveservices/voices/list" \
--header "Ocp-Apim-Subscription-Key: $SPEECH_KEY"