package dns

type RecordValid interface {
	Valid() error
}
