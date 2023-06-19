[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

# OpenCamp CLI

Simple CLI for interacting with the recreation.gov API.

Includes polling and notifications, so you can get notified if a fully booked campground has a cancelation for the dates you're interested in. Run locally or deploy to Render.

- [Usage](#usage)
  - [Search for a campground](#search-for-a-campground)
  - [Check campground availability](#check-campground-availability)
  - [Poll campground availability](#poll-campground-availability)
  - [Polling with email notifications](#polling-with-email-notifications)
    - [Using Gmail as your SMTP server](#using-gmail-as-your-smtp-server)
- [One-Click Deployment](#one-click-deployment)
- [License](#license)

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
➜ opencamp poll 233116 09-11-2023 09-12-2023 --interval=1m
INFO[06-09|14:24:37] No sites available atm, starting polling! interval=1m0s
INFO[06-09|14:25:37] Sorry, no available campsites were found for your dates. We'll try again
INFO[06-09|14:26:37] Sorry, no available campsites were found for your dates. We'll try again
...
```

### Polling with email notifications

```
➜ opencamp poll 233116 09-11-2023 09-12-2023 --interval=1m --notify=email
In order to get notified by email, please specify your email SMTP details
SMTP Server: smtp.gmail.com
SMTP Port: 587
Email address: your-email@gmail.com
Password: *************

INFO[06-09|14:24:37] No sites available at the moment, starting polling! interval=1m0s
INFO[06-09|14:25:37] Sorry, no available campsites were found for your dates. We'll try again...
...
INFO[06-11|18:14:37] Sorry, no available campsites were found for your dates. We'll try again...

Just in! The following sites are now available for those dates:
 - Site 004             Book at: https://www.recreation.gov/camping/campsites/70286
 - Site 005             Book at: https://www.recreation.gov/camping/campsites/70079

INFO[06-11|18:14:43] Notification email sent
```

_Note: SMTP credentials are stored in memory and never echoed to stdout, but you should still be conscious of the security implications of authenticating with an SMTP server like this._

#### Using Gmail as your SMTP server
If you want to use Gmail as your smtp server and you have two factor authentication enabled, you'll need to generate an app password here: https://myaccount.google.com/apppasswords.

## One-Click Deployment

You can run the CLI as a daemon but if you don't have a machine that runs 24/7 [Render](https://render.com) has one-click deployment functionality that is similar to Heroku and has free tier. Please be mindful of polling too frequently though, once every 10 minutes is the recommended max.

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
