[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

# OpenCamp CLI

Simple CLI for interacting with the recreation.gov API.

Includes polling and notifications, so you can get notified if a fully booked campground has a cancelation for the dates you're interested in. Run locally or deploy to Render.

- [Installation](#installation)
- [Usage](#usage)
  - [Search for a campground](#search-for-a-campground)
  - [Check campground availability](#check-campground-availability)
  - [Poll campground availability](#poll-campground-availability)
- [Email Notifications](#email-notifications)
    - [Using Gmail as your SMTP server](#using-gmail-as-your-smtp-server)
- [Configuration](#configuration)
- [One-Click Deployment](#one-click-deployment)
    - [Warning ⚠️](#warning-️)
  - [Steps](#steps)
- [License](#license)

## Installation

**Option 1**: Install with Brew

```
brew tap opencamp-hq/homebrew-opencamp
brew install opencamp
```

**Option 2**: Precompiled Binaries

Precompiled binaries for the project can be found in the [Releases](https://github.com/opencamp-hq/cli/releases) section.

## Usage

### Search for a campground
```
➜ opencamp search "kirk creek"
- Kirk Creek Campground      Big Sur, California        ID: 233116
- Bird Creek Campground      Ely, Nevada                ID: 234209
- Defeated Creek Park        Carthage, Tennessee        ID: 232572
- Grassy Creek Park          Clarksville, Virginia      ID: 10107534
```

### Check campground availability
```
➜ opencamp check 233116 09-11-2023 09-12-2023
The following sites are available for those dates:
 - Site 004                  Book at: https://www.recreation.gov/camping/campsites/70286
 - Site 007                  Book at: https://www.recreation.gov/camping/campsites/70079
 - Site 008                  Book at: https://www.recreation.gov/camping/campsites/70163
 - Site 014                  Book at: https://www.recreation.gov/camping/campsites/70857
 - Site 018                  Book at: https://www.recreation.gov/camping/campsites/70573
```

### Poll campground availability
```
➜ opencamp poll 233116 09-11-2023 09-12-2023 --interval=10m
INFO[06-09|14:24:37] No sites available atm, starting polling! interval=10m0s
INFO[06-09|14:34:37] Sorry, no available campsites were found for your dates. We'll try again in 10m0s
INFO[06-09|14:44:37] Sorry, no available campsites were found for your dates. We'll try again in 10m0s
...
```

## Email Notifications
Both the `check` and `poll` commands support email notifications when a campsite is found available.

To get notified via email, you'll be prompted to supply your SMTP credentials interactively. Alternately, you can supply these credentials as environment variables (SMTP_HOST, SMTP_EMAIL, etc) or in a config.yaml to allow the tool to run in a headless mode (ie: as a cron). See [Configuration](#configuration) for more info.

```
➜ opencamp poll 233116 09-11-2023 09-12-2023 --notify=email
In order to get notified by email, please specify your email SMTP details
SMTP Server: smtp.gmail.com
SMTP Port: 587
Email address: your-email@gmail.com
Password: *************

INFO[06-09|14:24:37] No sites available at the moment, starting polling! interval=10m0s
INFO[06-09|14:34:37] Sorry, no available campsites were found for your dates. We'll try again in 10m0s
...
INFO[06-11|18:14:37] Sorry, no available campsites were found for your dates. We'll try again in 10m0s

Just in! The following sites are now available for those dates:
 - Site 004             Book at: https://www.recreation.gov/camping/campsites/70286
 - Site 005             Book at: https://www.recreation.gov/camping/campsites/70079

INFO[06-11|18:14:43] Notification email sent
```

_Note: SMTP credentials are stored in memory and not echoed to stdout, but you should still be conscious of the security implications of authenticating with an SMTP server like this._

#### Using Gmail as your SMTP server
If you want to use Gmail as your smtp server and you have two factor authentication enabled, you'll need to generate an app password here: https://myaccount.google.com/apppasswords.

## Configuration

Configuration values can be passed as command line flags, set as environment variables, or defined in a config.yaml file, with precedence defined in that order.

Example config.yaml:
```
interval: 10m
notify: email
smtp:
    host: smtp.gmail.com
    port: 587
    email: your-email@gmail.com
    password: your-password
verbose: true
```

For nested configuration, the equivalent env var is a concatenation of keys with an `_`, ie: `host` in the yaml file above becomes `SMTP_HOST`.

## One-Click Deployment

You can run the CLI as a daemon but if you don't have a machine that runs 24/7 [Render](https://render.com) provides one-click deployment functionality similar to Heroku. Based on the settings defined in the [render.yaml](render.yaml) file, the CLI will be deployed as a cron job that runs once every 10m.

#### Warning ⚠️
Please be mindful of polling too frequently, once every 10 minutes is the recommended max. Running the tool at a higher frequency is unlikely to make a difference and risks recreation.gov authenticating their API and breaking this type of access for everyone.

### Steps
1. Sign up for [Render](https://render.com). In order to deploy the CLI as a cron job, you'll need to enter a credit card, but the billing rate is $0.0094/hr
   1. Note: The minimum for a cron job is $1/mo, which is what you'll probably end up spending
2. Click the Deploy to Render button below
3. Enter values for the environment variables. Example values for each environment variable are in the [render.yaml](render.yaml) file
   1. Note: You'll need to determine the id of the campground you're interested in by running the tool locally first, ie: [Search for a campground](#search-for-a-campground)
4. Check the logs to ensure the cronjob is successfully running
5. If a campground is available, you'll get an e-mail with a link to book

[![Deploy to Render](https://render.com/images/deploy-to-render-button.svg)](https://render.com/deploy?repo=https://github.com/opencamp-hq/cli)


## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

[contributors-shield]: https://img.shields.io/github/contributors/opencamp-hq/cli?style=for-the-badge
[contributors-url]: https://github.com/opencamp-hq/cli/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/opencamp-hq/cli?style=for-the-badge
[forks-url]: https://github.com/opencamp-hq/cli/network/members
[stars-shield]: https://img.shields.io/github/stars/opencamp-hq/cli?style=for-the-badge
[stars-url]: https://github.com/opencamp-hq/cli/stargazers
[issues-shield]: https://img.shields.io/github/issues/opencamp-hq/cli?style=for-the-badge
[issues-url]: https://github.com/opencamp-hq/cli/issues
[license-shield]: https://img.shields.io/github/license/opencamp-hq/cli?style=for-the-badge
[license-url]: https://github.com/opencamp-hq/cli/blob/main/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/kylechadha
