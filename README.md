# Status : Write status messages to your terminal
Status provides tools to output short status messages (in a terminal) in a Linux/SysV init fashion.

Basic usage :

	sw := status.NewWriter(termWidth)
	sw.Pendingf("Doing some stuff %v", x)
	sw.MkSuccess().Done()
    
Example :

![screenshot](https://github.com/fjecker/status/raw/master/img/screenshot.png)

Currently status allows you to use 4 different statuses :
	
* pending (....)
* success (PASS)
* warning (WARN)
* failure (FAIL)

Please have a look at the **godoc** and at the unit tests for further usage details

A status message can be updated (text or status) until the "Done()" method is called.

Please note that status relies on [ansicolor](github.com/shiena/ansicolor) to add a bit of color in its output (with a windows support)
    
*Status was initially developped for the [gobot](http://github.com/fjecker/gobot) project and was written while learning/discovering go* 