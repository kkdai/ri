package ri

type ServerNI struct {
	ClientNI

	Id    string
	EIPv4 string
	EPort int
}

func NewServerNI(c *ClientNI, eip string, eport int) *ServerNI {
	s := new(ServerNI)
	s.IIPv4 = c.IIPv4
	s.IIPv6 = c.IIPv6
	s.INetmask = c.INetmask
	s.IPort = c.IPort

	s.EIPv4 = eip
	s.EPort = eport
	return s
}

//Return the NAT type according its Internal IP and Public IP
func (s *ServerNI) UseNAT() bool {
	return s.IIPv4 == s.EIPv4 && s.IPort == s.EPort
}

//Check each NAT type and make sure if it is under the same routing.
//Return "true" if it is not NAT_Symmetric
func (s *ServerNI) ValidToP2P(si *ServerNI) bool {
	return si != nil
}
