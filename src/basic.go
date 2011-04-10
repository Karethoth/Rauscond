package rauscond

import (
  "fmt";
  "os";
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

