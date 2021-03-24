# tjsj.dev
### Yet another CS students personal web server

currently under development\
its very cool

## TODO List

- [x] Golang HTTPS server
- [ ] Routing & Navigation
	- [x] Routing & handlers
	- [x] HTML templates
	- [x] Navigation bar
	- [x] Custom error page
	- [x] Static public content / file server
	- [ ] Posts selection
- [ ] CSS Styling
	- [x] Dark theme
	- [ ] Font family
	- [ ] Mobile-friendly
- [ ] Home page
	- [ ] Title / Landing zone
	- [ ] About
	- [ ] Recent posts/projects
	- [ ] Links (Socials / CV / etc)
- [ ] Posts
	- [ ] List of posts
	- [ ] Search by content
	- [ ] Sort by date
	- [ ] Filter by tags
	- [ ] Embed web projects (p5.js/etc)
	- [ ] Download links
	- [ ] Seperate presentation for project?
- [ ] Static git pages
	- [ ] Stagit
	- [ ] Git post-hooks
	- [ ] Similar/Shared CSS styling
- [ ] Pageview analytics
	- [ ] Collect anonymous request stats
		- [ ] Visits aggregation
		- [ ] Referrals
		- [ ] Internal navigation
		- [ ] Outgoing links
		- [ ] Date / Time
		- [ ] Duration of visit? I don't want user state, no identifying information
	- [ ] Save data
	- [ ] Terminal output? Separate viewing program?
	- [ ] Present some data on website?
- [ ] Security
	- [ ] Sanitized URL parsing
	- [ ] Any input injection
	- [x] Response timeout
	- [ ] Anything else I'm currently forgetting
- [ ] Deployment
	- [x] Secure personal server
	- [ ] TUI status
	- [ ] Email notifcations
- [x] Title splashes
- [ ] RSS feed / Mailing list?
- [ ] Mail server / forwarding?
- [ ] VPN?

## Setting up & Running

Clone the repository
```
$ git clone https://github.com/tedski999/tjsj.dev.git
$ cd tjsj.dev
```

Build & run
```
$ make run
```

...or compile manually using the Go compiler
```
$ go build -o ./bin/tjsj ./src
```

The website should now be hosted on your machine accessable via port 443:
https://localhost

