# Frens


Frens is an open source social networking server written in Go, similar to [Mastodon](https://github.com/mastodon/mastodon), [Plemora](https://github.com/Hostdon/pleroma), or [GoToSocial](https://github.com/superseriousbusiness/gotosocial). 

This project differs from others given that it provides the speed and efficiency of a modern system languages while providing the most robust feature set. At the moment, this software can be classified as a side project. I am currently working on basic functionality to mimic the user API of Mastodon and testing with the [Pinafore](https://github.com/Pinafore) client. 

Anyone interested in the project is more than welcome to contribute.

### Built With
- https://go.dev/
- https://github.com/gofiber/fiber
- https://github.com/spf13/cobra
- https://github.com/spf13/viper
- https://github.com/dgraph-io/badger
- https://github.com/postgres/postgres


## Database Setup
sudo -u postgres psql
CREATE ROLE username WITH LOGIN SUPERUSER PASSWORD 'password';
CREATE DATABASE db_name;