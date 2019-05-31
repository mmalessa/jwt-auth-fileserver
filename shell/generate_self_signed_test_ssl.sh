#!/bin/bash

openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout ../ssl/test.key -out ../ssl/test.crt
