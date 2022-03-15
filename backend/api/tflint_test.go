package api

import (
    "github.com/jgroeneveld/trial/assert"
    "testing"
)

func TestTFLint(t *testing.T) {
	for _, test := range []apiTest{
		{
			in: []byte(`resource "aws_subnet" "subnet-123456789" {
  assign_ipv6_address_on_creation                = false
  availability_zone                              = "us-east-1a"
  availability_zone_id                           = "use1-az1"
  cidr_block                                     = "192.168.24.0/22"
  enable_dns64                                   = false
  enable_resource_name_dns_a_record_on_launch    = false
  enable_resource_name_dns_aaaa_record_on_launch = false
  ipv6_native                                    = false
  map_customer_owned_ip_on_launch                = false
  map_public_ip_on_launch                        = false
  private_dns_hostname_type_on_launch            = "ip-name"
  tags = {
    Env                    = "dev"
    Owner                  = "TeamX"
  }
  vpc_id = "vpc-112233"
}
`),
		},
	} {
		out, err := TFLint(test.in)
		if test.err == nil {
			assert.Equal(t, string(test.out), string(out), "out must match")
		} else {
			assert.Equal(t, test.err, err, "err must match")
		}
	}
}
