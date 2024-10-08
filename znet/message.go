package znet

// Message 消息
type Message struct {
	MagicCode  uint16 // 魔数（小安必须）
	ID         uint8  // 消息的ID（命令字，小安必须）
	Sn         uint8  // 序列号（小安必须）
	DataLen    uint16 // 消息的长度（小安必须）
	HeaderData []byte // 消息头
	Data       []byte // 消息的内容
	ExtendData []byte // 扩展消息
}

// NewMsgPackage 创建一个Message消息包
func NewMsgPackage(ID, Sn uint8, data []byte) *Message {
	return &Message{
		DataLen: uint16(len(data)),
		ID:      ID,
		Sn:      Sn,
		Data:    data,
	}
}

// GetMagicCode 获取魔数
func (msg *Message) GetMagicCode() uint16 {
	return msg.MagicCode
}

// GetDataLen 获取消息数据段长度
func (msg *Message) GetDataLen() uint16 {
	return msg.DataLen
}

// GetSn 获取头部序列号
func (msg *Message) GetSn() uint8 {
	return msg.Sn
}

// GetMsgID 获取消息ID
func (msg *Message) GetMsgID() uint8 {
	return msg.ID
}

// GetData 获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

// GetHeaderData 获取消息头内容
func (msg *Message) GetHeaderData() []byte {
	return msg.HeaderData
}

// GetExtendData 获取扩展消息
func (msg *Message) GetExtendData() []byte {
	return msg.ExtendData
}

// SetDataLen 设置消息数据段长度
func (msg *Message) SetDataLen(len uint16) {
	msg.DataLen = len
}

// SetMsgID 设置消息ID
func (msg *Message) SetMsgID(msgID uint8) {
	msg.ID = msgID
}

// SetHeaderData 设置消息头内容
func (msg *Message) SetHeaderData(data []byte) {
	msg.HeaderData = data
}

// SetData 设置消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}

// SetMagicCode 设置魔数
func (msg *Message) SetMagicCode(magicCode uint16) {
	msg.MagicCode = magicCode
}

// SetExtendData 设置扩展数据
func (msg *Message) SetExtendData(extendData []byte) error {
	msg.ExtendData = extendData
	return nil
}
