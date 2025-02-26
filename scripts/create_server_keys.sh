#!/usr/bin/env bash

# Clear old crt and key files
rm ./certs/*

# Generate a self-signed certificate with SAN (Subject Alternative Name)
openssl req -x509 -newkey rsa:4096 -nodes -keyout ./certs/server.key -out ./certs/server.crt -days 365 \
    -addext "subjectAltName=DNS:vpn.gophernest.net,DNS:www.vpn.gophernest.net,DNS:localhost,IP:127.0.0.1" \
    -subj "/O=Gophernest/OU=Notifications/L=Surprise/ST=Arizona/C=US"

# Generate a self-signed certificate with SAN (Subject Alternative Name)
openssl req -x509 -newkey rsa:4096 -nodes -keyout ./certs/client.key -out ./certs/client.crt -days 365 \
    -subj "/CN=testingclient/O=Gophernest/OU=Notifications"
