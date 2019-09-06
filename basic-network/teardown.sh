#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error, print all commands.
set -e

# Shut down the Docker containers for the system tests.
docker-compose -f docker-compose.yml kill && docker-compose -f docker-compose.yml down

# remove the local state
rm -rf ../fabcctv/hfc-key-store/*

# stop docker running container before deleting images
docker stop $(docker ps -a -q)
docker rm -f $(docker ps -a -q)

# remove chaincode docker images
docker rmi -f $(docker images -q)

# Your system is now clean
