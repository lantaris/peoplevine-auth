# Cloud Datalab PeopleVine auth server

Cloud Datalab PeopleVine authentication proxy server
Service features:
   The service accepts a standard for the PeopleVine API structure for authorization as input, accesses the PeopleVine authorization service with it, receives the result, and generates a formatted simplified response to the user.

## Topics

- [Building binaries](#Building-binaries)
  - [Building RPM package](#Building-RPM-package)
  - [Building Docker/Kubernetes image](#Building-Docker/Kubernetes-image)
- [Install and run service](#Install-and-run-service)
- [Endpoints](#Endpoints)
   - [Request for detailed response](#Request-for-detailed-response)
   - [Request for simple response](#Request-for-simple-response)
- [Response status codes](#Response-status-codes)

## Building binaries

### Building RPM package

Build requiremets: sudo, Docker

The assembly of the package is completely isolated and independent of the system environment.

Run build process:

```
make rpm
```

The build result is in: ./build/

### Building Docker/Kubernetes image

Build requiremets: sudo, Docker, j2cli( pip3 install j2cli )

The container image is also built in a sandboxed Docker environment.

Edit docker/SETTINGS to point to your docker registry to store the image.

Run build image:

```
make image
```

Upload image to registry

```
make push_image
```

Image run parameters:

```
[root@fc33 peoplevine-auth]# docker run -ti --rm hub.docker.com/lantaris/peoplevine-auth --help
Usage: peoplevine-auth [--debug DEBUG] [--listen LISTEN]

Options:
  --debug DEBUG, -d DEBUG
                         Debug level
  --listen LISTEN, -l LISTEN
                         REST listen
  --help, -h             display this help and exit
[root@fc33 peoplevine-auth]#
```

## Install and run service

Depends RPM based Linux system (Prefered: CentOS 8)

```
 # For new instalation sms-center
 yum install peoplevine-auth-0.0.1-1.el9.x86_64.rpm
 
 # For upgrade an existing installation
 yum update peoplevine-auth-<new version>.el8.x86_64.rpm 
 
 # Run peoplevine-auth service
 systemctl -now enable peoplevine-auth
```

## Endpoints

Endpoints for service 

### Request for detailed response

|  Request        | Value                                 |   Description                    |
| --------------- | ------------------------------------- | -------------------------------- |
| Request type:   | POST                                  |                                  |
| URL:            | http://<IP:port>/auth                 |                                  |
| Data:           | {                                     |                                  |
|                 |  "auth": {                            |                                  |
|                 |     "api_username": string,           | PeopleVine API user name         |  
|                 |     "api_password": string,           | PeopleVine API password          | 
|                 |     "api_key": string,                | PeopleVine API key               |
|                 |     "company_no": int,                | User company ID                  |
|                 |     "username": string,               | Company user name                |
|                 |     "password": string                | Company password                 |
|                 |     },                                |                                  |
|                 |  "credentials": {                     |                                  |
|                 |     "company_no": int,                | User company ID                  |
|                 |     "email": string,                  | User email                       |
|                 |     "password": string                | User password                    |
|                 |     }                                 |                                  |
|                 | }                                     |                                  |


|  Response       | Value                                  |   Description                    |
| --------------- | -------------------------------------- | -------------------------------- |
| Data:           | {                                      |                                  |
|                 |   "status": {                          | Return status code and message   |
|                 |     "errcode": int,                    | ( if 0 authentication succes )   |
|                 |     "errmsg": string                   |                                  |
|                 |   },                                   |                                  |
|                 |   "data": {                            |                                  |
|                 |     "responseCode": string,            |                                  |
|                 |     "message": string,                 |                                  |
|                 |     "isError": bool,                   |                                  |
|                 |     "methodFailed": string,            |                                  |
|                 |     "extendedMessage": string,         | if isError  == flase &&          |
|                 |     "reason": string,                  |    isMember == true = Auth OK    |
|                 |     "returnObject": {                  |                                  |
|                 |       "app_installed_on": string,      |                                  |
|                 |       "app_installed": bool,           |                                  |
|                 |       "network_company_no": int,       |                                  |
|                 |       "network_company_name": string,  |                                  |
|                 |       "adminOverride": bool,           |                                  |
|                 |       "map_image": string,             |                                  |
|                 |       "lifecycle_stage": string,       |                                  |
|                 |       "ip_address": string,            |                                  |
|                 |       "memberDirectoryEnabled": bool,  |                                  |
|                 |       "enableHouseAccount": bool,      |                                  |
|                 |       "keep_private": bool,            |                                  |
|                 |       "customer_token": string,        |                                  |
|                 |       "member_discount": int,          |                                  |
|                 |       "name": string,                  |                                  |
|                 |       "cc_email": string,              |                                  |
|                 |       "isMember": bool,                |                                  |
|                 |       "wasMember": bool,               |                                  |
|                 |       "hasApplied": bool,              |                                  |
|                 |       "member_since": string,          |                                  |
|                 |       "when_applied": string,          |                                  |
|                 |       "updateByReference": bool,       |                                  |
|                 |       "addresses": array,              |                                  |
|                 |       "email2": string,                |                                  |
|                 |       "notes": array,                  |                                  |
|                 |       "gender": string,                |                                  |
|                 |       "company_title": string,         |                                  |
|                 |       "anniversary_date": string,      |                                  |
|                 |       "hasPassword": bool,             |                                  |
|                 |       "last_login": string,            |                                  |
|                 |       "lastPageView": string,          |                                  |
|                 |       "lastEmailOpened": string,       |                                  |
|                 |       "lastNewsletter": string,        |                                  |
|                 |       "lastTouchpoint": string,        |                                  |
|                 |       "lastCommunication": string,     |                                  |
|                 |       "totalTouchpoints": int,         |                                  |
|                 |       "user": {                        |                                  |
|                 |         "business_location_no": int,   |                                  |
|                 |         "employee_id": string,         |                                  |
|                 |         "user_no": int,                |                                  |
|                 |         "first_name": string,          |                                  |
|                 |         "last_name": string,           |                                  |
|                 |         "mobile_number": string,       |                                  |
|                 |         "email": string,               |                                  |
|                 |         "username": string,            |                                  |
|                 |         "status": string               |                                  |
|                 |       },                               |                                  |
|                 |       "user_modified": {               |                                  |
|                 |         "business_location_no": int,   |                                  |
|                 |         "employee_id": string,         |                                  |
|                 |         "user_no": int,                |                                  |
|                 |         "first_name": string,          |                                  |
|                 |         "last_name": string,           |                                  |
|                 |         "mobile_number": string,       |                                  |
|                 |         "email": string,               |                                  |
|                 |         "username": string,            |                                  |
|                 |         "status": string               |                                  |
|                 |       },                               |                                  |
|                 |       "user_last_viewed": {            |                                  |
|                 |         "business_location_no": int,   |                                  |
|                 |         "employee_id": string,         |                                  |
|                 |         "user_no": int,                |                                  |
|                 |         "first_name": string,          |                                  |
|                 |         "last_name": string,           |                                  |
|                 |         "mobile_number": string,       |                                  |
|                 |         "email": string,               |                                  |
|                 |         "username": string,            |                                  |
|                 |         "status": string               |                                  |
|                 |       },                               |                                  |
|                 |       "last_viewed": string,           |                                  |
|                 |       "last_viewed_by": int,           |                                  |
|                 |       "modified_by": int,              |                                  |
|                 |       "affiliate_name": string,        |                                  |
|                 |       "affiliate_no": int,             |                                  |
|                 |       "session_id": string,            |                                  |
|                 |       "password_expires": string,      |                                  |
|                 |       "flag": string,                  |                                  |
|                 |       "vip_flag": bool,                |                                  |
|                 |       "last_contacted": string,        |                                  |
|                 |       "klout_score": int,              |                                  |
|                 |       "linkedin_handle": string,       |                                  |
|                 |       "linkedin_id": string,           |                                  |
|                 |       "lifecycle_stage_no": int,       |                                  |
|                 |       "customer_source": string,       |                                  |
|                 |       "timezone": string,              |                                  |
|                 |       "username": string,              |                                  |
|                 |       "settings": string,              |                                  |
|                 |       "generatePassword": bool,        |                                  |
|                 |       "profile_photo": string,         |                                  |
|                 |       "foursquare_handle": string,     |                                  |
|                 |       "foursquare_id": string,         |                                  |
|                 |       "customer_status": string,       |                                  |
|                 |       "attributes": array,             |                                  |
|                 |       "facebook_handle": string,       |                                  |
|                 |       "twitter_handle": string,        |                                  |
|                 |       "tiktok_handle": string,         |                                  |
|                 |       "instagram_handle": string,      |                                  |
|                 |       "latitude": string,              |                                  |
|                 |       "longitude": string,             |                                  |
|                 |       "customer_no": int,              |                                  |
|                 |       "first_name": string,            |                                  |
|                 |       "last_name": string,             |                                  |
|                 |       "address": string,               |                                  |
|                 |       "address2": string,              |                                  |
|                 |       "city": string,                  |                                  |
|                 |       "state": string,                 |                                  |
|                 |       "zip_code": string,              |                                  |
|                 |       "country": string,               |                                  |
|                 |       "phone_number": string,          |                                  |
|                 |       "mobile_number": string,         |                                  |
|                 |       "email": string,                 |                                  |
|                 |       "password": string,              |                                  |
|                 |       "created_on": string,            |                                  |
|                 |       "device_type": string,           |                                  |
|                 |       "passbook_enabled": string,      |                                  |
|                 |       "wallet_enabled": string,        |                                  |
|                 |       "opt_in_email": string,          |                                  |
|                 |       "opt_in_sms": string,            |                                  |
|                 |       "opt_in_updates": string,        |                                  |
|                 |       "company_no": int,               |                                  |
|                 |       "user_no": iny,                  |                                  |
|                 |       "company_name": string,          |                                  |
|                 |       "birthdate": string,             |                                  |
|                 |       "modified_on": string,           |                                  |
|                 |       "address3": string,              |                                  |
|                 |       "customer_reference": string,    |                                  |
|                 |       "facebook_id": string,           |                                  |
|                 |       "twitter_id": string,            |                                  |
|                 |       "tiktok_id": string,             |                                  |
|                 |       "instagram_id": string,          |                                  |
|                 |       "middle_name": string,           |                                  |
|                 |       "website": string,               |                                  |
|                 |       "customer_type": string,         |                                  |
|                 |       "schema": string,                |                                  |
|                 |       "totalFavorite": int,            |                                  |
|                 |       "skipTriggers": bool,            |                                  |
|                 |       "includeObjects": bool,          |                                  |
|                 |       "includeStats": bool,            |                                  |
|                 |       "includeTotalRecords": bool,     |                                  |
|                 |       "totalRecords": int,             |                                  |
|                 |       "rows_per_page": int,            |                                  |
|                 |       "page_number": int,              |                                  |
|                 |       "updateFields": string,          |                                  |
|                 |       "includeCompany": bool,          |                                  |
|                 |       "lightweight": bool,             |                                  |
|                 |       "itemDetailsFields": string,     |                                  |
|                 |       "includeItemDetails": bool,      |                                  |
|                 |       "search_value": string,          |                                  |
|                 |       "sort_by": string,               |                                  |
|                 |       "return_total": int,             |                                  |
|                 |       "from_date": string,             |                                  |
|                 |       "to_date": string,               |                                  |
|                 |       "includeDetails": bool,          |                                  |
|                 |       "timezone_offset": int,          |                                  |
|                 |       "timezone_id": string            |                                  |
|                 |     },                                 |                                  |
|                 |     "responseTime": int,               |                                  |
|                 |     "totalRecords": int,               |                                  |
|                 |     "minPrice": int,                   |                                  |
|                 |     "maxPrice": int,                   |                                  |
|                 |     "permission_level": string,        |                                  |
|                 |     "timezone_offset": int,            |                                  |
|                 |     "timezone_id": int                 |                                  |
|                 |   }                                    |                                  |
|                 | }                                      |                                  |
| Response code   | 200 OK                                 |                                  |

### Request for simple response

|  Request        | Value                                 |   Description                    |
| --------------- | ------------------------------------- | -------------------------------- |
| Request type:   | POST                                  |                                  |
| URL:            | http://<IP:port>/simple               |                                  |
| Data:           | {                                     |                                  |
|                 |  "auth": {                            |                                  |
|                 |     "api_username": string,           | PeopleVine API user name         |  
|                 |     "api_password": string,           | PeopleVine API password          | 
|                 |     "api_key": string,                | PeopleVine API key               |
|                 |     "company_no": int,                | User company ID                  |
|                 |     "username": string,               | Company user name                |
|                 |     "password": string                | Company password                 |
|                 |     },                                |                                  |
|                 |  "credentials": {                     |                                  |
|                 |     "company_no": int,                | User company ID                  |
|                 |     "email": string,                  | User email                       |
|                 |     "password": string                | User password                    |
|                 |     }                                 |                                  |
|                 | }                                     |                                  |


|  Response       | Value                                  |   Description                    |
| --------------- | -------------------------------------- | -------------------------------- |
| Data:           | {                                      |                                  |
|                 |   "status": {                          | Return status code and message   |
|                 |     "errcode": int,                    | ( if 0 authentication succes )   |
|                 |     "errmsg": string                   |                                  |
|                 |   },                                   |                                  |
|                 |   "data": null                         |                                  |
|                 | }                                      |                                  |
| Response code   | 200 OK                                 |                                  |

## Response status codes

Response status codes (response_status)

|  Code |  Status           |
| ----- | ----------------- |
| 0     | RET_OK
| 8     | RET_ERROR
| 16    | RET_PARAM_PARSE_ERROR
| 32    | RET_CALL_ERROR
