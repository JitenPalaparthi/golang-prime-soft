
take pingpong string

Two goroutunes should send the string and read the data.


go ping() // send the data ping -- to a Global String..
go pong() // it has to receive the data ping and send pong to the other go routine..

both ping and pong should continuosly  run with a sleep of 100 milli second
