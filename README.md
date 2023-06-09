[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

# OpenCamp CLI

Simple CLI for interacting with the recreation.gov API.

## Usage

### Search for a campground Id
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
Sorry we didn't find any available campsites!
```

### Poll campground availability
```
➜ opencamp poll 233116 09-11-2023 09-12-2023 --interval=1m
INFO[06-09|14:24:37] No sites available atm, starting polling! interval=1m0s
INFO[06-09|14:25:37] Sorry, no available campsites were found for your dates. We'll try again in 1m0s 
```

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
[license-url]: https://github.com/opencamp-hq/cli/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/kylechadha
