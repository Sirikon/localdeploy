#!/usr/bin/env bash
zip -r artifact ./molly
curl -F artifact=@./artifact.zip -F project=$MOLLY_PROJECT -F token=$MOLLY_TOKEN $MOLLY_URL"/deploy"
