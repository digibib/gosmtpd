# A very simple SMTP server with a REST API

[![Build Status](https://drone.io/github.com/sorenmat/gosmtpd/status.png)](https://drone.io/github.com/sorenmat/gosmtpd/latest)
[![Coverage Status](https://coveralls.io/repos/sorenmat/gosmtpd/badge.svg)](https://coveralls.io/r/sorenmat/gosmtpd)

A server that accepts smtp request and saves the emails in memory for later retrieval.

## Usage
```shell

usage: gosmtpd [<flags>]

Flags:
  --help              Show help (also see --help-long and --help-man).
  --webport=WEBPORT   Port the web server should run on (Default: 8000)
  --hostname="localhost"  
                      Hostname for the smtp server to listen to
  --port="2525"       Port for the smtp server to listen to
  --forwardhost=FORWARD_HOST
                      The hostname after the @ that we should forward i.e. gmail.com
  --forwardsmtp=FORWARD_SMTP
                      SMTP server to forward the mail to
  --forwardport=FORWARD_PORT
                      The port on which email should be forwarded (Default: 25)
  --forwarduser=FORWARD_USER
                      The username for the forward host
  --forwardpassword=FORWARD_PASSWORD
                      Password for the user
  --mailexpiration=300  
                      Time in seconds for a mail to expire, and be removed from database
  --version           Show application version.

``` 

## GET /status
Returns a 200 if the service is up

## GET /mail 
List all mails

## GET /inbox/:email 
List all email for a given email address

## GET /email/:id 
Get an email by id

## DELETE /inbox/:email 
Delete all mails for a given email

## DELETE /email/:id 
Delete a email via the id

# Trying it out

You can install it by doing 

``docker pull sorenmat/gosmtpd``


``docker start sorenmat/gosmtpd``
