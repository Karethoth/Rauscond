package rauscond

import (
  "net";
  "container/list";
)



type UserInfo struct {
  Id string;
  IN chan string;
  Con *net.Conn;
  Quit chan bool;
  Partner *UserInfo;
  UserList *list.List;
  WantsPartner bool;
  WantsToBreakUp bool;
}



func (u *UserInfo) Init( id string, conn *net.Conn, partner *UserInfo, userList *list.List, wantsPartner bool ) {
  u.Id = id;
  u.Con = conn;
  u.Partner = partner;
  u.UserList = userList;
  u.WantsPartner = wantsPartner;
  u.WantsToBreakUp = false;
  u.Quit = make(chan bool);
  u.IN = make(chan string);
}



func (u *UserInfo) Read( buf []byte ) bool {
  _, err := (*u.Con).Read(buf);
  if err != nil {
    u.Close();
    return false;
  }
  return true;
}



func (u *UserInfo) Close() {
  u.Quit <- true;
  (*u.Con).Close();
  u.DeleteFromList();
}



func (u *UserInfo) Equal( b *UserInfo ) bool {
  if( u.Id == b.Id ) {
    if u.Con == b.Con {
      return true;
    }
  }
  return false;
}



func (u *UserInfo) DeleteFromList() {
  for e:= u.UserList.Front(); e != nil; e = e.Next() {
    client := e.Value.(*UserInfo);
    if u.Equal( client ) {
      if u.IsPaired() {
        u.WantsToBreakUp = true;
        u.BreakUp();
      }
      u.UserList.Remove( e );
    }
  }
}



func (u *UserInfo) IsPaired() bool {
  if u.Partner != nil {
   return true;
  }
  return false;
}



func (u *UserInfo) BreakUp() {
  if u.IsPaired() {
    (*u.Partner.Con).Write( []byte( "Stranger disconnected.\n" ) );
    u.Partner.Partner = nil;
  }
  u.Partner = nil;
}



func (u *UserInfo) Marry( b *UserInfo ) {
  if u.IsPaired() {
    return;
  }
  if b.IsPaired() {
    return;
  }

  u.Partner = b;
  b.Partner = u;

  u.WantsPartner = false;
  b.WantsPartner = false;
  u.WantsToBreakUp = false;
  b.WantsToBreakUp = false;
  Log( "Marrying was a success!" );
  (*u.Con).Write( []byte( "Talking with a stranger now.\n" ) );
  (*b.Con).Write( []byte( "Talking with a stranger now.\n" ) );
}

