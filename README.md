# go-ec2meta

This Library is to get metadata for EC2 Instances using Golang.    
It can get information like public-ip, hostname, mac .. and so on if this machine is EC2 instance in AWS.  
Maybe you are useful when your application in Docker Container is getting IP information(host machine IP .. and so on).

## Usage
```golang
# Example
import  "github.com/gjbae1212/go-ec2meta"

hostnames, err := ec2meta.Hostname()
```

## Method
| method name | description |
| ------------|-------------|
| *Hostname*  | hostname |
| *LocalHostname*  | local hostname |
| *LocalIPV4*  | local ip information(V4) |
| *PublicHostname*  | public hostname |
| *PublicIPV4*  | public ip information(V4) |
| *Mac*  | instance mac address |


## Features (TODO)
* [X] hostname 
* [X] local-hostname
* [X] local-ipv4
* [X] public-hostname
* [X] public-ipv4
* [X] mac
* [ ] ami-id
* [ ] ami-launch-index
* [ ] ami-manifest-path
* [ ] block-device-mapping/ (multiple option)
* [ ] block-device-mapping/ (multiple option)
* [ ] events/ (multiple option)
* [ ] iam/ (multiple option)
* [ ] instance-action
* [ ] instance-id
* [ ] instance-type
* [ ] profile
* [ ] reservation-id
* [ ] metrics/ (multiple option)
* [ ] network/ (multiple option)
* [ ] placement/ (multiple option)
* [ ] public-keys/ (multiple option)
* [ ] security-groups (multiple option)
* [ ] services/ (multiple option)   
 
## LICENSE
This project is following The MIT.
