package message

type Header int
const (
    BAD  Header = iota
    PING
    GET
    PUT
)
func (header Header) String() string {
    headers := [...]string{
               "Bad",
               "Ping",
               "Get",
               "Put"}

    if(!header.Valid()) {
        return "Bad"
    }

    return headers[header]
}

func (header Header) Valid() bool {
    return header >= PING && header <= PUT
}
