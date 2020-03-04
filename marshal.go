package config

// Marshaller marshal and unmarshall changes set into config values
type Marshaller interface {
	Marshal(*Snapshot) (Values, error)
	Unmarshal(Values) (*Snapshot, error)
}

// Encoder provide encode and decode data stream to obfuscate config content
type Encoder interface {
	Encode([]byte) []byte
	Decode([]byte) []byte
}
