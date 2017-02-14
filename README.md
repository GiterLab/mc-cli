# mc-cli
client tool for memcached

### Build ###

		go get -u -v github.com/GiterLab/mc-cli

		or

		go get -u -v github.com/tools/godep
		git clone https://github.com/GiterLab/mc-cli.git
		cd mc-cli
		./build

### Using ###

	Toby@Giter-Toby-PC MINGW64 /d/prj/go/src/github.com/GiterLab/mc-cli (master)
	$ ./mc-cli.exe  --> ./mc-cli.exe -host 127.0.0.1:11211 to specify the address of the memcached server
	==============================
	memcache client v0.0.1
	==============================
	MC>> set hello tobyzxj
	OK
	
	MC>> get hello
	tobyzxj
	
	MC>> set hello tobyzxj 10
	OK
	
	MC>> get hello
	tobyzxj
	
	MC>> get hello
	[E] memcache: cache miss
	MC>> list
	------------------------
	  set: set key value, set key value expiration_time
	  get: get key
	  getmore: getmore key
	  list(l): list commands
	  quit(q): quit this app
	  exit(e): quit this app
	
	MC>> exit
	  Quit
	
	Toby@Giter-Toby-PC MINGW64 /d/prj/go/src/github.com/GiterLab/mc-cli (master)
	$

