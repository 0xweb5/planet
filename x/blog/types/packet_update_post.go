package types

// ValidateBasic is used for validating the packet
func (p UpdatePostPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p UpdatePostPacketData) GetBytes() ([]byte, error) {
	var modulePacket BlogPacketData

	modulePacket.Packet = &BlogPacketData_UpdatePostPacket{&p}

	return modulePacket.Marshal()
}
