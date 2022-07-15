# Frens

[![Open Source Love svg1](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)

[![Go](https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff)](https://golang.org/)


Frens is an open source social networking server, similar to [Mastodon](https://github.com/mastodon/mastodon), [Plemora](https://github.com/Hostdon/pleroma), and [GoToSocial](https://github.com/superseriousbusiness/gotosocial).

We plan to implement the [ActivityPub](https://activitypub.rocks/) standard, allowing cross-communicative federation with other social networks.

At the moment, this software can be classified as a side project. I am currently working on basic functionality to mimic the user API of Mastodon and testing with the [Pinafore](https://github.com/Pinafore) client.  

### Authentication Flow
1. Application makes a POST request to /api/v1/apps to register application
	1. Server stores registration information and returns registration information including newly generated client ID and client secret
2. Application makes a GET request to /oauth/authorize
	1. Server returns login page
3. User enters credentials and makes a login POST request to /oauth/authorize.
	1. Server returns access code
4. Application uses client ID, client secret