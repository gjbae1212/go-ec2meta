# go-ec2meta
<p align="left"> 
<a href="https://hits.seeyoufarm.com"/><img src="https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https://github.com/gjbae1212/go-ec2meta"/></a>
<a href="https://goreportcard.com/report/github.com/gjbae1212/go-ec2meta"><img src="https://goreportcard.com/badge/github.com/gjbae1212/go-ec2meta"/></a>
<a href="https://godoc.org/github.com/gjbae1212/go-ec2meta"><img src="https://godoc.org/github.com/gjbae1212/go-ec2meta?status.svg"/></a>
<a href="/LICENSE"><img src="https://img.shields.io/badge/license-MIT-GREEN.svg" alt="license"/></a>
</p>
This Library is to get metadata for EC2 Instances using Golang.    
It can get information like public-ip, hostname, mac .. and so on, if this machine is EC2 instance in AWS.  
Maybe you are useful when your application in Docker Container is getting IP information(host machine IP .. and so on).

## Usage
```golang
import  "github.com/gjbae1212/go-ec2meta"

hosts, err := ec2meta.PublicHostname()
ips, err := ec2meta.PublicIPV4()

// and so on.
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
