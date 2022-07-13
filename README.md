

SBOM-poc
---

A simple microservice to produce the Software Bill of Materials.


## Prerequisites

To make the microservice work have a
 - go version >= 1.17.8



## Run SBOMService [using Go]

SBOMService requires  **go1.17 or higher**  to run successfully. Run the following commands to build the latest version
```sh
git clone https://github.com/vishnusomank/sbom-poc.git
cd sbom-poc
go mod tidy
go build -o SBOMService main.go 
```
To run the program use
```sh
./SBOMService
```

## Run SBOMService [using Docker]

```sh
git clone https://github.com/vishnusomank/sbom-poc.git
cd sbom-poc
docker build -t sbomsvc . 
```
To run the program use
```sh
docker run -d -p 8080:8080 sbomsvc
```

## Usage
The server exposes 4 APIs 

 - /sbomservice/api/v1/all_scanned-images
 - /sbomservice/api/v1/scanned-image/:id
 - /sbomservice/api/v1/add-image
 - /sbomservice/api/v1/get-vuln/:id

| ENDPOINT | TYPE  | DATA | EXPLANATION |
|--|--|--|--|
| /sbomservice/api/v1/all_scanned-images | GET  | NIL | Returns keyvalue pair of Images already scanned |
|/sbomservice/api/v1/scanned-image/:id| GET  | ID | Displays SBOM value for the image stored with the specific ID|
|/sbomservice/api/v1/get-vuln/:id| GET  | ID | Displays vulnerabilities from the SBOM result of specific ID|
|/sbomservice/api/v1/add-image| POST | {  "image": "value" , "version":"value" }| Generated SBOM data for the input image. eg: image:ubuntu, version:20.04|


## Examples

 1. List all scanned images

	```sh
	curl "<IP:PORT>/sbomservice/api/v1/all_scanned-images"

	{"ID":1,"IMAGE NAME":"ubuntu","IMAGE VERSION":"latest"}
	{"ID":2,"IMAGE NAME":"debian","IMAGE VERSION":"latest"}
	Total Records loaded = 2
	```
2. List SBOM for specific ID

    ```
    curl -s "http://localhost:8080/sbomservice/api/v1/scanned-image/1"
    
    {"ID":1,"IMAGE NAME":"ubuntu","IMAGE VERSION":"latest"}
    {
    	        "SBOM": {
    		    	"artifacts": [
    	                {
    	                    "cpes": [
    	                        "cpe:2.3:a:adduser:adduser:3.118ubuntu5:*:*:*:*:*:*:*"
    	                    ],
    	                    "foundBy": "dpkgdb-cataloger",
    	                    "id": "78ce150ba8cd5542",
    	                    "language": "",
    	                    "licenses": [
    	                        "GPL-2"
    	                    ],
    	                    "locations": [
    	                        {
    	                            "layerID": "sha256:a790f937a6aea2600982de54a5fb995c681dd74f26968d6b74286e06839e4fb3",
    	                            "path": "/usr/share/doc/adduser/copyright"
    	                        },
    	                        {
    	                            "layerID": "sha256:a790f937a6aea2600982de54a5fb995c681dd74f26968d6b74286e06839e4fb3",
    	                            "path": "/var/lib/dpkg/info/adduser.conffiles"
    	                        },
    	                        {
    	                            "layerID": "sha256:a790f937a6aea2600982de54a5fb995c681dd74f26968d6b74286e06839e4fb3",
    	                            "path": "/var/lib/dpkg/info/adduser.md5sums"
    	                        },
    	                        {
    	                            "layerID": "sha256:a790f937a6aea2600982de54a5fb995c681dd74f26968d6b74286e06839e4fb3",
    	                            "path": "/var/lib/dpkg/status"
    	                        }
    	                    ],
   ......................................................................................................................	                    

3. Generate SBOM for a particular image
	
	```
    curl -X POST "http://localhost:8080/sbomservice/api/v1/add-image" -d '{"image": "alpine", "version": "latest"}'
    {"Submitted":"alpine:latest"}
    ``` 
4. View vulnerabilites for a particular image [Generated from SBOM]
	
    ```
    curl -s "http://localhost:8080/sbomservice/api/v1/get-vuln/1"
    
    {"ID":1,"IMAGE NAME":"redis","IMAGE VERSION":"latest"}
	{
    "Vulnerabilities": {
        "descriptor": {
            "configuration": {
                "add-cpes-if-none": false,
                "attestation": {
                    "public-key": "",
                    "skip-verification": false
                },
                "check-for-app-update": true,
                "configPath": "",
                "db": {
                    "auto-update": true,
                    "ca-cert": "",
                    "cache-dir": "/home/zero/.cache/grype/db",
                    "max-allowed-built-age": 432000000000000,
                    "update-url": "https://toolbox-data.anchore.io/grype/databases/listing.json",
                    "validate-age": true,
                    "validate-by-hash-on-start": false
                },
                "dev": {
                    "profile-cpu": false,
                    "profile-mem": false
                },
                "distro": "",
                "exclude": [],
                "externalSources": {
                    "enable": false,
                    "maven": {
                        "baseUrl": "https://search.maven.org/solrsearch/select",
                        "searchUpstreamBySha1": true
                    }
                },
	................................................................................................
    ``` 


