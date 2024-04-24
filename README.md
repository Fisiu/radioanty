# RadioAnty

It seems that Antyradio has recently changed its streaming services. They are using some kind of load balancer now. Sending a request to the main stream results in adding a query param with a timestamp and then a redirection to a specific server instance.

Some streaming devices (like mine Revo SuperConnect Stereo) do not handle redirects properly, and this app solves above issue - it handles the redirect, adds a timestamp and returns the stream.
