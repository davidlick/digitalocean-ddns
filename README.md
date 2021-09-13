# digitalocean-ddns

A small Go DDNS application to update DigitalOcean Domain records to your current public IP.

## Usage

This application requires 3 environment variables to be set:

DIGITALOCEAN_TOKEN
  - A personal access token from your [DigitalOcean account](https://cloud.digitalocean.com/account/api/tokens). It requires read and write scopes.

DIGITALOCEAN_URL
  - This is the DigitalOcean API URL. At the time of writing this the API is located at https://api.digitalocean.com.

DIGITALOCEAN_DOMAIN
  - This is the domain name as configured in DigitalOcean. This application will not create a record if one does not already exist and it will only update an A record with the name "vpn". In the future I may generalize this, if you would like to change this yourself in the meantime you may submit a PR for it.

