package jetoo

const (
    ValueUnknownString = 'UNKNOWN'
    ValueUnknownInt = -1
)

type DataSet struct {
    Scheme string
    Host string
    Port int32
    User string
    Pass string
    Path string
    Query string
    Fragment string
}
