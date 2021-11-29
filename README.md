Ticket Viewer
==
# Quick Start

## 1. Server
First, navigate into `server` directory

### Read Config Mechanism
There are two way to read config, read from a file, and read from environment variables. The environment variable priority is higher than reading from a file when the configs conflict.
It will read `.env` from the project root directory as the config file. The format of the config file is a standard envfile format.

Please fill in `ZENDESK_DOMAIN`, `ZENDESK_CRED_EMAIL`, and `ZENDESK_CRED_API_TOKEN` enviroment variables, which are required to let server work.
#### Available Environment Variable
- `PORT`: Default value is `:25976`
- `RUN_MODE`: Default value is `release`
- `READ_TIMEOUT`: Default value is `180` seconds
- `WRITE_TIMEOUT`: Default value is `60` seconds
- `ALLOWED_ORIGINS`: Allowed frontend origins **Required**
- `ZENDESK_DOMAIN`: Zendesk domain. **Required**
- `ZENDESK_CRED_EMAIL`: Zendesk credential: email **Required**
- `ZENDESK_CRED_API_TOKEN`: Zendesk credential: API token **Required**

### Build from Source
After setting up config file, it's now to build the project as an executable file.

#### Executable file
0. Manually fill in the .env file under `/server/env` directory
1. Execute `make install ENV=<env filename>` in the server directory
2. Executable file will show in the server directory
3. Specified config file will be copied to to the project root directory and been renamed to `.env` or pass environment variable when execute. For example, execute `make install ENV=local`, local.env will be copied to `/server/.env`
4. Execute the binary (./server)
5. Your service is started at port 25976


### Run Tests
To run tests, please execute `make test` it will run all unit tests written in this project. After running, a `test.out` coverage report will show in project root directory.

## 2. Client
Navigate into `client` directory

Run `npm start`
It will runs the app in the development mode.

Open http://localhost:3000 to view the ticket viewer frontend UI in the browser.

### Demo
#### Ticket Details
<img width="743" alt="截圖 2021-11-29 上午3 56 51" src="https://user-images.githubusercontent.com/17703676/143863995-3508076c-db56-474d-a3ad-7a2d8c23d46a.png">

#### Ticket List
<img width="817" alt="4" src="https://user-images.githubusercontent.com/17703676/143864355-674ccc29-68a8-4eae-96f6-e2e54423f4ba.png">

#### Pagination
<img width="765" alt="3" src="https://user-images.githubusercontent.com/17703676/143864395-0f11a83c-435c-45ea-bc87-fc2ec175c5b7.png">

#### Error Handle
<img width="1027" alt="1" src="https://user-images.githubusercontent.com/17703676/143864472-7ee81e48-3c05-41bb-8ff2-492b7447149b.png">

