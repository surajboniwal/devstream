package idgen

import (
	"crypto/sha1"
	"fmt"
	"net"

	"github.com/bwmarrin/snowflake"
)

type SnowflakeIdGen struct {
}

func NewSnowflakeIdGen() *SnowflakeIdGen {
	return &SnowflakeIdGen{}
}

var node *snowflake.Node

func init() {
	n, err := snowflake.NewNode(int64(getMachineID()))

	if err != nil {
		logger.E(err)
	}

	node = n
}

func (sf SnowflakeIdGen) New() int64 {
	sfId := node.Generate()
	return sfId.Int64()
}

func getMachineID() int {
	macAddr, err := getFirstMACAddress()
	if err != nil {
		return 0
	}

	id := fmt.Sprintf("%s-%s", macAddr, "additional-unique-info")
	hash := sha1.New()
	_, _ = hash.Write([]byte(id))
	hashResult := hash.Sum(nil)

	machineID := int(hashResult[0]) % 1024

	return machineID
}

func getFirstMACAddress() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, intf := range interfaces {
		if intf.Flags&net.FlagUp != 0 && intf.Flags&net.FlagLoopback == 0 {
			return intf.HardwareAddr.String(), nil
		}
	}
	return "", fmt.Errorf("no active MAC address found")
}
