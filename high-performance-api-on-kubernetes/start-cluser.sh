#!/bin/bash

gcloud container clusters create $1 \
	--machine-type=n1-standard-16 \
	--num-nodes=3
