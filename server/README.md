# Ticket Viewer Server
## Build from Source
First, please build the project as an executable file.

### Executable file
0. Manually fill in the .env file under `/server/env` directory
1. Execute `make install ENV=<env filename>` in the project root (/server) directory
2. Executable file will show in the project root (/server) directory
3. Specified config file will be copied to to the project root directory and been renamed to `.env` or pass environment variable when execute. For example, execute `make install ENV=local`, local.env will be copied to `/server/.env`
4. Execute the binary
5. Your service is started at port 25976

## Read Config Mechanism
There are two way to read config, read from a file, and read from environment variables. The environment variable priority is higher than reading from a file when the configs conflict.
It will read `.env` from the project root directory as the config file. The format of the config file is a standard envfile format.
### Available Environment Variable
- `PORT`: Default value is `:25976`
- `RUN_MODE`: Default value is `release`
- `READ_TIMEOUT`: Default value is `180` seconds
- `WRITE_TIMEOUT`: Default value is `60` seconds
- `ALLOWED_ORIGINS`: Allowed frontend origins **Required**
- `ZENDESK_DOMAIN`: Zendesk domain. **Required**
- `ZENDESK_CRED_EMAIL`: Zendesk credential: email **Required**
- `ZENDESK_CRED_API_TOKEN`: Zendesk credential: API token **Required**