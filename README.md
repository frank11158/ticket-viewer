# Ticket Viewer Server
## Build from Source
There are two ways to build your program. First, build the project as an executable file. The other one is to build a docker image to execute the program.

### Docker Image
1. Execute `make build` in the project root directory
2. Execute `docker run -p 12345:25976 --env-file env/production.env ticket-viewer-server`
3. Your service is started at port 12345

### Executable file
1. Execute `make install ENV=<env filename>` in the project root (/server) directory
2. Executable file will show in the project root (/server) directory
3. Specified config file will be copied to to the project root directory and been renamed to `.env` or pass environment variable when execute
4. Execute the binary
5. Your service is started at port 25976

## Read Config Mechanism
There is two way to read config, read from a file, and read from environment variables. The environment variable priority is higher than reading from a file when the configs conflict.
It will read `.env` from the project root directory as the config file. The format of the config file is a standard envfile format.
### Available Environment Variable
- `PORT`: Default value is `:25976`
- `RUN_MODE`: Default value is `release`
- `READ_TIMEOUT`: Default value is `180` seconds
- `WRITE_TIMEOUT`: Default value is `60` seconds
- `ZENDESK_DOMAIN`: Zendesk domain. **Required**