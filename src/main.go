package rauscond

import (
  "net";
  "container/list";
  "strconv";
)



func ClearBuffer( buf *[]byte ) {
  for i:=0; i<2048; i++ {
    (*buf)[i]=0x00;
  }
}



func ClientReceiver( client *UserInfo ) {
  buf := make([]byte, 2048);
  Log( "ClientReceiver(): start for " + client.Id );
  for client.Read( buf ) {
    line := string( buf );
    line = line[:StrLen(line)-2];

    if line == "/quit" {
      client.Close();
      break;

    } else if line == "/next" {
      client.WantsToBreakUp = true;
      client.WantsPartner = true;
      ClearBuffer( &buf );
      continue;

    } else if line == "/break" {
      client.WantsToBreakUp = true;
      ClearBuffer( &buf );
      continue;
    }

    Log( "ClientReceiver(): received from " + client.Id + " (" + line + ")" );
    if client.Partner != nil {
      client.Partner.IN <- string(buf);
    }
    ClearBuffer( &buf );
  }
  Log( "ClientReceiver(): stop for: " + client.Id );
}



func ClientSender( client *UserInfo ) {
  Log( "ClientSender(): start for: " + client.Id );
  for {
    Log( "ClientSender(): wait for input to send" );
    select {
      case buf := <- client.IN:
        Log( "ClientSender(): send to \""+ client.Id + "\": "+ string(buf));
        (*client.Con).Write( []byte(buf) );

      case <-client.Quit:
        Log( "ClientSender(): client wants to quit" );
        (*client.Con).Close();
        break;
    }
  }
  Log( "ClientSender(): stop for: "+ client.Id );
}



func HandleNewClient( con *net.Conn, lst *list.List ) {
  id := "USER"+strconv.Uitoa64(*idCounter);
  *idCounter++;
  var newClient = new(UserInfo);
  newClient.Init( id, con, nil, lst, true );
  Log( "HandleClient(): for "+ id );
  PrintUsageToClient( newClient );
  go ClientSender( newClient );
  go ClientReceiver( newClient );
  lst.PushBack(newClient);
}




var debug *bool;
var idCounter *uint64;

func main() {
  debug = new(bool);
  *debug = true;
  idCounter = new(uint64);
  *idCounter = 1;
  Log( "main(): start" );

  userList := list.New();

  addr := net.TCPAddr{ net.ParseIP("127.0.0.1"), 9988 };
  netlisten, err := net.ListenTCP( "tcp", &addr );
  Test( err, "main Listen" );
  defer netlisten.Close();

  Log( "main(): starting UserPairer..." );
  var pairer UserPairer;
  pairer.Init( userList );
  go pairer.Start();


  for {
    Log( "main(): waiting for client..." );
    conn, err := netlisten.Accept();
    Test( err, "main: Accept for client" );
    go HandleNewClient( &conn, userList );
  }
}

