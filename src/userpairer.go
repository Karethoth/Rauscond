package rauscond

import (
  "container/list";
  "time";
);

type UserPairer struct {
  userList *list.List;
  unpairedUsers *list.List;
}


func (up *UserPairer) Init( userlist *list.List ) {
  up.userList = userlist;
}


func (up *UserPairer) Start() {
  for {
    up.unpairedUsers = new(list.List);

    // Run through pairs and breakup those in need to
    for e := up.userList.Front(); e != nil; e = e.Next() {
      user := e.Value.(*UserInfo);
      if user.WantsToBreakUp {
        user.BreakUp();
      }
    }
    // Run through all users and generate list from the one's who need a partner
    for e := up.userList.Front(); e != nil; e = e.Next() {
      user := e.Value.(*UserInfo);
      if user.WantsPartner {
        if !user.IsPaired() {
          up.unpairedUsers.PushBack( user )
        }
      }
    }
    // Randomize the list of users in need of a pair
    // Generate pairs from them
    for e := up.unpairedUsers.Front(); e != nil; e = e.Next() {
      user := e.Value.(*UserInfo);
      if e.Next() != nil {
        next := e.Next().Value.(*UserInfo);
        if !user.Equal( next ) {
          Log( "Marry "+user.Id+" & "+next.Id );
          user.Marry( next );
          e = e.Next();
        }
      }
    }

    time.Sleep( 10 );
  }
}

