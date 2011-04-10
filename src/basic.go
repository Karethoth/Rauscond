package rauscond

import (
  "fmt";
  "os";
  "strings";
)


func Log(v string) {
  if *debug == true {
    ret := fmt.Sprint( v );
    fmt.Printf( "SERVER: %s\n", ret );
  }
}



func Test( err os.Error, mesg string ) {
  if err != nil {
    fmt.Printf( "SERVER: ERROR: " + mesg );
    os.Exit(-1);
  } else
    Log( "Ok: "+ mesg );
}



func PrintUsageToClient( client *UserInfo ) {
  msg := string(
   "Welcome!\n"+
   "Basic commands:\n"+
   "  /next  - If talking already, disconnect and\n"+
   "           connect with a new stranger.\n"+
   "  /break - Break the conversation with stranger.\n"+
   "  /quit  - Quit.\n\n" );
  (*client.Con).Write( []byte( msg ) );
}



func StrLen( str string ) int {
  return strings.Index( str, string(0) );
}

