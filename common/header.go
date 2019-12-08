package message

type Header int
const ( 
    PING Header = iota
    GET
    PUT
)
func (header Header) String() string {
    headers := [...]string{
               "PING",
               "GET",
               "PUT"}
    
    if(!header.Valid()) {
        return "BAD"
    }

    return headers[header]
}

func (header Header) Valid() bool {
    return header >= PING && header <= PUT
}
